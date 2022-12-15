package seqio_test

import (
	"genomedb/seqio"
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
	parser := seqio.NewFastaParser()

	for _, line := range strings.Split(fasta, "\n") {
		parser.ConsumeLine(line)
	}

	parser.Flush()

	expected := []seqio.FastaRecord{
		*seqio.NewFastaRecord("A", "", "ATGC"),
		*seqio.NewFastaRecord("B", "", "ATGCCGTA"),
		*seqio.NewFastaRecord("C", "test", "ATGCCGTAATTTTTTTGGGGGGGG"),
	}

	for i, e := range expected {
		t.Run("Basic io", func(t *testing.T) {
			if e != parser.Records[i] {
				t.Errorf("error in record %d\n expected: %v\n actual: %v", i, e, parser.Records[i])
			}
		})
	}
}
