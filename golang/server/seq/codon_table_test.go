package seq_test

import (
	"fmt"
	"genomedb/seq"
	"testing"
)

func TestStandardCodon(t *testing.T) {
	codon_table := seq.NewCodonTable(seq.STANDART_CODE)
	fmt.Println(codon_table)

	nuc := `atgacgacgacagctgcggaaagtgatccccgcctcgcccacatcaaggagtcgattcgtaccatccccggtttcccaaaggccggtatccttttccgcgatgtgacgactcttctgcttgatcacaaagcttttaaggacaccactgacatttttgtcgaacgatacaaggacatgaaaatagacgtcgttgttggaattgaagcacgaggtttcatatttggcccaccgattgcactcgctattggagcgaaatttgttcctcaacgaaagcccaagaagctaccaggccctgttgtcggtgaggagtacgagttggagtacggtacagaccgaattgagatgcacgttggagccgtgcagcccggtgaacgctgtttgattgtggacgatcttattgcaaccgggggtacacttgctgcaggaataaggctactagagcgcgtcggtgccgaggttgtggagtgtgcctgcttgatcgagctttcggatctgaagggccgtgagaaactgaaaggcaagcctctgttcatcctggtagacttcgagggagagtaa`
	prot := `MTTTAAESDPRLAHIKESIRTIPGFPKAGILFRDVTTLLLDHKAFKDTTDIFVERYKDMKIDVVVGIEARGFIFGPPIALAIGAKFVPQRKPKKLPGPVVGEEYELEYGTDRIEMHVGAVQPGERCLIVDDLIATGGTLAAGIRLLERVGAEVVECACLIELSDLKGREKLKGKPLFILVDFEGE*`

	t.Run("Basic translation", func(t *testing.T) {
		actual, err := codon_table.Transrate(nuc)

		if err != nil {
			t.Error(err)
		}

		if prot != actual {
			t.Errorf("Translation failed\nexpected: %v\nactual  : %v", prot, actual)
		}
	})
}
