package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Gene holds the schema definition for the Gene entity.
type Gene struct {
	ent.Schema
}

// Fields of the Gene.
func (Gene) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").StorageKey("geneId"),
	}
}

// Edges of the Gene.
func (Gene) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("transcripts", Transcript.Type),
		edge.From("genome", Genome.Type).Ref("genes").Unique(),
	}
}
