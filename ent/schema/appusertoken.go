package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/boshangad/v1/app/model/mixins"
	"github.com/google/uuid"
)

// AppUserToken holds the schema definition for the AppUserToken entity.
type AppUserToken struct {
	ent.Schema
}

func (AppUserToken) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{
			Table: "app_user_token",
		},
	}
}

func (AppUserToken) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.Primary{},
		mixins.CreateTime{},
	}
}

// Fields of the AppUserToken.
func (AppUserToken) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("app_id").Default(0).Comment("应用"),
		field.Uint64("app_user_id").Default(0).Comment("应用用户"),
		field.Uint64("user_id").Default(0).Comment("用户"),
		field.String("client_version").Default("").Comment("客户端版本").MaxLen(255),
		field.Bytes("uuid").GoType(&uuid.UUID{}).Comment("设备唯一标识"),
		field.String("ip").Default("127.0.0.1").Comment("IP地址").MaxLen(16),
		field.Int64("expire_time").Default(0).Comment("失效时间"),
	}
}

// Edges of the AppUserToken.
func (AppUserToken) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("app", App.Type).Field("app_id").Required().Unique(),
		edge.To("appUser", AppUser.Type).Field("app_user_id").Required().Unique(),
		edge.To("user", User.Type).Field("user_id").Required().Unique(),
	}
}

func (AppUserToken) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("app_id"),
		index.Fields("app_user_id"),
		index.Fields("uuid"),
	}
}
