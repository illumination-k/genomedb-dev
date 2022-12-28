package seq

import (
	"fmt"
	"strings"
)

// https://www.ncbi.nlm.nih.gov/Taxonomy/taxonomyhome.html/index.cgi?chapter=tgencodes
const STANDART_CODE = 1

type CodonTable struct {
	id          int32
	name        string
	description string
	table       map[string]string
}

func (c CodonTable) Transrate(seq string) (string, error) {
	var err error
	seq = strings.ToUpper(seq)

	var prot strings.Builder

	for i := 0; i < len(seq); i += 3 {
		if i+3 > len(seq) {
			err = fmt.Errorf("Sequence should be 3x")
			continue
		}

		codon := seq[i : i+3]

		aminoAcid, ok := c.table[codon]

		if ok {
			prot.WriteString(aminoAcid)
		} else {
			prot.WriteString(Unknown)
		}
	}

	return prot.String(), err
}

func (c *CodonTable) fromText(aa string, base1 string, base2 string, base3 string) error {
	if !(len(aa) == 64 && len(base1) == 64 && len(base2) == 64 && len(base3) == 64) {
		return fmt.Errorf("All of aa, base1, base2 and base3 Should be 64 base.")
	}

	table := make(map[string]string, 64)

	for i := 0; i < 64; i++ {
		var codon strings.Builder
		codon.Grow(3)
		codon.WriteByte(base1[i])
		codon.WriteByte(base2[i])
		codon.WriteByte(base3[i])

		table[codon.String()] = string(aa[i])
	}

	c.table = table
	return nil
}

func NewCodonTable(id int32) CodonTable {
	var codon_table CodonTable

	switch id {
	case STANDART_CODE:
		codon_table = CodonTable{id: STANDART_CODE, name: "The Standard Code",
			description: `By default all transl_table in GenBank flatfiles are equal to id 1, and this is not shown. When transl_table is not equal to id 1, it is shown as a qualifier on the CDS feature.`}

		codon_table.fromText(
			"FFLLSSSSYY**CC*WLLLLPPPPHHQQRRRRIIIMTTTTNNKKSSRRVVVVAAAADDEEGGGG",
			"TTTTTTTTTTTTTTTTCCCCCCCCCCCCCCCCAAAAAAAAAAAAAAAAGGGGGGGGGGGGGGGG",
			"TTTTCCCCAAAAGGGGTTTTCCCCAAAAGGGGTTTTCCCCAAAAGGGGTTTTCCCCAAAAGGGG",
			"TCAGTCAGTCAGTCAGTCAGTCAGTCAGTCAGTCAGTCAGTCAGTCAGTCAGTCAGTCAGTCAG",
		)
	}

	return codon_table
}
