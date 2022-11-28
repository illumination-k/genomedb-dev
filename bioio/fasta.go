package bioio

import (
	"strings"
)

type FastaRecord struct {
	id   string
	name string
	seq  string
}

func NewFastaRecord(id string, name string, seq string) *FastaRecord {
	record := new(FastaRecord)
	record.id = id
	record.name = name
	record.seq = seq

	return record
}

type FastaParser struct {
	curId   string
	curName string
	curSeq  string
	Records []FastaRecord
}

func NewFastaParser() *FastaParser {
	parser := new(FastaParser)
	return parser
}

func (p FastaParser) isEmpty() bool {
	return p.curId == "" && p.curName == "" && p.curSeq == ""
}

func (p *FastaParser) ConsumeLine(line string) {
	line = strings.TrimSpace(line)

	if line == "" {
		return
	}

	if strings.HasPrefix(line, ">") {
		if p.curId != "" {
			var record *FastaRecord = NewFastaRecord(p.curId, p.curName, p.curSeq)
			p.Records = append(p.Records, *record)
		}

		attrs := strings.Split(line, " ")
		p.curId = strings.ReplaceAll(attrs[0], ">", "")

		if len(attrs) > 1 {
			p.curName = attrs[1]
		}

		p.curSeq = ""
	} else {
		p.curSeq = p.curSeq + line
	}
}

func (p *FastaParser) Flush() {
	if p.isEmpty() {
		return
	}

	var record *FastaRecord = NewFastaRecord(p.curId, p.curName, p.curSeq)
	p.Records = append(p.Records, *record)
}
