package interproscan

import (
	"fmt"
	"strconv"
	"strings"
)

type InterproscanParser struct{}

func splitTerms(attr string) []string {
	var arr []string
	if attr == "" {
		return arr
	}

	return strings.Split("|", attr)
}

func (p InterproscanParser) ConsumeLine(line string) (record InterproscanRecord, err error) {
	line = strings.TrimSpace(line)

	attrs := strings.Split(line, "\t")

	if len(attrs) < 11 {
		return record, fmt.Errorf("Invalid line: %s\n", line)
	}

	for {
		if len(attrs) == 15 {
			break
		}

		attrs = append(attrs, "")
	}

	record.Accession = attrs[0]

	if record.Length, err = strconv.Atoi(attrs[2]); err != nil {
		return record, fmt.Errorf("Invalid line, length field should be integer: %s\n", line)
	}

	record.Analysis = attrs[3]

	if record.Start, err = strconv.Atoi(attrs[6]); err != nil {
		return record, fmt.Errorf("Invalid line, start field should be integer: %s\n", line)
	}

	if record.Stop, err = strconv.Atoi(attrs[7]); err != nil {
		return record, fmt.Errorf("Invalid line, end field should be integer: %s\n", line)
	}

	if record.Score, err = strconv.ParseFloat(attrs[8], 64); err != nil {
		return record, fmt.Errorf("Invalid line, end field should be float: %s\n", line)
	}

	record.SignatureAccession = attrs[11]
	record.SignatureDescription = attrs[12]
	record.GoTermIDs = splitTerms(attrs[13])
	record.PathwayIDs = splitTerms(attrs[14])

	return record, nil
}
