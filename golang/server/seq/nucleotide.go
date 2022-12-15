package seq

import "strings"

func Reverse(s string) string {

	runes := []rune(s)

	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	return string(runes)
}

func ReverseComplement(seq string) string {
	seq = Reverse(strings.ToUpper(seq))

	compMap := map[rune]rune{
		'A': 'T',
		'T': 'A',
		'U': 'A',
		'C': 'G',
		'G': 'C',
	}

	var compSeq strings.Builder
	compSeq.Grow(len(seq))

	for _, nuc := range seq {
		compNuc, ok := compMap[nuc]

		if ok {
			compSeq.WriteRune(compNuc)
		} else {
			compSeq.WriteRune('N')
		}
	}

	return compSeq.String()
}
