package interproscan

import (
	"fmt"
	"strconv"
	"strings"
)

type InterproscanTsvParser struct {
	Records []InterproscanRecord
}

func splitTerms(attr string) []string {
	var arr []string
	if attr == "" {
		return arr
	}

	return strings.Split(attr, "|")
}

func (p *InterproscanTsvParser) ConsumeLine(line string) (err error) {
	var record InterproscanRecord
	line = strings.TrimSpace(line)

	if line == "" {
		return nil
	}

	attrs := strings.Split(line, "\t")

	if len(attrs) < 11 {
		return fmt.Errorf("Expected more than 11 records but only %d record.\nInvalid line: %s\n", len(attrs), line)
	}

	for {
		if len(attrs) == 15 {
			break
		}

		attrs = append(attrs, "")
	}

	record.Accession = attrs[0]

	if record.Length, err = strconv.Atoi(attrs[2]); err != nil {
		return fmt.Errorf("Invalid line, length field should be integer: %s\n", line)
	}

	record.Analysis = attrs[3]

	if record.Start, err = strconv.Atoi(attrs[6]); err != nil {
		return fmt.Errorf("Invalid line, start field should be integer: %s\n", line)
	}

	if record.Stop, err = strconv.Atoi(attrs[7]); err != nil {
		return fmt.Errorf("Invalid line, end field should be integer: %s\n", line)
	}

	if record.Score, err = strconv.ParseFloat(attrs[8], 64); err != nil {
		return fmt.Errorf("Invalid line, score field should be float: %s\n", line)
	}

	record.SignatureAccession = attrs[4]
	record.SignatureDescription = attrs[5]
	record.InterproscanAccession = attrs[11]
	record.InterproscanDescription = attrs[12]
	record.GoTermIDs = splitTerms(attrs[13])
	record.PathwayIDs = splitTerms(attrs[14])

	p.Records = append(p.Records, record)
	return nil
}
