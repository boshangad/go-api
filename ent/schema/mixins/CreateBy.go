package mixins

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

type CreateBy struct{
	mixin.Schema
}

func (CreateBy) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("create_by").Default(0).Comment("创建人").Immutable(),
	}
}
