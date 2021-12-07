package mixins

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

type UpdateTime struct {
	mixin.Schema
}

func (UpdateTime) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("update_time").UpdateDefault(func() int64 {
			return time.Now().Unix()
		}).DefaultFunc(func() int64 {
			return time.Now().Unix()
		}).Comment("更新时间"),
	}
}
