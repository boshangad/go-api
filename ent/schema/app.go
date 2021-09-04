package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/boshangad/go-api/ent/schema/mixins"
	"github.com/google/uuid"
)

// App holds the schema definition for the App entity.
type App struct {
	ent.Schema
}

func (App) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{
			Table: "app",
		},
	}
}

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
		field.String("alias").DefaultFunc(uuid.New().String).Comment("别名").
			SchemaType(map[string]string{dialect.MySQL: "char(36)"}).Unique(),
		field.Uint64("type_id").Default(0).Comment("类型"),
		field.String("title").Default("").Comment("标题").MaxLen(128),
		field.String("intro").Default("").Comment("简介").MaxLen(255),
		field.String("mp_origin_id").Default("").Comment("原始ID").MaxLen(32),
		field.String("app_id").Default("").Comment("应用ID").MaxLen(64),
		field.String("app_secret").Default("").Comment("应用密钥").MaxLen(128),
		field.Bool("has_payment_auth").Default(false).Comment("是否有支付权限"),
		field.Uint64("register_user_number").Default(0).Comment("注册用户数量"),
		field.Uint("status").Default(0).Comment("状态").Min(0),
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