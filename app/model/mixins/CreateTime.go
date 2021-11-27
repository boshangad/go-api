package mixins

import (
	"time"

	"entgo.io/ent"

	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

type CreateTime struct {
	mixin.Schema
}

// 字段
func (CreateTime) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("create_time").Default(time.Now().Unix()).Comment("创建时间").Immutable(),
	}
}

// // 时间钩子
// func (CreateTime) Hook() []ent.Hook {
// 	return []ent.Hook{
// 		hook.On(func(next ent.Mutator) ent.Mutator {
// 			return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
// 				// if s, ok := m.(interface{ SetCreateTime(time.Time) }); ok {
// 				// 	o, e := m.OldField(nil, "create_time")
// 				// 	if e == nil && o == nil {
// 				// 		s.SetCreateTime(time.Now())
// 				// 	}
// 				// }
// 				v, err := next.Mutate(ctx, m)
// 				return v, err
// 			})
// 		},
// 			ent.OpCreate,
// 		),
// 	}
// }
