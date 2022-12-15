package seqio

import (
	"fmt"
	"strings"
)

type FastaRecord struct {
	Id   string
	Name string
	Seq  string
}

func NewFastaRecord(id string, name string, seq string) *FastaRecord {
	record := new(FastaRecord)
	record.Id = id
	record.Name = name
	record.Seq = seq

	return record
}

type FastaParser struct {
	curLineNumber int
	curId         string
	curName       string
	curSeq        strings.Builder
	Records       []FastaRecord
}

func NewFastaParser() *FastaParser {
	parser := new(FastaParser)
	return parser
}

func (p FastaParser) isEmpty() bool {
	return p.curId == "" && p.curName == "" && p.curSeq.String() == ""
}

func (p *FastaParser) ConsumeLine(line string) error {
	line = strings.TrimSpace(line)

	if line == "" {
		return nil
	}

	if strings.HasPrefix(line, ">") {
		if p.curId != "" {
			var record *FastaRecord = NewFastaRecord(p.curId, p.curName, p.curSeq.String())
			p.Records = append(p.Records, *record)
		}

		attrs := strings.Split(line, " ")
		p.curId = strings.ReplaceAll(attrs[0], ">", "")

		if p.curId == "" {
			return fmt.Errorf("Empty id in the fasta file in line %d", p.curLineNumber)
		}

		if len(attrs) > 1 {
			p.curName = attrs[1]
		}

		p.curSeq.Reset()
	} else {
		p.curSeq.WriteString(line)
	}

	p.curLineNumber++

	return nil
}

func (p *FastaParser) Flush() {
	if p.isEmpty() {
		return
	}

	var record *FastaRecord = NewFastaRecord(p.curId, p.curName, p.curSeq.String())
	p.Records = append(p.Records, *record)
}
