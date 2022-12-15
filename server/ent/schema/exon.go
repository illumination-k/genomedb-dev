package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Exon holds the schema definition for the Exon entity.
type Exon struct {
	ent.Schema
}

// Fields of the Exon.
func (Exon) Fields() []ent.Field {
	return []ent.Field{
		field.String("seqname"),
		field.Int32("start").Positive(),
		field.Int32("end").Positive(),
		field.String("strand"),
	}
}

// Edges of the Exon.
func (Exon) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("transcript", Transcript.Type).Ref("exon").Unique(),
	}
}
