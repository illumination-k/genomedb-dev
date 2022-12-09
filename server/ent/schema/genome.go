package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Genome holds the schema definition for the Genome entity.
type Genome struct {
	ent.Schema
}

// Fields of the Genome.
func (Genome) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Unique(),
		field.Int32("codon_table"),
		field.Text("seq"),
	}
}

// Edges of the Genome.
func (Genome) Edges() []ent.Edge {
	return nil
}
