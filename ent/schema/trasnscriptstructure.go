package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// TrasnscriptStructure holds the schema definition for the TrasnscriptStructure entity.
type TrasnscriptStructure struct {
	ent.Schema
}

// Fields of the TrasnscriptStructure.
func (TrasnscriptStructure) Fields() []ent.Field {
	return []ent.Field{
		field.String("transcriptId"),
		field.String("feature"),
		field.String("seqname"),
		field.Int32("start"),
		field.Int32("end"),
		field.String("strand"),
	}
}

// Edges of the TrasnscriptStructure.
func (TrasnscriptStructure) Edges() []ent.Edge {
	return nil
}
