package seq_test

import (
	"genomedb/bio/seq"
	"testing"
)

func TestReverseComplement(t *testing.T) {
	t.Run(
		"Basic Reverse", func(t *testing.T) {
			if seq.Reverse("ATGC") != "CGTA" {
				t.Fatalf("Reverse not work")
			}
		},
	)

	t.Run(
		"Basic DNA Reverse Complement", func(t *testing.T) {
			if seq.ReverseComplement("ATGCN") != "NGCAT" {
				t.Fatalf("DNA Reverse Complement not work")
			}
		},
	)
}
