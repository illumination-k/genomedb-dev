/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"genomedb/cmd/importgenome"

	_ "github.com/mattn/go-sqlite3"
	"github.com/spf13/cobra"
)

// importGenomeCmd represents the importGenome command
var genomeName string
var genomeFasta string
var genomeGff string
var codonCode int32

var importGenomeCmd = &cobra.Command{
	Use:   "importGenome",
	Short: "Import genome information",
	Long:  `Import genomic features from a genomic fasta file and a genomic gff3 file into the database.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return importgenome.Run(genomeName, genomeFasta, genomeGff, databaseUri, codonCode)
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
