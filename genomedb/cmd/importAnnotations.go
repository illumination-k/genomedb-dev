/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bufio"
	"context"
	"fmt"
	"genomedb/bio/interproscan"
	"genomedb/ent"
	"os"

	"github.com/spf13/cobra"
)

// importAnnotationsCmd represents the importAnnotations command
var importAnnotationsCmd = &cobra.Command{
	Use:   "importAnnotations",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("importAnnotations called")
		interproscanRecords, err := readInterproscanTsv("")

		if err != nil {
			return err
		}

		client, err := ent.Open("sqlite3", databaseUri)
		if err != nil {
			return err
		}
		defer client.Close()

		// ## Run the auto migration tool.
		ctx := context.Background()
		if err := client.Schema.Create(ctx); err != nil {
			return fmt.Errorf("failed creating schema resources: %v", err)
		}

		dbStatus, err = CheckDBStatus(ctx, client)

		if err != nil {
			return err
		}

		domainAnnotationCreate := []*ent.DomainAnnotationCreate{}
		domainToTranscriptCreate := []*ent.DomainAnnotationToTranscriptCreate{}
		goTermToTranscriptCreat := []*ent.GoTermOnTranscriptsCreate{}
		for _, record := range interproscanRecords {
			domainAnnotationCreate = append(
				domainAnnotationCreate,
				client.DomainAnnotation.
					Create().
					SetID(record.SignatureAccession).
					SetDescription(record.SignatureDescription).
					SetAnalysis(record.Analysis),
			)

			domainToTranscriptCreate = append(domainToTranscriptCreate,
				client.DomainAnnotationToTranscript.
					Create().
					SetDomainAnnotationID(record.SignatureAccession).
					SetTranscriptID(record.Accession).
					SetStart(int32(record.Start)).
					SetStop(int32(record.Stop)).
					SetScore(record.Score),
			)

			if dbStatus.IsGoTermImported {
				for _, go_term_id := range record.GoTermIDs {
					goTermToTranscriptCreat = append(goTermToTranscriptCreat,
						client.GoTermOnTranscripts.Create().
							SetEvidenceCode("ISO").
							SetGoTermID(go_term_id).
							SetTranscriptID(record.Accession),
					)
				}
			}
		}

		if err := client.DomainAnnotation.
			CreateBulk(domainAnnotationCreate...).
			OnConflict().
			UpdateNewValues().
			Exec(ctx); err != nil {
			return err
		}

		if err := client.DomainAnnotationToTranscript.
			CreateBulk(domainToTranscriptCreate...).
			OnConflict().
			UpdateNewValues().
			Exec(ctx); err != nil {
			return err
		}

		if err := client.GoTermOnTranscripts.
			CreateBulk(goTermToTranscriptCreat...).
			OnConflict().
			UpdateNewValues().
			Exec(ctx); err != nil {
			return err
		}

		return nil
	},
}

func readInterproscanTsv(path string) ([]interproscan.InterproscanRecord, error) {
	r, err := os.Open(path)
	parser := interproscan.InterproscanTsvParser{}

	if err != nil {
		return parser.Records, err
	}

	defer r.Close()

	br := bufio.NewScanner(r)

	for br.Scan() {
		if err := parser.ConsumeLine(br.Text()); err != nil {
			return parser.Records, err
		}
	}

	return parser.Records, nil
}

func init() {
	rootCmd.AddCommand(importAnnotationsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// importAnnotationsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// importAnnotationsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
