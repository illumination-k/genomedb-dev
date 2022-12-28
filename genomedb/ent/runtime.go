// Code generated by ent, DO NOT EDIT.

package ent

import (
	"genomedb/ent/domainannotationtotranscript"
	"genomedb/ent/genome"
	"genomedb/ent/keggorthlogy"
	"genomedb/ent/schema"
	"genomedb/ent/transcript"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	domainannotationtotranscriptFields := schema.DomainAnnotationToTranscript{}.Fields()
	_ = domainannotationtotranscriptFields
	// domainannotationtotranscriptDescStart is the schema descriptor for start field.
	domainannotationtotranscriptDescStart := domainannotationtotranscriptFields[2].Descriptor()
	// domainannotationtotranscript.StartValidator is a validator for the "start" field. It is called by the builders before save.
	domainannotationtotranscript.StartValidator = domainannotationtotranscriptDescStart.Validators[0].(func(int32) error)
	// domainannotationtotranscriptDescStop is the schema descriptor for stop field.
	domainannotationtotranscriptDescStop := domainannotationtotranscriptFields[3].Descriptor()
	// domainannotationtotranscript.StopValidator is a validator for the "stop" field. It is called by the builders before save.
	domainannotationtotranscript.StopValidator = domainannotationtotranscriptDescStop.Validators[0].(func(int32) error)
	genomeFields := schema.Genome{}.Fields()
	_ = genomeFields
	// genomeDescCodonTable is the schema descriptor for codon_table field.
	genomeDescCodonTable := genomeFields[1].Descriptor()
	// genome.CodonTableValidator is a validator for the "codon_table" field. It is called by the builders before save.
	genome.CodonTableValidator = func() func(int32) error {
		validators := genomeDescCodonTable.Validators
		fns := [...]func(int32) error{
			validators[0].(func(int32) error),
			validators[1].(func(int32) error),
		}
		return func(codon_table int32) error {
			for _, fn := range fns {
				if err := fn(codon_table); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	keggorthlogyFields := schema.KeggOrthlogy{}.Fields()
	_ = keggorthlogyFields
	// keggorthlogyDescID is the schema descriptor for id field.
	keggorthlogyDescID := keggorthlogyFields[0].Descriptor()
	// keggorthlogy.IDValidator is a validator for the "id" field. It is called by the builders before save.
	keggorthlogy.IDValidator = keggorthlogyDescID.Validators[0].(func(string) error)
	transcriptFields := schema.Transcript{}.Fields()
	_ = transcriptFields
	// transcriptDescStart is the schema descriptor for start field.
	transcriptDescStart := transcriptFields[5].Descriptor()
	// transcript.StartValidator is a validator for the "start" field. It is called by the builders before save.
	transcript.StartValidator = transcriptDescStart.Validators[0].(func(int32) error)
	// transcriptDescEnd is the schema descriptor for end field.
	transcriptDescEnd := transcriptFields[6].Descriptor()
	// transcript.EndValidator is a validator for the "end" field. It is called by the builders before save.
	transcript.EndValidator = transcriptDescEnd.Validators[0].(func(int32) error)
}
