package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Genome holds the schema definition for the Genome entity.
type Genome struct {
	ent.Schema
}

// Fields of the Genome.
func (Genome) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").StorageKey("name"),
		field.Int32("codon_table").
			Min(1).
			Max(31).
			Comment("See https://www.ncbi.nlm.nih.gov/Taxonomy/taxonomyhome.html/index.cgi?chapter=tgencodes"),
	}
}

// Edges of the Genome.
func (Genome) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("genes", Gene.Type),
		edge.To("scaffolds", Scaffold.Type),
	}
}
