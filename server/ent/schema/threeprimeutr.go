package schema

import (
	"entgo.io/contrib/entproto"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// ThreePrimeUtr holds the schema definition for the ThreePrimeUtr entity.
type ThreePrimeUtr struct {
	ent.Schema
}

// Fields of the ThreePrimeUtr.
func (ThreePrimeUtr) Fields() []ent.Field {
	return []ent.Field{
		field.String("seqname").Annotations(entproto.Field(2)),
		field.Int32("start").Positive().Annotations(entproto.Field(3)),
		field.Int32("end").Positive().Annotations(entproto.Field(4)),
		field.String("strand").Annotations(entproto.Field(5)),
	}
}

// Edges of the ThreePrimeUtr.
func (ThreePrimeUtr) Edges() []ent.Edge {
	return []ent.Edge{edge.From("transcript", Transcript.Type).Ref("three_prime_utr").Unique().Annotations(entproto.Field(6))}
}

func (ThreePrimeUtr) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entproto.Message(),
	}
}
