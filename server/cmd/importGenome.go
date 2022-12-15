/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bufio"
	"context"
	"fmt"
	"genomedb/ent"
	"genomedb/gffio"
	"genomedb/seq"
	"genomedb/seqio"
	"os"
	"strconv"
	"strings"

	_ "github.com/mattn/go-sqlite3"
	"github.com/spf13/cobra"
)

// importGenomeCmd represents the importGenome command
var genomeName string
var genomeFasta string
var genomeGff string
var databaseUri string
var codonCode int32

type TranscriptDataCreates struct {
	Cds           []*ent.CdsCreate
	Exon          []*ent.ExonCreate
	FivePrimeUtr  []*ent.FivePrimeUtrCreate
	ThreePrimeUtr []*ent.ThreePrimeUtrCreate
}

func (c *TranscriptDataCreates) PushGffRecord(client *ent.Client, transcriptId string, rec gffio.GffRecord) {
	seqname := rec.Seqname
	start := rec.Start
	end := rec.End
	strand := rec.Strand

	if rec.IsCds() {
		_frame, _ := strconv.Atoi(rec.Frame)
		var frame int8 = int8(_frame)
		c.Cds = append(c.Cds,
			client.Cds.Create().
				SetTranscriptID(transcriptId).
				SetSeqname(seqname).
				SetStart(start).
				SetEnd(end).
				SetFrame(frame).
				SetStrand(strand),
		)
	} else if rec.IsExon() {
		c.Exon = append(
			c.Exon,
			client.Exon.Create().
				SetTranscriptID(transcriptId).
				SetSeqname(seqname).
				SetStart(start).
				SetEnd(end).
				SetStrand(strand),
		)
	} else if rec.IsFivePrimeUtr() {
		c.FivePrimeUtr = append(
			c.FivePrimeUtr,
			client.FivePrimeUtr.Create().
				SetTranscriptID(transcriptId).
				SetSeqname(seqname).
				SetStart(start).
				SetEnd(end).
				SetStrand(strand),
		)
	} else if rec.IsThreePrimeUtr() {
		c.ThreePrimeUtr = append(c.ThreePrimeUtr,
			client.ThreePrimeUtr.Create().
				SetTranscriptID(transcriptId).
				SetSeqname(seqname).
				SetStart(start).
				SetEnd(end).
				SetStrand(strand),
		)
	}
}

func (c TranscriptDataCreates) BulkCreate(client *ent.Client, ctx context.Context, stepNum int) error {
	fmt.Printf("Upsert %d Exons ...\n", len(c.Exon))
	for i := 0; i < len(c.Exon); i += stepNum {
		var dtos []*ent.ExonCreate
		if i+stepNum > len(c.Exon) {
			dtos = c.Exon[i:]
		} else {
			dtos = c.Exon[i : i+stepNum]
		}
		if err := client.Exon.CreateBulk(dtos...).OnConflict().UpdateNewValues().Exec(ctx); err != nil {
			return err
		}
	}

	fmt.Printf("Upsert %d Cds ...\n", len(c.Cds))
	for i := 0; i < len(c.Cds); i += stepNum {
		var dtos []*ent.CdsCreate
		if i+stepNum > len(c.Cds) {
			dtos = c.Cds[i:]
		} else {
			dtos = c.Cds[i : i+stepNum]
		}
		if err := client.Cds.CreateBulk(dtos...).OnConflict().UpdateNewValues().Exec(ctx); err != nil {
			return err
		}
	}

	fmt.Printf("Upsert %d FivePrimeUtrs ...\n", len(c.FivePrimeUtr))
	for i := 0; i < len(c.FivePrimeUtr); i += stepNum {
		var dtos []*ent.FivePrimeUtrCreate
		if i+stepNum > len(c.FivePrimeUtr) {
			dtos = c.FivePrimeUtr[i:]
		} else {
			dtos = c.FivePrimeUtr[i : i+stepNum]
		}
		if err := client.FivePrimeUtr.CreateBulk(dtos...).OnConflict().UpdateNewValues().Exec(ctx); err != nil {
			return err
		}
	}

	fmt.Printf("Upsert %d ThreePrimeUtrs ...\n", len(c.ThreePrimeUtr))
	for i := 0; i < len(c.ThreePrimeUtr); i += stepNum {
		var dtos []*ent.ThreePrimeUtrCreate
		if i+stepNum > len(c.ThreePrimeUtr) {
			dtos = c.ThreePrimeUtr[i:]
		} else {
			dtos = c.ThreePrimeUtr[i : i+stepNum]
		}
		if err := client.ThreePrimeUtr.CreateBulk(dtos...).OnConflict().UpdateNewValues().Exec(ctx); err != nil {
			return err
		}
	}

	return nil
}

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

		seqname2seq, fastaReadError := readGenomeFasta(genomeFasta)
		fmt.Println("Finish fasta reading...")

		if fastaReadError != nil {
			return fastaReadError
		}

		transcriptId2recs, gtfReadError := readGffFile(genomeGff)
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
		transcriptDataCreates := TranscriptDataCreates{}

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

