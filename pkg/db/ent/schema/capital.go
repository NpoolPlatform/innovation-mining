package schema

import "entgo.io/ent"

// Capital holds the schema definition for the Capital entity.
type Capital struct {
	ent.Schema
}

// Fields of the Capital.
func (Capital) Fields() []ent.Field {
	return nil
}

// Edges of the Capital.
func (Capital) Edges() []ent.Edge {
	return nil
}
