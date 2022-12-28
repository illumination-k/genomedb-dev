package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// GoTerm holds the schema definition for the GoTerm entity.
type GoTerm struct {
	ent.Schema
}

// Fields of the GoTerm.
func (GoTerm) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").Comment("Go term id: GO:NNNNNN"),
		field.Enum("namespace").Values("BP", "MF", "CC"),
		field.String("name").Comment("short name of GO term"),
		field.String("def").Comment("Go term description"),
		field.Int32("level").Comment("shortest distance from root node"),
		field.Int32("depth").Comment("longest distance from root node"),
	}
}

// Edges of the GoTerm.
func (GoTerm) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("children", GoTerm.Type).
			From("parent").
			Unique(),
		edge.To("transcripts", Transcript.Type).Through("goterm_transcript", GoTermOnTranscripts.Type),
	}
}

// GoTermOnTranscripts
type GoTermOnTranscripts struct {
	ent.Schema
}

func (GoTermOnTranscripts) Fields() []ent.Field {
	return []ent.Field{
		field.String("evidence_code"),
		field.String("go_term_id"),
		field.String("transcript_id"),
	}
}

func (GoTermOnTranscripts) Annotations() []schema.Annotation {
	return []schema.Annotation{
		field.ID("go_term_id", "transcript_id"),
	}
}

func (GoTermOnTranscripts) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("go_term", GoTerm.Type).
			Unique().
			Required().
			Field("go_term_id"),
		edge.To("transcript", Transcript.Type).
			Required().
			Unique().
			Field("transcript_id"),
	}
}
