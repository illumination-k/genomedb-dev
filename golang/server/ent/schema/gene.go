package schema

import (
	"entgo.io/contrib/entproto"
	"entgo.io/ent"
	"entgo.io/ent/schema"
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
		field.String("id").StorageKey("geneId").Annotations(entproto.Field(1)),
	}
}

// Edges of the Gene.
func (Gene) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("transcripts", Transcript.Type).Annotations(entproto.Field(2)),
		edge.From("genome", Genome.Type).Ref("genes").Unique().Annotations(entproto.Field(3)),
	}
}

// gRPC annotation
func (Gene) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entproto.Message(), entproto.Service(),
	}
}
