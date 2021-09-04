package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/boshangad/go-api/ent/schema/mixins"
)

// AppUserLogin holds the schema definition for the AppUserLogin entity.
type AppUserLogin struct {
	ent.Schema
}

func (AppUserLogin) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{
			Table: "app_user_login",
		},
	}
}

func (AppUserLogin) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.Primary{},
		mixins.CreateTime{},
	}
}

// Fields of the AppUserLogin.
func (AppUserLogin) Fields() []ent.Field {
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

// Edges of the AppUserLogin.
func (AppUserLogin) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("app", App.Type).Field("app_id").Required().Unique(),
		edge.To("appUser", AppUser.Type).Field("app_user_id").Required().Unique(),
		edge.To("user", User.Type).Field("user_id").Required().Unique(),
	}
}

func (AppUserLogin) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("app_user_id"),
	}
}