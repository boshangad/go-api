package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/boshangad/v1/app/model/mixins"
	"github.com/google/uuid"
)

// Sort holds the schema definition for the Sort entity.
type Sort struct {
	ent.Schema
}

// 混入
func (Sort) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.Primary{},
		mixins.DeleteTime{},
		mixins.CreateTime{},
		mixins.CreateBy{},
		mixins.UpdateTime{},
		mixins.UpdateBy{},
	}
}

// Fields of the Sort.
func (Sort) Fields() []ent.Field {
	return []ent.Field{
		field.Bytes("uuid").GoType(&uuid.UUID{}).DefaultFunc(func() *uuid.UUID {
			uid := uuid.New()
			return &uid
		}).MaxLen(16).Comment("UUID").Immutable(),
		field.Uint64("parent_id").Default(0).Comment("父类"),
		field.String("title").MaxLen(128).Comment("标题"),
		field.String("parent_list").MaxLen(255).Default("").Comment("族谱"),
		field.Uint64("icon_id").Default(0).Comment("图标"),
		field.String("icon_url").MaxLen(128).Comment("图标路径"),
		field.Uint16("display_order").Default(0).Max(65535).Comment("排序"),
		field.Uint("level").Default(0).Max(255).Comment("等级"),
		field.Uint("status").Default(0).Max(255).Comment("状态"),
	}
}

// Edges of the Sort.
func (Sort) Edges() []ent.Edge {
	return nil
}
