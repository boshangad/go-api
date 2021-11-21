package schema

import "entgo.io/ent"

// AuthAssgiment holds the schema definition for the AuthAssgiment entity.
type AuthAssgiment struct {
	ent.Schema
}

// Fields of the AuthAssgiment.
func (AuthAssgiment) Fields() []ent.Field {
	return nil
}

// Edges of the AuthAssgiment.
func (AuthAssgiment) Edges() []ent.Edge {
	return nil
}
