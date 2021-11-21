package schema

import "entgo.io/ent"

// AuthItemChild holds the schema definition for the AuthItemChild entity.
type AuthItemChild struct {
	ent.Schema
}

// Fields of the AuthItemChild.
func (AuthItemChild) Fields() []ent.Field {
	return nil
}

// Edges of the AuthItemChild.
func (AuthItemChild) Edges() []ent.Edge {
	return nil
}
