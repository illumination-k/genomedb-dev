// https://github.com/The-Sequence-Ontology/Specifications/blob/master/gff3.md

package gffio

import (
	"fmt"
	"genomedb/seq"
	"strconv"
	"strings"
)

type MultiMap map[string][]string

func NewMultiMap() MultiMap {
	return make(map[string][]string)
}

func (m MultiMap) Get(key string) (value string, found bool) {
	values, found := m[key]

	if !found {
		return "", found
	}

	return values[0], found
}

func (m MultiMap) GetAll(key string) (values []string, found bool) {
	values, found = m[key]

	return values, found
}

func (m *MultiMap) Put(key string, value string) {
	values, found := (*m)[key]

	if !found {
		(*m)[key] = []string{value}
	} else {
		(*m)[key] = append(values, value)
	}
}

func (m *MultiMap) PutAll(key string, values ...string) {
	curValues, found := (*m)[key]

	if !found {
		(*m)[key] = values
	} else {
		(*m)[key] = append(curValues, values...)
	}
}

type GffRecord struct {
	Seqname    string
	Source     string
	Type       string
	Start      int32
	End        int32
	Score      string
	Strand     string
	Phase      string
	Attributes MultiMap
}

func NewGffRecord(seqname string, source string, type_ string, start int32, end int32, score string, strand string, phase string, attrs MultiMap) *GffRecord {
	record := new(GffRecord)
	record.Seqname = seqname
	record.Source = source
	record.Type = type_
	record.Start = start
	record.End = end
	record.Score = score
	record.Strand = strand
	record.Phase = phase
	record.Attributes = attrs
	return record
}

func ParseGffAttrs(attrs_string string) (MultiMap, error) {
	multimap := NewMultiMap()
	attrs_string = strings.ReplaceAll(attrs_string, "\"", "")

	for _, kv_pair := range strings.Split(attrs_string, ";") {
		kv_arr := strings.Split(kv_pair, "=")
		if len(kv_arr) != 2 {
			return multimap, fmt.Errorf("attrs should be separated by =: %v\n", attrs_string)
		}
		k := kv_arr[0]
		values := kv_arr[1]

		multimap.PutAll(k, strings.Split(values, ",")...)
	}

	return multimap, nil
}

func (r GffRecord) IsGeneChild() bool {
	return r.Type == "mRNA" || r.Type == "tRNA" || r.Type == "rRNA" || r.Type == "miRNA" || r.Type == "pre-miRNA"
}

func (r GffRecord) IsExon() bool {
	return r.Type == "exon"
}

func (r GffRecord) IsCds() bool {
	return r.Type == "CDS"
}

func (r GffRecord) IsFivePrimeUtr() bool {
	/*
		[Term]
		id: SO:0000204
		name: five_prime_UTR
		namespace: sequence
		def: "A region at the 5' end of a mature transcript (preceding the initiation codon) that is not translated into a protein." [http://www.insdc.org/files/feature_table.html]
		subset: SOFA
		synonym: "5' UTR" EXACT []
		synonym: "five prime UTR" EXACT []
		synonym: "five_prime_untranslated_region" EXACT []
		synonym: "INSDC_feature:5'UTR" EXACT []
		xref: http://en.wikipedia.org/wiki/5'_UTR "wiki"
		is_a: SO:0000203 ! UTR
	*/
	return r.Type == "five_prime_UTR"
}

func (r GffRecord) IsThreePrimeUtr() bool {
	/*
		[Term]
		id: SO:0000205
		name: three_prime_UTR
		namespace: sequence
		def: "A region at the 3' end of a mature transcript (following the stop codon) that is not translated into a protein." [http://www.insdc.org/files/feature_table.html]
		subset: SOFA
		synonym: "INSDC_feature:3'UTR" EXACT []
		synonym: "three prime untranslated region" EXACT []
		synonym: "three prime UTR" EXACT []
		xref: http://en.wikipedia.org/wiki/Three_prime_untranslated_region "wiki"
		is_a: SO:0000203 ! UTR
	*/
	return r.Type == "three_prime_UTR"
}

func (r GffRecord) ExtractSequence(seqnameSeq string) string {
	sequence := seqnameSeq[r.Start:r.End]

	if r.Strand == "-" {
		sequence = seq.ReverseComplement(sequence)
	}

	return sequence
}

type GffParser struct {
	Version  string
	Comments []string
	Records  []GffRecord
}

func NewGffParser() *GffParser {
	return &GffParser{Version: "gff3"}
}

func (p *GffParser) ConsumeLine(line string) error {
	line = strings.TrimSpace(line)
	if strings.HasPrefix(line, "#") {
		p.Comments = append(p.Comments, line)
		return nil
	}

	if line == "" {
		return nil
	}

	recs := strings.Split(line, "\t")

	if len(recs) != 9 {
		return fmt.Errorf("Incorrect line, expected 9 records but %d records: %v", len(recs), line)
	}

	seqname := recs[0]
	source := recs[1]
	type_ := recs[2]
	start, err := strconv.Atoi(recs[3])

	if err != nil {
		return fmt.Errorf("Incorrect line, start should be integer: %v", line)
	}

	end, err := strconv.Atoi(recs[4])

	if err != nil {
		return fmt.Errorf("Incorrect line, end should be integer: %v", line)
	}

	score := recs[5]
	strand := recs[6]

	if !(strand == "-" || strand == "+") {
		return fmt.Errorf("Incorrect line, strand must be - or + but %v: %v", strand, line)
	}

	phase := recs[7]
	attrs, err := ParseGffAttrs(recs[8])

	if err != nil {
		return fmt.Errorf("Errer: %v in %v", err, line)
	}

	gffRecord := NewGffRecord(seqname, source, type_, int32(start), int32(end), score, strand, phase, attrs)
	p.Records = append(p.Records, *gffRecord)

	return nil
}
