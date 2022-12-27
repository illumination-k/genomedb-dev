package genome

import (
	"bufio"
	"context"
	"fmt"
	"genomedb/bio/gffio"
	"genomedb/bio/seq"
	"genomedb/bio/seqio"
	"genomedb/ds/hashset"
	"genomedb/ent"
	"os"
	"strings"
)

func Run(genomeName string, genomeFasta string, genomeGff string, databaseUri string, codonCode int32) error {
	fmt.Printf("fasta: %v\n gtf : %v\n", genomeFasta, genomeGff)

	seqname2seq, fastaReadError := ReadGenomeFasta(genomeFasta)
	fmt.Println("Finish fasta reading...")

	if fastaReadError != nil {
		return fastaReadError
	}

	transcriptId2recs, gtfReadError := ReadGffFile(genomeGff)
	fmt.Println("Finish gtf reading...")

	if gtfReadError != nil {
		return gtfReadError
	}

	// # import into database
	// ## init client
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

	// ## create genome
	if err := client.Genome.Create().SetID(genomeName).SetCodonTable(codonCode).OnConflict().UpdateNewValues().Exec(ctx); err != nil {
		return err
	}

	// ## set codon table for translation
	codonTable := seq.NewCodonTable(codonCode)

	genes := hashset.NewHashSet[string]()
	transcriptDtos := []*ent.TranscriptCreate{}

	for transcriptId, rec := range transcriptId2recs {
		scaffoldSeq, ok := seqname2seq[rec.SeqName]

		genes.Add(rec.LocusId)

		if !ok {
			return fmt.Errorf("seqname=%v does not exist in %v\n rec: %v\n", rec.SeqName, genomeFasta, rec)
		}

		cds := rec.CdsSequence(scaffoldSeq)
		protein, _ := codonTable.Transrate(cds)
		exon := rec.ExonSequence(scaffoldSeq)
		genomic := rec.GenomicSequence(scaffoldSeq)

		transcriptDtos = append(
			transcriptDtos,
			client.Transcript.
				Create().
				SetID(transcriptId).
				SetStrand(rec.Strand).
				SetType(rec.Type).
				SetSource(rec.Source).
				SetSeqname(rec.SeqName).
				SetStart(rec.Start).
				SetEnd(rec.End).
				SetGenomicSequence(genomic).
				SetExon(rec.ExonRecords).
				SetExonSequence(exon).
				SetCds(rec.CdsRecords).
				SetCdsSequence(cds).
				SetProteinSequence(protein).
				SetFivePrimeUtr(rec.FivePrimeUtrRecords).
				SetThreePrimeUtr(rec.ThreePrimeUtrRecords).
				SetLocusID(rec.LocusId),
		)
	}

	geneDtos := []*ent.LocusCreate{}
	for geneId := range genes {
		geneDtos = append(geneDtos, client.Locus.Create().SetID(geneId).SetGenomeID(genomeName))
	}

	stepNum := 100

	fmt.Printf("Upsert %d Locuses ...\n", len(geneDtos))
	for i := 0; i < len(geneDtos); i += stepNum {
		var dtos []*ent.LocusCreate
		if i+stepNum > len(geneDtos) {
			dtos = geneDtos[i:]
		} else {
			dtos = geneDtos[i : i+stepNum]
		}

		if err := client.Locus.CreateBulk(dtos...).OnConflict().UpdateNewValues().Exec(ctx); err != nil {
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

	return nil
}

func ReadGenomeFasta(fastaPath string) (map[string]string, error) {
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
	LocusId              string
	Source               string
	Type                 string
	SeqName              string
	Strand               string
	Start                int32
	End                  int32
	CdsRecords           []gffio.GffRecord
	ExonRecords          []gffio.GffRecord
	FivePrimeUtrRecords  []gffio.GffRecord
	ThreePrimeUtrRecords []gffio.GffRecord
}

func NewTranscriptRecords(record gffio.GffRecord) (TranscriptRecords, error) {
	transcriptRecord := new(TranscriptRecords)

	if !record.IsGeneChild() {
		return *transcriptRecord, fmt.Errorf("New transcript record should be created from gene child")
	}

	geneId, found := record.Attributes.Get("Parent")

	if !found {
		return *transcriptRecord, fmt.Errorf("")
	}

	transcriptRecord.LocusId = geneId
	transcriptRecord.Source = record.Source
	transcriptRecord.Type = record.Type
	transcriptRecord.SeqName = record.Seqname
	transcriptRecord.Start = record.Start
	transcriptRecord.End = record.End
	transcriptRecord.Strand = record.Strand

	return *transcriptRecord, nil
}

func (r *TranscriptRecords) PushGffRecord(record gffio.GffRecord) error {
	if record.IsCds() {
		r.CdsRecords = append(r.CdsRecords, record)
	} else if record.IsExon() {
		r.ExonRecords = append(r.ExonRecords, record)
	} else if record.IsFivePrimeUtr() {
		r.FivePrimeUtrRecords = append(r.FivePrimeUtrRecords, record)
	} else if record.IsThreePrimeUtr() {
		r.ThreePrimeUtrRecords = append(r.ThreePrimeUtrRecords, record)
	} else {
	}

	return nil
}

func (r *TranscriptRecords) CdsSequence(scaffoldSeq string) string {
	var cds strings.Builder

	for _, r := range r.CdsRecords {
		cds.WriteString(r.ExtractSequence(scaffoldSeq))
	}

	return cds.String()
}

func (r *TranscriptRecords) ExonSequence(scaffoldSeq string) string {
	var exon strings.Builder
	for _, r := range r.ExonRecords {
		exon.WriteString(r.ExtractSequence(scaffoldSeq))
	}

	return exon.String()
}

func (r TranscriptRecords) GenomicSequence(scaffoldSeq string) string {
	sequence := scaffoldSeq[r.Start:r.End]

	if r.Strand == "-" {
		sequence = seq.ReverseComplement(sequence)
	}

	return sequence
}

func ReadGffFile(gtfPath string) (map[string]TranscriptRecords, error) {
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
			id, found := rec.Attributes.Get("ID")

			if !found {
				return nil, fmt.Errorf("")
			}

			var recs TranscriptRecords
			if recs, err = NewTranscriptRecords(rec); err != nil {
				return nil, fmt.Errorf("")
			}
			id2rec[id] = recs

		} else if rec.IsExon() {
			// exon can have multiple parents
			ids, found := rec.Attributes.GetAll("Parent")
			if !found {
				return nil, fmt.Errorf("Exon must have Parent attributes.")
			}

			for _, id := range ids {
				recs := id2rec[id]
				recs.PushGffRecord(rec)

				id2rec[id] = recs
			}

		} else if rec.IsCds() || rec.IsFivePrimeUtr() || rec.IsThreePrimeUtr() {
			id, found := rec.Attributes.Get("Parent")
			if !found {
				return nil, fmt.Errorf("%s should have Parent attributes", rec.Type)
			}

			recs := id2rec[id]
			recs.PushGffRecord(rec)

			id2rec[id] = recs
		}
	}

	return id2rec, nil
}
