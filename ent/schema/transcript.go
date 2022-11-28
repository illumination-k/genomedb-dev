package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// Transcript holds the schema definition for the Transcript entity.
type Transcript struct {
	ent.Schema
}

// Fields of the Transcript.
func (Transcript) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").Unique(),
		field.String("gene"),
		field.Text("mrna"),
		field.Text("cds"),
		field.Text("protein"),
	}
}

// Edges of the Transcript.
func (Transcript) Edges() []ent.Edge {
	return nil
}

func (Transcript) Index() []ent.Index {
	return []ent.Index{index.Fields("id")}
}