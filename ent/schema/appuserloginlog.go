package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/boshangad/v1/app/model/mixins"
)

// AppUserLoginLog holds the schema definition for the AppUserLoginLog entity.
type AppUserLoginLog struct {
	ent.Schema
}

func (AppUserLoginLog) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{
			Table: "app_user_login_log",
		},
	}
}

func (AppUserLoginLog) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.Primary{},
		mixins.CreateTime{},
	}
}

// Fields of the AppUserLoginLog.
func (AppUserLoginLog) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("app_id").Default(0).Comment("应用"),
		field.Uint64("app_user_id").Default(0).Comment("应用用户"),
		field.Uint64("user_id").Default(0).Comment("用户"),
		field.Uint("login_type_id").Default(0).Comment("登录方案"),
		field.String("ip").Default("127.0.0.1").Comment("登录IP").MaxLen(16),
		field.String("content").Comment("内容").Optional(),
		field.Uint("status").Default(0).Comment("状态"),
	}
}

// Edges of the AppUserLoginLog.
func (AppUserLoginLog) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("app", App.Type).Field("app_id").Required().Unique(),
		edge.To("appUser", AppUser.Type).Field("app_user_id").Required().Unique(),
		edge.To("user", User.Type).Field("user_id").Required().Unique(),
	}
}

func (AppUserLoginLog) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("app_user_id"),
	}
}
