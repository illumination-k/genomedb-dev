package schema

import (
	"entgo.io/contrib/entproto"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Cds holds the schema definition for the Cds entity.
type Cds struct {
	ent.Schema
}

// Fields of the Cds.
func (Cds) Fields() []ent.Field {
	return []ent.Field{
		field.String("seqname").Annotations(entproto.Field(2)),
		field.Int32("start").Positive().Annotations(entproto.Field(3)),
		field.Int32("end").Positive().Annotations(entproto.Field(4)),
		field.Int8("phase").Annotations(entproto.Field(5)),
		field.String("strand").Annotations(entproto.Field(6)),
	}
}

// Edges of the Cds.
func (Cds) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("transcript", Transcript.Type).Ref("cds").Unique().Annotations(entproto.Field(7)),
	}
}

func (Cds) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entproto.Message(),
	}
}
