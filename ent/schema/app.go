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
	"github.com/google/uuid"
)

// App holds the schema definition for the App entity.
type App struct {
	ent.Schema
}

// 注释
func (App) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{
			Table: "app",
		},
	}
}

// 混入
func (App) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.Primary{},
		mixins.DeleteTime{},
		mixins.CreateTime{},
		mixins.CreateBy{},
		mixins.UpdateTime{},
		mixins.UpdateBy{},
	}
}

// Fields of the App.
func (App) Fields() []ent.Field {
	return []ent.Field{
		field.Bytes("uuid").GoType(&uuid.UUID{}).MaxLen(16).DefaultFunc(func() *uuid.UUID {
			u, _ := uuid.NewUUID()
			return &u
		}).Comment("UUID"),
		field.Uint64("type_id").Default(0).Comment("类型").Annotations(
			cannotations.ConstAnnotation{
				Values: []cannotations.ConstData{
					{
						Name:       "TypePC",
						Value:      1,
						Annotation: "网页端",
					},
					{
						Name:       "TypeMobile",
						Value:      2,
						Annotation: "移动网页端",
					},
					{
						Name:       "TypeAndroid",
						Value:      3,
						Annotation: "安卓应用",
					},
					{
						Name:       "TypeIOS",
						Value:      4,
						Annotation: "苹果应用",
					},
					{
						Name:       "TypeMiniWechat",
						Value:      5,
						Annotation: "微信小程序",
					},
				},
			},
		),
		field.String("title").Default("").Comment("标题").MaxLen(128),
		field.String("intro").Default("").Comment("简介").MaxLen(255),
		field.String("mp_origin_id").Default("").Comment("原始ID").MaxLen(32),
		field.String("app_id").Default("").Comment("应用ID").MaxLen(64),
		field.String("app_secret").Default("").Comment("应用密钥").MaxLen(128),
		field.Bool("has_payment_auth").Default(false).Comment("是否有支付权限"),
		field.Uint64("register_user_number").Default(0).Comment("注册用户数量"),
		field.Uint("status").Default(0).Comment("状态").Min(0).Annotations(
			cannotations.DefaultStatusAnnotation,
		),
	}
}

// Edges of the App.
func (App) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("appOptions", AppOption.Type).Comment("应用配置项"),
	}
}

func (App) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("app_id"),
	}
}
