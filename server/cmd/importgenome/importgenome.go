package importgenome

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
)

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
		_frame, _ := strconv.Atoi(rec.Phase)
		var frame int8 = int8(_frame)
		c.Cds = append(c.Cds,
			client.Cds.Create().
				SetTranscriptID(transcriptId).
				SetSeqname(seqname).
				SetStart(start).
				SetEnd(end).
				SetPhase(frame).
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
			// check gene id
			geneId, found := rec.Attributes.Get("Parent")

			if !found {
				return nil, fmt.Errorf("")
			}

			id, found := rec.Attributes.Get("ID")

			if !found {
				return nil, fmt.Errorf("")
			}

			recs := TranscriptRecords{GeneId: geneId, Type: rec.Type, SeqName: rec.Seqname, Strand: rec.Strand, Records: []gffio.GffRecord{rec}}
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
				return nil, fmt.Errorf("%s should have Parent attributes", rec.Type)
			}

			recs := id2rec[id]
			recs.Records = append(recs.Records, rec)

			id2rec[id] = recs
		}
	}

	return id2rec, nil
}
