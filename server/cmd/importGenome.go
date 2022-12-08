/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bufio"
	"fmt"
	"genomedb/bioio"
	"os"

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
	PreRunE: func(cmd *cobra.Command, args []string) error {
		databaseUri = os.Getenv("DATABASE_URI")
		fmt.Println(databaseUri)
		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
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

		for transcriptId, recs := range transcriptId2recs {
			scaffoldSeq, ok := seqname2seq[recs.SeqName]

			if !ok {
				return fmt.Errorf("seqname=%v does not exist in %v", recs.SeqName, genomeFasta)
			}

			for _, rec := range recs.Records {
				seq := scaffoldSeq[rec.Start-1 : rec.End-1]
				fmt.Printf(">%v %v %v:%d:%d\n%v\n", transcriptId, rec.Feature, rec.Seqname, rec.Start, rec.End, seq)
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

	importGenomeCmd.MarkFlagRequired("fasta")
	importGenomeCmd.MarkFlagRequired("gtf")
}
