package schema

import (
	"entgo.io/ent"
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
		field.String("seqname"),
		field.Int32("start").Positive(),
		field.Int32("end").Positive(),
		field.String("strand"),
	}
}

// Edges of the ThreePrimeUtr.
func (ThreePrimeUtr) Edges() []ent.Edge {
	return []ent.Edge{edge.From("transcript", Transcript.Type).Ref("three_prime_utr").Unique()}
}
