package interproscan_test

import (
	"genomedb/bio/interproscan"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInterproscanTsvParser(t *testing.T) {
	example_tsv := "P51587\t14086411a2cdf1c4cba63020e1622579\t3418\tPfam\tPF09103\tBRCA2, oligonucleotide/oligosaccharide-binding, domain 1\t2670\t2799\t7.9E-43\tT\t15-03-2013\n" +
		"P51587\t14086411a2cdf1c4cba63020e1622579\t3418\tProSiteProfiles\tPS50138\tBRCA2 repeat profile.\t1002\t1036\t0.0\tT\t18-03-2013\tIPR002093\tBRCA2 repeat\tGO:0005515|GO:0006302\n" +
		"P51587\t14086411a2cdf1c4cba63020e1622579\t3418\tGene3D\tG3DSA:2.40.50.140\t\t2966\t3051\t3.1E-52\tT\t15-03-2013"

	parser := interproscan.InterproscanTsvParser{}

	for _, line := range strings.Split(example_tsv, "\n") {
		t.Run("consume line test", func(t *testing.T) {
			if err := parser.ConsumeLine(line); err != nil {
				t.Errorf("err %v in %s", err, line)
			}
		})
	}

	t.Run("Interproscan tsv parse result", func(t *testing.T) {
		expected := []interproscan.InterproscanRecord{
			{Accession: "P51587", Length: 3418, Start: 2670, Stop: 2799, Analysis: "Pfam", SignatureAccession: "PF09103", SignatureDescription: "BRCA2, oligonucleotide/oligosaccharide-binding, domain 1", Score: 7.9e-43, InterproscanAccession: "", InterproscanDescription: "", GoTermIDs: []string(nil), PathwayIDs: []string(nil)},
			{Accession: "P51587", Length: 3418, Start: 1002, Stop: 1036, Analysis: "ProSiteProfiles", SignatureAccession: "PS50138", SignatureDescription: "BRCA2 repeat profile.", Score: 0, InterproscanAccession: "IPR002093", InterproscanDescription: "BRCA2 repeat", GoTermIDs: []string{"GO:0005515", "GO:0006302"}, PathwayIDs: []string(nil)},
			{Accession: "P51587", Length: 3418, Start: 2966, Stop: 3051, Analysis: "Gene3D", SignatureAccession: "G3DSA:2.40.50.140", SignatureDescription: "", Score: 3.1e-52, InterproscanAccession: "", InterproscanDescription: "", GoTermIDs: []string(nil), PathwayIDs: []string(nil)},
		}

		assert.EqualValues(t, expected, parser.Records)
	})
}
