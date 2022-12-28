package schema

import (
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Kegg holds the schema definition for the Kegg entity.
type KeggOrthlogy struct {
	ent.Schema
}

// Fields of the Kegg.
func (KeggOrthlogy) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").Validate(func(s string) error {
			if !strings.HasPrefix(s, "ko") {
				return fmt.Errorf("Kegg Ontology ID should has prefix ko: %s", s)
			}
			return nil
		}),
		field.String("name"),
	}
}

// Edges of the Kegg.
func (KeggOrthlogy) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("pathways", KeggPathway.Type).Ref("orthologies"),
	}
}

type KeggPathway struct {
	ent.Schema
}

// Fields of the Kegg.
func (KeggPathway) Fields() []ent.Field {
	return []ent.Field{
		field.String("id"),
		field.String("name"),
	}
}

// Edges of the Kegg.
func (KeggPathway) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("related_map", KeggPathway.Type).
			From("relating_map"),
		edge.To("reactions", KeggReaction.Type),
		edge.To("orthologies", KeggOrthlogy.Type),
	}
}

type KeggModule struct {
	ent.Schema
}

// Fields of the Kegg.
func (KeggModule) Fields() []ent.Field {
	return []ent.Field{
		field.String("id"),
	}
}

// Edges of the Kegg.
func (KeggModule) Edges() []ent.Edge {
	return nil
}

type KeggReaction struct {
	ent.Schema
}

// Fields of the Kegg.
func (KeggReaction) Fields() []ent.Field {
	return []ent.Field{
		field.String("id"),
		field.String("name"),
	}
}

// Edges of the Kegg.
func (KeggReaction) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("pathways", KeggPathway.Type).Ref("reactions"),
	}
}

type KeggCompound struct {
	ent.Schema
}

// Fields of the Kegg.
func (KeggCompound) Fields() []ent.Field {
	return []ent.Field{
		field.String("id"),
	}
}

// Edges of the Kegg.
func (KeggCompound) Edges() []ent.Edge {
	return nil
}
