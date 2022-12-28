//go:build file

package goterm_test

import (
	"bufio"
	"genomedb/bio/goterm"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOboFileParse(t *testing.T) {
	r, err := os.Open("../../files/go-basic.obo")

	if err != nil {
		t.Error(err)
	}

	defer r.Close()

	br := bufio.NewScanner(r)

	parser := goterm.NewOboParser()

	for br.Scan() {
		if err := parser.ConsumeLine(br.Text()); err != nil {
			t.Errorf("%v", err)
		}
	}

	dag := goterm.CreateGoDAGFromGoterms(parser.GoTerms)

	t.Run("Test top (biological process, GO:0008150)", func(t *testing.T) {
		expected := goterm.GoDAGRecord{
			GoTermId: "GO:0008150",
			Depth:    0,
			Level:    0,
			Parent:   []string(nil),
			Children: []string{"GO:0000003", "GO:0002376", "GO:0008152", "GO:0009987", "GO:0016032", "GO:0022414", "GO:0023052", "GO:0032501", "GO:0032502", "GO:0040007", "GO:0040011", "GO:0042592", "GO:0043473", "GO:0044419", "GO:0044848", "GO:0048511", "GO:0050896", "GO:0051179", "GO:0051703", "GO:0065007", "GO:0098754"}}

		assert.Equal(t, dag["GO:0008150"], expected)
	})

	t.Run("Test middle, GO:0006807", func(t *testing.T) {
		expected := goterm.GoDAGRecord{
			GoTermId: "GO:0006807",
			Depth:    2,
			Level:    2,
			Parent:   []string{"GO:0008152"},
			Children: []string{"GO:0006653", "GO:0008291", "GO:0018887", "GO:0018916", "GO:0018937", "GO:0018952", "GO:0018954", "GO:0018960", "GO:0019326", "GO:0019329", "GO:0019330", "GO:0019666", "GO:0034641", "GO:0043603", "GO:0044597", "GO:0044598", "GO:0071941", "GO:0097164", "GO:1901119", "GO:1901150", "GO:1901564", "GO:1901742", "GO:1902061", "GO:2001057"},
		}

		assert.Equal(t, dag["GO:0006807"], expected)
	})

	t.Run("Test bottom, GO:2001315", func(t *testing.T) {
		expected := goterm.GoDAGRecord{
			GoTermId: "GO:2001315", Depth: 19, Level: 6, Parent: []string{"GO:0009226", "GO:0046349", "GO:2001313"}, Children: []string(nil),
		}
		assert.Equal(t, dag["GO:2001315"], expected)
	})
	// t.Errorf("Finish: %#v", dag["GO:2001315"])
}
