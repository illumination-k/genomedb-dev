package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type DomainAnnotation struct {
	ent.Schema
}

// https://interproscan-docs.readthedocs.io/en/latest/HowToRun.html

func (DomainAnnotation) Fields() []ent.Field {
	return []ent.Field{
		field.String("id"),
		field.String("description"),
		field.String("Analysis").Comment(
			"CDD" +
				"COILS" +
				"Gene3D" +
				"HAMAP" +
				"MOBIDB" +
				"PANTHER" +
				"Pfam" +
				"PIRSF" +
				"PRINTS" +
				"PROSITE" +
				"SFLD" +
				"SMART" +
				"SUPERFAMILY" +
				"TIGRFAMs",
		),
	}
}

func (DomainAnnotation) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("transcripts", Transcript.Type).Through("domain_transcript", DomainAnnotationToTranscript.Type),
	}
}

type DomainAnnotationToTranscript struct {
	ent.Schema
}

func (DomainAnnotationToTranscript) Fields() []ent.Field {
	return []ent.Field{
		field.String("domain_annotation_id"),
		field.String("transcript_id"),
		field.Int32("start").Positive(),
		field.Int32("stop").Positive(),
		field.Float("score"),
	}
}

func (DomainAnnotationToTranscript) Annotations() []schema.Annotation {
	return []schema.Annotation{
		field.ID("domain_annotation_id", "transcript_id"),
	}
}

func (DomainAnnotationToTranscript) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("domain", DomainAnnotation.Type).
			Unique().
			Required().
			Field("domain_annotation_id"),
		edge.To("transcript", Transcript.Type).
			Required().
			Unique().
			Field("transcript_id"),
	}
}
