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

// AppUser holds the schema definition for the AppUser entity.
type AppUser struct {
	ent.Schema
}

func (AppUser) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{
			Table: "app_user",
		},
	}
}

func (AppUser) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.Primary{},
		mixins.CreateTime{},
		mixins.UpdateTime{},
	}
}

// Fields of the AppUser.
func (AppUser) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("app_id").Default(0).Comment("应用"),
		field.Uint64("user_id").Default(0).Comment("用户"),
		field.String("open_id").Default("").Comment("微信开放ID").MaxLen(64),
		field.String("unionid").Default("").Comment("微信联合用户").MaxLen(32),
		field.String("session_key").Default("").Comment("会话密钥").MaxLen(64),
		field.Bool("is_load_user_profile").Default(false).Comment("是否加载用户信息"),
		field.String("nickname").Default("").Comment("昵称").MaxLen(64).
			Annotations(
				schema.Annotation(entsql.Annotation{Collation: "utf8mb4_general_ci"}),
			),
		field.String("avatar").Default("").Comment("头像").MaxLen(255),
		field.String("avatar_url").Default("").Comment("头像").MaxLen(255),
		field.Uint("gender").Default(0).Comment("性别").Range(0, 2),
		field.String("county").Default("").Comment("国籍").MaxLen(64),
		field.String("country_code").Default("").Comment("国籍代码").MaxLen(20),
		field.String("province").Default("").Comment("省份").MaxLen(64),
		field.String("city").Default("").Comment("城市").MaxLen(64),
		field.String("language").Default("zh_CN").Comment("语言").MaxLen(20),
		field.String("phone_number").Default("").Comment("联系号码").MaxLen(32),
		field.String("pure_phone_number").Default("").Comment("联系号码").MaxLen(32),
		field.String("watermark").Default("").Comment("备注").MaxLen(255),
		field.Uint64("load_user_profile_time").Default(0).Comment("加载用户信息时间"),
		field.Uint64("last_login_time").Default(0).Comment("最后登录时间"),
	}
}

// Edges of the AppUser.
func (AppUser) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("app", App.Type).Field("app_id").Required().Unique(),
		edge.To("user", User.Type).Field("user_id").Required().Unique(),
	}
}

func (AppUser) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("app_id"),
		index.Fields("user_id"),
		index.Fields("open_id"),
		index.Fields("unionid"),
	}
}
