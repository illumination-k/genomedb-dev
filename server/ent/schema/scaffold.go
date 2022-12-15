package schema

import (
	"entgo.io/contrib/entproto"
	"entgo.io/ent"
	"entgo.io/ent/schema"
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
		field.String("seqname").Annotations(entproto.Field(2)),
		field.Text("seq").Annotations(entproto.Field(3)),
	}
}

// Edges of the Scaffold.
func (Scaffold) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("genome", Genome.Type).Ref("scaffolds").Unique().Annotations(entproto.Field(4)),
	}
}

// gRPC annotation
func (Scaffold) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entproto.Message(),
	}
}
