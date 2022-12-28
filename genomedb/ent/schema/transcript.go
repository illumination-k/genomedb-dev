package schema

import (
	"genomedb/bio/gffio"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Transcript holds the schema definition for the Transcript entity.
type Transcript struct {
	ent.Schema
}

// Fields of the Transcript.
func (Transcript) Fields() []ent.Field {
	return []ent.Field{
		field.String("id"),
		field.String("seqname"),
		field.String("strand"),
		field.String("type"),
		field.String("source"),
		field.Int32("start").Positive(),
		field.Int32("end").Positive(),
		field.JSON("exon", []gffio.GffRecord{}),
		field.JSON("five_prime_utr", []gffio.GffRecord{}),
		field.JSON("three_prime_utr", []gffio.GffRecord{}),
		field.JSON("cds", []gffio.GffRecord{}),

		field.Text("genomic_sequence"),
		field.Text("exon_sequence"),
		field.Text("cds_sequence"),
		field.Text("protein_sequence"),
	}
}

// Edges of the Transcript.
func (Transcript) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("gene", Gene.Type).Ref("transcripts").Unique(),
		edge.From("goterms", GoTerm.Type).Ref("transcripts").Through("goterm_transcript", GoTermOnTranscripts.Type),
		edge.From("domains", DomainAnnotation.Type).Ref("transcripts").Through("domain_transcript", DomainAnnotationToTranscript.Type),
	}
}
