/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bufio"
	"context"
	"fmt"
	"genomedb/bioio"
	"genomedb/ent"
	"os"
	"strings"

	_ "github.com/mattn/go-sqlite3"
	"github.com/spf13/cobra"
)

// importGenomeCmd represents the importGenome command
var genomeFasta string
var genomeGtf string
var databaseUri string

var importGenomeCmd = &cobra.Command{
	Use:   "importGenome",
	Short: "Import genome information",
	Long:  `Import genomic features from a genomic fasta file and a genomic gtf file into the database.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if databaseUri == "" {
			databaseUri = os.Getenv("DATABASE_URI")
		}

		if databaseUri == "" {
			return fmt.Errorf("You should specify DATABASE URI by enviroment variables or flag")
		}

		fmt.Printf("fasta: %v\n gtf : %v\n", genomeFasta, genomeGtf)

		seqname2seq, fastaReadError := readGenomeFasta(genomeFasta)
		fmt.Println("Finish fasta reading...")

		if fastaReadError != nil {
			return fastaReadError
		}

		transcriptId2recs, gtfReadError := readGtfFile(genomeGtf)
		fmt.Println("Finish gtf reading...")

		if gtfReadError != nil {
			return gtfReadError
		}

		// import into database
		client, err := ent.Open("sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
		if err != nil {
			return err
		}
		defer client.Close()

		// Run the auto migration tool.
		ctx := context.Background()
		if err := client.Schema.Create(ctx); err != nil {
			return fmt.Errorf("failed creating schema resources: %v", err)
		}

		transcriptDtos := []*ent.TranscriptCreate{}
		for transcriptId, recs := range transcriptId2recs {
			scaffoldSeq, ok := seqname2seq[recs.SeqName]

			if !ok {
				return fmt.Errorf("seqname=%v does not exist in %v", recs.SeqName, genomeFasta)
			}

			var genome string
			var mrna strings.Builder
			var cds strings.Builder

			for _, rec := range recs.Records {
				seq := scaffoldSeq[rec.Start-1 : rec.End-1]

				if rec.Feature == "mRNA" || rec.Feature == "rRNA" {
					genome = seq
				} else if rec.Feature == "exon" {
					mrna.WriteString(seq)
				} else if rec.Feature == "CDS" {
					cds.WriteString(seq)
				} else {
					// unknown feature
				}
			}

			transcriptDtos = append(
				transcriptDtos,
				client.Transcript.
					Create().
					SetTranscriptId(transcriptId).
					SetGeneId(recs.GeneId).
					SetCds(cds.String()).
					SetGenome(genome).
					SetMrna(mrna.String()).
					SetProtein(""),
			)

			if len(transcriptDtos) == 10 {
				err = client.Transcript.CreateBulk(transcriptDtos...).OnConflict().UpdateNewValues().Exec(ctx)

				if err != nil {
					return err
				}

				transcriptDtos = []*ent.TranscriptCreate{}
			}
		}

		return nil
	},
}

func readGenomeFasta(fastaPath string) (map[string]string, error) {
	r, err := os.Open(fastaPath)

	if err != nil {
		return nil, err
	}

	defer r.Close()

	br := bufio.NewScanner(r)

	parser := bioio.NewFastaParser()

	for br.Scan() {
		parser.ConsumeLine(br.Text())
	}

	if err := br.Err(); err != nil {
		return nil, err
	}

	parser.Flush()

	seqname2seq := map[string]string{}
	for _, rec := range parser.Records {
		seqname2seq[rec.Id] = rec.Seq
	}

	return seqname2seq, nil
}

type GffRecords struct {
	GeneId  string
	SeqName string
	Records []bioio.GffRecord
}

func readGtfFile(gtfPath string) (map[string]GffRecords, error) {
	r, err := os.Open(gtfPath)

	if err != nil {
		return nil, err
	}

	defer r.Close()

	br := bufio.NewScanner(r)

	parser := bioio.NewGffParser()

	for br.Scan() {
		parser.ConsumeLine(br.Text())
	}

	id2rec := map[string]GffRecords{}
	for _, rec := range parser.Records {
		// check transcript id
		transcriptId := ""

		v, ok := rec.Attributes["ID"]

		if ok {
			transcriptId = v
		}

		v, ok = rec.Attributes["Parent"]

		if ok {
			transcriptId = v
		}

		if transcriptId == "" {
			return nil, fmt.Errorf("ID or Parent is required in %v", rec)
		}

		// check gene id
		geneId := rec.Attributes["geneID"]

		recs, exist := id2rec[transcriptId]

		if exist {
			if geneId != "" {
				recs.GeneId = geneId
			}

			recs.Records = append(recs.Records, rec)
			id2rec[transcriptId] = recs
		} else {
			recs = GffRecords{GeneId: geneId, SeqName: rec.Seqname, Records: []bioio.GffRecord{rec}}
			id2rec[transcriptId] = recs
		}
	}

	return id2rec, nil
}

func init() {
	rootCmd.AddCommand(importGenomeCmd)

	// Here you will define your flags and configuration settings.
	importGenomeCmd.Flags().StringVarP(&genomeFasta, "fasta", "f", "", "Path to the genomic fasta file")
	importGenomeCmd.Flags().StringVarP(&genomeGtf, "gtf", "g", "", "Path to the genomic gtf file")
	importGenomeCmd.Flags().StringVarP(&databaseUri, "database-uri", "d", "", "Database Uri")

	importGenomeCmd.MarkFlagRequired("fasta")
	importGenomeCmd.MarkFlagRequired("gtf")
}
