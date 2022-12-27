package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Locus holds the schema definition for the Locus entity.
type Locus struct {
	ent.Schema
}

// Fields of the Locus.
func (Locus) Fields() []ent.Field {
	return []ent.Field{
		field.String("id"),
	}
}

// Edges of the Locus.
func (Locus) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("transcripts", Transcript.Type),
		edge.From("genome", Genome.Type).Ref("locuses").Unique(),
	}
}
