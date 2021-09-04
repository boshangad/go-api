package mixins

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

type UpdateTime struct{
	mixin.Schema
}

func (UpdateTime) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("update_time").Default(0).Comment("更新时间"),
	}
}
