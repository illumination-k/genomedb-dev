package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// FivePrimeUtr holds the schema definition for the FivePrimeUtr entity.
type FivePrimeUtr struct {
	ent.Schema
}

// Fields of the FivePrimeUtr.
func (FivePrimeUtr) Fields() []ent.Field {
	return []ent.Field{
		field.String("seqname"),
		field.Int32("start").Positive(),
		field.Int32("end").Positive(),
		field.String("strand"),
	}
}

// Edges of the FivePrimeUtr.
func (FivePrimeUtr) Edges() []ent.Edge {
	return []ent.Edge{edge.From("transcript", Transcript.Type).Ref("five_prime_utr").Unique()}
}
