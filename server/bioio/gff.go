package bioio

import (
	"fmt"
	"strconv"
	"strings"
)

type GffRecord struct {
	seqname    string
	source     string
	feature    string
	start      int32
	end        int32
	score      string
	strand     string
	frame      string
	attributes map[string]string
}

func NewGffRecord(seqname string, source string, feature string, start int32, end int32, score string, strand string, frame string, attrs map[string]string) *GffRecord {
	record := new(GffRecord)
	record.seqname = seqname
	record.source = source
	record.feature = feature
	record.start = start
	record.end = end
	record.score = score
	record.strand = strand
	record.frame = frame
	record.attributes = attrs
	return record
}

type GffParser struct {
	Records []GffRecord
}

func NewGffParser() *GffParser {
	return &GffParser{}
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
	feature := recs[2]
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

	frame := recs[7]
	attrs := map[string]string{}

	for _, _kv := range strings.Split(recs[8], ";") {
		_kv = strings.TrimSpace(_kv)
		kv := strings.Split(_kv, "=")
		attrs[kv[0]] = kv[1]
	}

	gffRecord := NewGffRecord(seqname, source, feature, int32(start), int32(end), score, strand, frame, attrs)
	fmt.Println(gffRecord)
	p.Records = append(p.Records, *gffRecord)

	return nil
}