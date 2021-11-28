package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	cannotations "github.com/boshangad/v1/app/model/annotations"
	"github.com/boshangad/v1/app/model/mixins"
)

// AppOption holds the schema definition for the AppOption entity.
type AppOption struct {
	ent.Schema
}

func (AppOption) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{
			Table: "app_option",
		},
	}
}

func (AppOption) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.Primary{},
		mixins.CreateTime{},
		mixins.CreateBy{},
		mixins.UpdateTime{},
		mixins.UpdateBy{},
	}
}

// Fields of the AppOption.
func (AppOption) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("app_id").Default(0).Comment("应用"),
		field.String("title").Default("").Comment("标题").MaxLen(64),
		field.String("description").Default("").Comment("描述").MaxLen(225),
		field.String("name").Default("").Comment("键").MaxLen(128),
		field.String("value").Default("").Comment("值").MaxLen(255),
		field.Int64("expire_time").Default(0).Comment("失效时间"),
		field.Uint("edit_type").Default(0).Comment("编辑类型").Annotations(
			cannotations.ConstAnnotation{
				Values: []cannotations.ConstData{
					{
						Name:       "EditTypeOnlyRead",
						Value:      0,
						Annotation: "只读",
					},
					{
						Name:       "EditTypeNumberInput",
						Value:      1,
						Annotation: "数值输入框",
					},
					{
						Name:       "EditTypeTextInput",
						Value:      2,
						Annotation: "文本输入框",
					},
					{
						Name:       "EditTypeTextarea",
						Value:      3,
						Annotation: "多行文本框",
					},
				},
			},
		),
	}
}

// Edges of the AppOption.
func (AppOption) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("app", App.Type).Ref("appOptions").Field("app_id").Required().Unique(),
	}
}

func (AppOption) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("app_id"),
		index.Fields("name"),
	}
}
