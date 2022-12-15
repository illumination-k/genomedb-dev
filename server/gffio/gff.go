// https://github.com/The-Sequence-Ontology/Specifications/blob/master/gff3.md

package gffio

import (
	"fmt"
	"strconv"
	"strings"
)

type MultiMap struct {
	m map[string][]string
}

func NewMultiMap() MultiMap {
	return MultiMap{make(map[string][]string)}
}

func (m MultiMap) Get(key string) (value string, found bool) {
	values, found := m.m[key]

	if !found {
		return "", found
	}

	return values[0], found
}

func (m MultiMap) GetAll(key string) (values []string, found bool) {
	values, found = m.m[key]

	return values, found
}

func (m *MultiMap) Put(key string, value string) {
	values, found := m.m[key]

	if !found {
		m.m[key] = []string{value}
	} else {
		m.m[key] = append(values, value)
	}
}

func (m *MultiMap) PutAll(key string, values ...string) {
	curValues, found := m.m[key]

	if !found {
		m.m[key] = values
	} else {
		m.m[key] = append(curValues, values...)
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
	return r.Type == "five_prime_UTR"
}

func (r GffRecord) IsThreePrimeUtr() bool {
	return r.Type == "three_prime_UTR"
}

type GffParser struct {
	Version string
	Records []GffRecord
}

func NewGffParser() *GffParser {
	return &GffParser{Version: "gff3"}
}

func (p *GffParser) ConsumeLine(line string) error {
	line = strings.TrimSpace(line)
	if strings.HasPrefix(line, "#") {
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
