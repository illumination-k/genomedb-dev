/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"fmt"
	"genomedb/cmd/importgenome"
	"genomedb/ent"
	"genomedb/seq"

	_ "github.com/mattn/go-sqlite3"
	"github.com/spf13/cobra"
)

// importGenomeCmd represents the importGenome command
var genomeName string
var genomeFasta string
var genomeGff string
var databaseUri string
var codonCode int32

var importGenomeCmd = &cobra.Command{
	Use:   "importGenome",
	Short: "Import genome information",
	Long:  `Import genomic features from a genomic fasta file and a genomic gff3 file into the database.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		// if databaseUri == "" {
		// 	databaseUri = os.Getenv("DATABASE_URI")
		// }

		// if databaseUri == "" {
		// 	return fmt.Errorf("You should specify DATABASE URI by enviroment variables or flag")
		// }

		fmt.Printf("fasta: %v\n gtf : %v\n", genomeFasta, genomeGff)

		seqname2seq, fastaReadError := importgenome.ReadGenomeFasta(genomeFasta)
		fmt.Println("Finish fasta reading...")

		if fastaReadError != nil {
			return fastaReadError
		}

		transcriptId2recs, gtfReadError := importgenome.ReadGffFile(genomeGff)
		fmt.Println("Finish gtf reading...")

		if gtfReadError != nil {
			return gtfReadError
		}

		// # import into database
		// ## init client
		client, err := ent.Open("sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
		if err != nil {
			return err
		}
		defer client.Close()

		// ## Run the auto migration tool.
		ctx := context.Background()
		if err := client.Schema.Create(ctx); err != nil {
			return fmt.Errorf("failed creating schema resources: %v", err)
		}

		// ## create genome

		if err := client.Genome.Create().SetID(genomeName).SetCodonTable(codonCode).OnConflict().UpdateNewValues().Exec(ctx); err != nil {
			return err
		}

		// ## set codon table for translation
		codonTable := seq.NewCodonTable(codonCode)

		genes := map[string]struct{}{}
		transcriptDtos := []*ent.TranscriptCreate{}
		transcriptDataCreates := importgenome.TranscriptDataCreates{}

		for transcriptId, rec := range transcriptId2recs {
			scaffoldSeq, ok := seqname2seq[rec.SeqName]

			genes[rec.GeneId] = struct{}{}

			if !ok {
				return fmt.Errorf("seqname=%v does not exist in %v\n rec: %v\n", rec.SeqName, genomeFasta, rec)
			}

			transcriptSeq := rec.ToTranscriptSeq(scaffoldSeq, codonTable)
			transcriptDtos = append(
				transcriptDtos,
				client.Transcript.
					Create().
					SetID(transcriptId).
					SetStrand(rec.Strand).
					SetType(rec.Type).
					SetGenomeSeq(transcriptSeq.Genome).
					SetTranscriptSeq(transcriptSeq.Mrna).
					SetCdsSeq(transcriptSeq.Cds).
					SetProteinSeq(transcriptSeq.Protein).
					SetGeneID(rec.GeneId),
			)

			for _, gffrec := range rec.Records {
				transcriptDataCreates.PushGffRecord(client, transcriptId, gffrec)
			}
		}

		geneDtos := []*ent.GeneCreate{}
		for geneId := range genes {
			geneDtos = append(geneDtos, client.Gene.Create().SetID(geneId).SetGenomeID(genomeName))
		}

		stepNum := 100

		fmt.Printf("Upsert %d Genes ...\n", len(geneDtos))
		for i := 0; i < len(geneDtos); i += stepNum {
			var dtos []*ent.GeneCreate
			if i+stepNum > len(geneDtos) {
				dtos = geneDtos[i:]
			} else {
				dtos = geneDtos[i : i+stepNum]
			}

			if err := client.Gene.CreateBulk(dtos...).OnConflict().UpdateNewValues().Exec(ctx); err != nil {
				return err
			}
		}

		fmt.Printf("Upsert %d Transcript ...\n", len(transcriptDtos))
		for i := 0; i < len(transcriptDtos); i += stepNum {
			var dtos []*ent.TranscriptCreate
			if i+stepNum > len(transcriptDtos) {
				dtos = transcriptDtos[i:]
			} else {
				dtos = transcriptDtos[i : i+stepNum]
			}
			if err := client.Transcript.CreateBulk(dtos...).OnConflict().UpdateNewValues().Exec(ctx); err != nil {
				return err
			}
		}

		transcriptDataCreates.BulkCreate(client, ctx, stepNum)

		return nil
	},
}

func init() {
	rootCmd.AddCommand(importGenomeCmd)

	// Here you will define your flags and configuration settings.
	importGenomeCmd.Flags().StringVarP(&genomeName, "genomeName", "n", "", "Genome Name")
	importGenomeCmd.Flags().StringVarP(&genomeFasta, "fasta", "f", "", "Path to the genomic fasta file")
	importGenomeCmd.Flags().StringVarP(&genomeGff, "gtf", "g", "", "Path to the genomic gtf file")
	importGenomeCmd.Flags().StringVarP(&databaseUri, "database-uri", "d", "", "Database Uri")
	importGenomeCmd.Flags().Int32VarP(&codonCode, "codon-code", "c", 1, "Codon Code")

	importGenomeCmd.MarkFlagRequired("fasta")
	importGenomeCmd.MarkFlagRequired("gtf")
}
