package schema

import "entgo.io/ent"

// AuthRule holds the schema definition for the AuthRule entity.
type AuthRule struct {
	ent.Schema
}

// Fields of the AuthRule.
func (AuthRule) Fields() []ent.Field {
	return nil
}

// Edges of the AuthRule.
func (AuthRule) Edges() []ent.Edge {
	return nil
}
