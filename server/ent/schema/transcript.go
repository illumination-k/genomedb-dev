package schema

import (
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
		field.String("strand"),
		field.String("type"),
		field.Text("genome_seq"),
		field.Text("transcript_seq"),
		field.Text("cds_seq"),
		field.Text("protein_seq"),
	}
}

// Edges of the Transcript.
func (Transcript) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("gene", Gene.Type).Ref("transcripts").Unique(),
		edge.To("cds", Cds.Type),
		edge.To("exon", Exon.Type),
		edge.To("five_prime_utr", FivePrimeUtr.Type),
		edge.To("three_prime_utr", ThreePrimeUtr.Type),
	}
}
