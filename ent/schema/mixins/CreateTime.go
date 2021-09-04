package mixins

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

type CreateTime struct{
	mixin.Schema
}

func (CreateTime) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("create_time").Default(0).Comment("创建时间").Immutable(),
	}
}
