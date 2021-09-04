package mixins

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

type DeleteTime struct{
	mixin.Schema
}

func (DeleteTime) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("delete_time").Default(0).Comment("删除时间"),
	}
}
