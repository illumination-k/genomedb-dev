package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Scaffold holds the schema definition for the Scaffold entity.
type Scaffold struct {
	ent.Schema
}

// Fields of the Scaffold.
func (Scaffold) Fields() []ent.Field {
	return []ent.Field{
		field.String("seqname"),
		field.Text("seq"),
	}
}

// Edges of the Scaffold.
func (Scaffold) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("genome", Genome.Type).Ref("scaffolds").Unique(),
	}
}
