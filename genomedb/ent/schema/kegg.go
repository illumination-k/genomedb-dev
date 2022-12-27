package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Kegg holds the schema definition for the Kegg entity.
type KeggOntology struct {
	ent.Schema
}

// Fields of the Kegg.
func (KeggOntology) Fields() []ent.Field {
	return []ent.Field{
		field.String("id"),
		field.String("name"),
	}
}

// Edges of the Kegg.
func (KeggOntology) Edges() []ent.Edge {
	return nil
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
	return nil
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
	}
}

// Edges of the Kegg.
func (KeggReaction) Edges() []ent.Edge {
	return nil
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
