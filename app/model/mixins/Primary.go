package mixins

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

type Primary struct {
	mixin.Schema
}

func (Primary) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id").Comment("ID").Unique().Immutable(),
	}
}
