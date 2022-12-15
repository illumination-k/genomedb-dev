package schema

import (
	"entgo.io/ent"
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
		field.String("seqname"),
		field.Int32("start").Positive(),
		field.Int32("end").Positive(),
		field.Int8("frame"),
		field.String("strand"),
	}
}

// Edges of the Cds.
func (Cds) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("transcript", Transcript.Type).Ref("cds").Unique(),
	}
}
