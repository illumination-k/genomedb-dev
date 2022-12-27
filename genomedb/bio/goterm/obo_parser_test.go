package goterm_test

import (
	"genomedb/bio/goterm"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOboParserBasic(t *testing.T) {
	example_go_obo := `
format-version: 1.2
data-version: releases/2022-12-04
subsetdef: goslim_synapse "synapse GO slim"
subsetdef: goslim_yeast "Yeast GO slim"
synonymtypedef: syngo_official_label "label approved by the SynGO project"
synonymtypedef: systematic_synonym "Systematic synonym" EXACT

[Term]
id: GO:0000001
name: mitochondrion inheritance
namespace: biological_process
def: "The distribution of mitochondria, including the mitochondrial genome, into daughter cells after mitosis or meiosis, mediated by interactions between mitochondria and the cytoskeleton." [GOC:mcc, PMID:10873824, PMID:11389764]
synonym: "mitochondrial inheritance" EXACT []
is_a: GO:0048308 ! organelle inheritance
is_a: GO:0048311 ! mitochondrion distribution

[Term]
id: GO:0000009
name: alpha-1,6-mannosyltransferase activity
namespace: molecular_function
def: "Catalysis of the transfer of a mannose residue to an oligosaccharide, forming an alpha-(1->6) linkage." [GOC:mcc, PMID:2644248]
synonym: "1,6-alpha-mannosyltransferase activity" EXACT []
xref: EC:2.4.1.232
xref: Reactome:R-HSA-449718 "Addition of a third mannose to the N-glycan precursor by ALG2"
is_a: GO:0000030 ! mannosyltransferase activity

[Term]
id: GO:0000039
name: obsolete plasma membrane long-chain fatty acid transporter
namespace: molecular_function
def: "OBSOLETE. (Was not defined before being made obsolete)." [GOC:ai]
comment: This term was made obsolete because it describes a gene product and it contains component information.
synonym: "plasma membrane long-chain fatty acid transporter" EXACT []
is_obsolete: true
consider: GO:0005324
consider: GO:0005886

[Term]
id: GO:2001286
name: regulation of caveolin-mediated endocytosis
namespace: biological_process
def: "Any process that modulates the frequency, rate or extent of caveolin-mediated endocytosis." [GOC:obol]
synonym: "regulation of caveolae-dependent endocytosis" EXACT [GOC:obol]
synonym: "regulation of caveolae-mediated endocytosis" EXACT [GOC:obol]
synonym: "regulation of caveolin-dependent endocytosis" EXACT [GOC:obol]
is_a: GO:0030100 ! regulation of endocytosis
relationship: regulates GO:0072584 ! caveolin-mediated endocytosis

[Typedef]
id: negatively_regulates
name: negatively regulates
namespace: external
xref: RO:0002212
is_a: regulates ! regulates
`
	var err error
	parser := *goterm.NewOboParser()

	for _, line := range strings.Split(example_go_obo, "\n") {
		if err = parser.ConsumeLine(line); err != nil {
			t.Errorf("Error in consumeline: %v", err)
		}
	}

	t.Run("Check Header", func(t *testing.T) {
		assert.Equal(t, parser.Header.Version, "1.2")
		assert.Equal(t, parser.Header.DataVersion, "releases/2022-12-04")
	})

	t.Run("Check goterms", func(t *testing.T) {
		expected := []goterm.GoTerm{
			{
				Id:            "GO:0000001",
				Name:          "mitochondrion inheritance",
				Namespace:     "BP",
				Def:           "\"The distribution of mitochondria, including the mitochondrial genome, into daughter cells after mitosis or meiosis, mediated by interactions between mitochondria and the cytoskeleton.\" [GOC:mcc, PMID:10873824, PMID:11389764]",
				IsObsolete:    false,
				Xrefs:         []string(nil),
				Synonyms:      []string{"\"mitochondrial inheritance\" EXACT []"},
				IsAs:          []string{"GO:0048308", "GO:0048311"},
				Relationships: []goterm.Relationship(nil),
			},
			{
				Id:            "GO:0000009",
				Name:          "alpha-1,6-mannosyltransferase activity",
				Namespace:     "MF",
				Def:           "\"Catalysis of the transfer of a mannose residue to an oligosaccharide, forming an alpha-(1->6) linkage.\" [GOC:mcc, PMID:2644248]",
				IsObsolete:    false,
				Xrefs:         []string{"EC:2.4.1.232", "Reactome:R-HSA-449718 \"Addition of a third mannose to the N-glycan precursor by ALG2\""},
				Synonyms:      []string{"\"1,6-alpha-mannosyltransferase activity\" EXACT []"},
				IsAs:          []string{"GO:0000030"},
				Relationships: []goterm.Relationship(nil),
			},
			{
				Id:         "GO:0000039",
				Name:       "obsolete plasma membrane long-chain fatty acid transporter",
				Namespace:  "MF",
				Def:        "\"OBSOLETE. (Was not defined before being made obsolete).\" [GOC:ai]",
				IsObsolete: true,
				Xrefs:      []string(nil),
				Synonyms:   []string{"\"plasma membrane long-chain fatty acid transporter\" EXACT []"},
				IsAs:       []string(nil), Relationships: []goterm.Relationship(nil),
			},
			{
				Id:            "GO:2001286",
				Name:          "regulation of caveolin-mediated endocytosis",
				Namespace:     "BP",
				Def:           "\"Any process that modulates the frequency, rate or extent of caveolin-mediated endocytosis.\" [GOC:obol]",
				IsObsolete:    false,
				Xrefs:         []string(nil),
				Synonyms:      []string{"\"regulation of caveolae-dependent endocytosis\" EXACT [GOC:obol]", "\"regulation of caveolae-mediated endocytosis\" EXACT [GOC:obol]", "\"regulation of caveolin-dependent endocytosis\" EXACT [GOC:obol]"},
				IsAs:          []string{"GO:0030100"},
				Relationships: []goterm.Relationship{{Name: "regulates", Target: "GO:0072584"}},
			},
		}

		assert.EqualValues(t, parser.GoTerms, expected)
	})
}
