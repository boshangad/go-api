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

// EmailLog holds the schema definition for the EmailLog entity.
type EmailLog struct {
	ent.Schema
}

func (EmailLog) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{
			Table: "email_log",
		},
	}
}

func (EmailLog) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.Primary{},
		mixins.CreateTime{},
		mixins.CreateBy{},
		mixins.UpdateTime{},
		mixins.UpdateBy{},
	}
}

// Fields of the EmailLog.
func (EmailLog) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("app_id").Default(0).Comment("应用"),
		field.String("email").Comment("邮箱").MaxLen(128),
		field.String("scope").Default("common").Comment("范围").MaxLen(32),
		field.Uint64("type_id").Default(0).Comment("类型"),
		field.String("gateway").Default("default").Comment("网关").MaxLen(64),
		field.String("ip").Default("127.0.0.1").Comment("IP").MaxLen(16),
		field.String("from_name").Default("").Comment("发信名称").MaxLen(128),
		field.String("from_address").Default("").Comment("发信地址").MaxLen(128),
		field.String("title").Default("").Comment("标题").MaxLen(255).
			Annotations(
				schema.Annotation(entsql.Annotation{Collation: "utf8mb4_general_ci"}),
			),
		field.String("content").Default("").Comment("短信内容").MaxLen(1024).
			Annotations(
				schema.Annotation(entsql.Annotation{Collation: "utf8mb4_general_ci"}),
			),
		field.String("data").Default("").Comment("内容参数").MaxLen(1024).
			Annotations(
				schema.Annotation(entsql.Annotation{Collation: "utf8mb4_general_ci"}),
			),
		field.Uint8("check_count").Default(0).Comment("检查次数"),
		field.Uint("status").Default(0).Comment("状态"),
		field.String("return_msg").Default("").Comment("返回消息").MaxLen(1024),
	}
}

// Edges of the EmailLog.
func (EmailLog) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("app", App.Type).Field("app_id").Required().Unique(),
	}
}

func (EmailLog) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("app_id"),
	}
}