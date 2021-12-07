package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	cannotations "github.com/boshangad/v1/app/model/annotations"
	"github.com/boshangad/v1/app/model/mixins"
	"github.com/google/uuid"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

func (User) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{
			Table: "user",
		},
	}
}

func (User) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.Primary{},
		mixins.DeleteTime{},
		mixins.CreateTime{},
		mixins.CreateBy{},
		mixins.UpdateTime{},
		mixins.UpdateBy{},
	}
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("uuid", uuid.New()).Default(uuid.New).Comment("UUID").Immutable().Unique(),
		field.String("username").Default("").Comment("用户名").MaxLen(64),
		field.String("nickname").Default("").Comment("昵称").MaxLen(64),
		field.String("password").Default("").Comment("登录密码").MaxLen(255),
		field.String("dial_code").Default("86").Comment("拨号区号").MaxLen(4),
		field.String("mobile").Default("").Comment("手机号").MaxLen(32),
		field.String("mobile_hash").Default("").Comment("手机号哈希").MaxLen(32),
		field.String("email").Default("").Comment("邮箱").MaxLen(128),
		field.String("email_hash").Default("").Comment("邮箱hash").MaxLen(32),
		field.String("avatar").Default("").Comment("头像").MaxLen(255),
		field.String("name").Default("").Comment("下面").MaxLen(64),
		field.Uint("sex").Default(0).Comment("性别").Range(0, 2),
		field.Uint64("birthday").Default(0).Comment("生日"),
		field.Uint("age").Default(0).Comment("年龄"),
		field.String("last_login_ip").Default("127.0.0.1").Comment("最后登录IP").MaxLen(16),
		field.Uint64("last_login_time").Default(0).Comment("最后登录时间"),
		field.Uint("status").Default(0).Comment("状态").Annotations(
			cannotations.ConstAnnotation{
				Values: []cannotations.ConstData{
					{
						Name:       "StatusDisabled",
						Value:      0,
						Annotation: "禁用",
					},
					{
						Name:       "StatusWaitActive",
						Value:      1,
						Annotation: "待激活",
					},
					{
						Name:       "StatusActived",
						Value:      10,
						Annotation: "激活",
					},
				},
			},
		),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return nil
}
