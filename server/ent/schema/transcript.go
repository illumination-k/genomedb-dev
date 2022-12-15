package schema

import (
	"entgo.io/contrib/entproto"
	"entgo.io/ent"
	"entgo.io/ent/schema"
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
		field.String("id").Annotations(entproto.Field(1)),
		field.String("strand").Annotations(entproto.Field(2)),
		field.String("type").Annotations(entproto.Field(3)),
		field.Text("genome_seq").Annotations(entproto.Field(4)),
		field.Text("transcript_seq").Annotations(entproto.Field(5)),
		field.Text("cds_seq").Annotations(entproto.Field(6)),
		field.Text("protein_seq").Annotations(entproto.Field(7)),
	}
}

// Edges of the Transcript.
func (Transcript) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("gene", Gene.Type).Ref("transcripts").Unique().Annotations(entproto.Field(8)),
		edge.To("cds", Cds.Type).Annotations(entproto.Field(9)),
		edge.To("exon", Exon.Type).Annotations(entproto.Field(10)),
		edge.To("five_prime_utr", FivePrimeUtr.Type).Annotations(entproto.Field(11)),
		edge.To("three_prime_utr", ThreePrimeUtr.Type).Annotations(entproto.Field(12)),
	}
}

// gRPC annotation
func (Transcript) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entproto.Message(),
		entproto.Service(),
	}
}
