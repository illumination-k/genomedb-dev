package bioio_test

import (
	"genomedb/bioio"
	"strings"
	"testing"
)

func TestFastaIOBasic(t *testing.T) {
	fasta := `>A
ATGC
>B
ATGC
CGTA
>C test
ATGCCGTA
ATTTTTTT
GGGGGGGG
`
	parser := bioio.NewFastaParser()

	for _, line := range strings.Split(fasta, "\n") {
		parser.ConsumeLine(line)
	}

	parser.Flush()

	expected := []bioio.FastaRecord{
		*bioio.NewFastaRecord("A", "", "ATGC"),
		*bioio.NewFastaRecord("B", "", "ATGCCGTA"),
		*bioio.NewFastaRecord("C", "test", "ATGCCGTAATTTTTTTGGGGGGGG"),
	}

	for i, e := range expected {
		t.Run("Basic io", func(t *testing.T) {
			if e != parser.Records[i] {
				t.Errorf("error in record %d\n expected: %v\n actual: %v", i, e, parser.Records[i])
			}
		})
	}
}
