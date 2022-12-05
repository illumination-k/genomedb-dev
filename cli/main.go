package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	importTranscriptsCmd := flag.NewFlagSet("import-transcripts", flag.ExitOnError)

	if len(os.Args) < 2 {
		fmt.Println("Expected import-transcripts")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "import-transcripts":
		name := importTranscriptsCmd.String("name", "", "genome name")
		cds := importTranscriptsCmd.String("cds", "", "Path to CDS fasta")
		mrna := importTranscriptsCmd.String("mrna", "", "Path to mrna fasta")
		gene := importTranscriptsCmd.String("gene", "", "Path to gene fasta")
		protein := importTranscriptsCmd.String("protein", "", "Path to protein fasta")

		importTranscriptsCmd.Parse(os.Args[2:])
		fmt.Printf("%v %v %v %v %v", *name, *cds, *mrna, *gene, *protein)
	default:
		fmt.Println("Expected import-transcripts")
		os.Exit(1)
	}

	fmt.Printf("OK!")
}
