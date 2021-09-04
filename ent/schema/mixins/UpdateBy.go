package mixins

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

type UpdateBy struct{
	mixin.Schema
}

func (UpdateBy) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("update_by").Default(0).Comment("更新人"),
	}
}
