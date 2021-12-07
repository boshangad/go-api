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
		field.Uint("login_type_id").Default(0).Comment("登录方案").Annotations(
			cannotations.ConstAnnotation{
				Values: []cannotations.ConstData{
					{
						Name:       "LoginTypeUnknow",
						Value:      0,
						Annotation: "未知的登录方式",
					}, {
						Name:       "LoginTypeUsername",
						Value:      1,
						Annotation: "使用用户名登录",
					}, {
						Name:       "LoginTypeMobile",
						Value:      2,
						Annotation: "使用手机号登录",
					}, {
						Name:       "LoginTypeEmail",
						Value:      3,
						Annotation: "使用邮箱登录",
					}, {
						Name:       "LoginTypeThired",
						Value:      4,
						Annotation: "第三方快捷登录",
					},
				},
			},
		),
		field.String("ip").Default("127.0.0.1").Comment("登录IP").MaxLen(16),
		field.String("user_agent").Default("").Comment("用户代理").MaxLen(255),
		field.String("client_version").Default("").Comment("客户端版本").MaxLen(255),
		field.String("content").Comment("内容").Optional(),
		field.Uint("status").Default(0).Comment("状态").Annotations(
			cannotations.ConstAnnotation{
				Values: []cannotations.ConstData{
					{
						Name:       "StatusWaitConfirm",
						Value:      0,
						Annotation: "待确认",
					}, {
						Name:       "StatusSuccess",
						Value:      1,
						Annotation: "成功",
					}, {
						Name:       "StatusFailed",
						Value:      2,
						Annotation: "失败",
					}, {
						Name:       "StatusServerFailed",
						Value:      3,
						Annotation: "服务异常失败",
					},
				},
			},
		),
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
