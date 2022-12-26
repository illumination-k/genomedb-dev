package schema

import "entgo.io/ent"

type Pfam struct {
	ent.Schema
}

func (Pfam) Fields() []ent.Field {
	return []ent.Field{}
}

func (Pfam) Edges() []ent.Edge {
	return nil
}
