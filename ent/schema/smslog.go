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

// SmsLog holds the schema definition for the SmsLog entity.
type SmsLog struct {
	ent.Schema
}

func (SmsLog) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{
			Table: "sms_log",
		},
	}
}

func (SmsLog) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.Primary{},
		mixins.CreateTime{},
		mixins.CreateBy{},
		mixins.UpdateTime{},
		mixins.UpdateBy{},
	}
}

// Fields of the SmsLog.
func (SmsLog) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("app_id").Default(0).Comment("应用"),
		field.String("dial_code").Default("86").Comment("拨号区号").MaxLen(4),
		field.String("mobile").Default("").Comment("拨号区号").MaxLen(32),
		field.String("scope").Default("common").Comment("范围").MaxLen(32),
		field.Uint64("type_id").Default(0).Comment("类型"),
		field.String("gateway").Default("default").Comment("网关").MaxLen(128),
		field.String("ip").Default("127.0.0.1").Comment("IP").MaxLen(16),
		field.String("template_id").Default("").Comment("短信模板").MaxLen(128),
		field.String("template_text").Default("").Comment("模板文案").MaxLen(1024).
			Annotations(
				schema.Annotation(entsql.Annotation{Collation: "utf8mb4_general_ci"}),
			),
		field.String("content").Default("").Comment("内容").MaxLen(1024).
			Annotations(
				schema.Annotation(entsql.Annotation{Collation: "utf8mb4_general_ci"}),
			),
		field.Uint8("check_count").Default(0).Comment("检查次数"),
		field.Uint("status").Default(0).Comment("状态"),
		field.String("return_msg").Default("").Comment("返回内容").MaxLen(1024),
	}
}

// Edges of the SmsLog.
func (SmsLog) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("app", App.Type).Field("app_id").Required().Unique().Comment("关联应用"),
	}
}

func (SmsLog) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("app_id"),
	}
}
