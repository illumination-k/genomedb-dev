package schema

import (
	"entgo.io/contrib/entproto"
	"entgo.io/ent"
	"entgo.io/ent/schema"
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
		field.String("seqname").Annotations(entproto.Field(2)),
		field.Int32("start").Positive().Annotations(entproto.Field(3)),
		field.Int32("end").Positive().Annotations(entproto.Field(4)),
		field.String("strand").Annotations(entproto.Field(5)),
	}
}

// Edges of the Exon.
func (Exon) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("transcript", Transcript.Type).Ref("exon").Unique().Annotations(entproto.Field(6)),
	}
}

func (Exon) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entproto.Message(),
	}
}