func readGenomeFasta(fastaPath string) (map[string]string, error) {
	r, err := os.Open(fastaPath)

	if err != nil {
		return nil, err
	}

	defer r.Close()

	br := bufio.NewScanner(r)

	parser := seqio.NewFastaParser()

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

// gene -(Parent)-> mRNA, rRNA, miRNA, pre-miRNA -(Parent)-> exon, cds, five_prime_UTR, three_rime_UTR
// map[transcriptId][]transcript_rec

// reading a gff file
type TranscriptRecords struct {
	GeneId  string
	Type    string
	SeqName string
	Strand  string
	Records []gffio.GffRecord
}

type TranscriptSeq struct {
	Genome  string
	Mrna    string
	Cds     string
	Protein string
}

func (r TranscriptRecords) ToTranscriptSeq(scaffoldSeq string, codonTable seq.CodonTable) TranscriptSeq {
	var genome string
	var mrna strings.Builder
	var cds strings.Builder

	for _, rec := range r.Records {
		seq := scaffoldSeq[rec.Start-1 : rec.End]

		if rec.IsGeneChild() {
			genome = seq
		} else if rec.IsExon() {
			mrna.WriteString(seq)
		} else if rec.IsCds() {
			cds.WriteString(seq)
		} else {
			// unknown feature
		}
	}

	var Cds string
	var Mrna string
	if r.Strand == "+" {
		Cds = cds.String()
		Mrna = mrna.String()
	} else {
		Cds = seq.ReverseComplement(cds.String())
		Mrna = seq.ReverseComplement(cds.String())
	}

	prot, err := codonTable.Transrate(Cds)

	if err != nil {
		fmt.Printf("%v can be invalid CDS, length=%d, length mod 3=%d\n", r.GeneId, len(cds.String()), len(cds.String())%3)
	}

	return TranscriptSeq{Genome: genome, Cds: Cds, Mrna: Mrna, Protein: prot}
}

func readGffFile(gtfPath string) (map[string]TranscriptRecords, error) {
	r, err := os.Open(gtfPath)

	if err != nil {
		return nil, err
	}

	defer r.Close()

	br := bufio.NewScanner(r)

	parser := gffio.NewGffParser()

	for br.Scan() {
		parser.ConsumeLine(br.Text())
	}

	id2rec := map[string]TranscriptRecords{}

	for _, rec := range parser.Records {
		if rec.IsGeneChild() {
			// check gene id
			geneId, found := rec.Attributes.Get("Parent")

			if !found {
				return nil, fmt.Errorf("")
			}

			id, found := rec.Attributes.Get("ID")

			if !found {
				return nil, fmt.Errorf("")
			}

			recs := TranscriptRecords{GeneId: geneId, Type: rec.Feature, SeqName: rec.Seqname, Strand: rec.Strand, Records: []gffio.GffRecord{rec}}
			id2rec[id] = recs
		} else if rec.IsExon() {
			// exon can have multiple parents
			ids, found := rec.Attributes.GetAll("Parent")
			if !found {
				return nil, fmt.Errorf("Exon must have Parent attributes.")
			}

			for _, id := range ids {
				recs := id2rec[id]
				recs.Records = append(recs.Records, rec)

				id2rec[id] = recs
			}

		} else if rec.IsCds() || rec.IsFivePrimeUtr() || rec.IsThreePrimeUtr() {
			id, found := rec.Attributes.Get("Parent")
			if !found {
				return nil, fmt.Errorf("%s should have Parent attributes", rec.Feature)
			}

			recs := id2rec[id]
			recs.Records = append(recs.Records, rec)

			id2rec[id] = recs
		}
	}

	return id2rec, nil
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
