package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// KOG holds the schema definition for the KOG entity.
type KOG struct {
	ent.Schema
}

// Fields of the KOG.
func (KOG) Fields() []ent.Field {
	return []ent.Field{
		field.String("id"),
		field.String("description"),
		field.String("category"),
	}
}

// Edges of the KOG.
func (KOG) Edges() []ent.Edge {
	return nil
}
