/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bufio"
	"context"
	"fmt"
	"genomedb/bio/goterm"
	"genomedb/ent"
	"genomedb/util"
	"os"

	entgoterm "genomedb/ent/goterm"

	"github.com/spf13/cobra"
)

var oboPath string

// importAnnotationBaseCmd represents the importAnnotationBase command
var importAnnotationBaseCmd = &cobra.Command{
	Use:   "importAnnotationBase",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	RunE: func(cmd *cobra.Command, args []string) error {
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

		go_terms, err := ReadOboFile(oboPath)

		if err != nil {
			return err
		}

		dag := goterm.CreateGoDAGFromGoterms(go_terms)

		goTermCreates := []*ent.GoTermCreate{}

		for _, go_term := range go_terms {
			goDagRecord := dag[go_term.Id]
			if err != nil {
				return err
			}

			goTermCreates = append(goTermCreates,
				client.GoTerm.Create().
					SetID(go_term.Id).
					SetName(go_term.Name).
					SetNamespace(entgoterm.Namespace(go_term.Namespace)).
					SetDef(go_term.Def).
					SetDepth(int32(goDagRecord.Depth)).
					SetLevel(int32(goDagRecord.Level)),
			)
		}

		batch_iter := util.NewBatchIter(1000, goTermCreates)

		for {
			dtos, finish := batch_iter.Next()

			if finish {
				break
			}

			if err := client.GoTerm.CreateBulk(dtos...).OnConflict().UpdateNewValues().Exec(ctx); err != nil {
				return err
			}
		}

		return nil
	},
}

func ReadOboFile(path string) ([]goterm.GoTerm, error) {
	r, err := os.Open(path)
	parser := goterm.NewOboParser()

	if err != nil {
		return parser.GoTerms, err
	}

	defer r.Close()

	br := bufio.NewScanner(r)

	for br.Scan() {
		if err := parser.ConsumeLine(br.Text()); err != nil {
			return parser.GoTerms, err
		}
	}

	return parser.GoTerms, nil
}

func init() {
	rootCmd.AddCommand(importAnnotationBaseCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// importAnnotationBaseCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// importAnnotationBaseCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	importAnnotationBaseCmd.Flags().StringVarP(&oboPath, "obo", "", "", "Path to obo file")
}
