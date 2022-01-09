package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	cannotations "github.com/boshangad/v1/app/model/annotations"
	"github.com/boshangad/v1/app/model/mixins"
	"github.com/google/uuid"
)

// ResourceFile holds the schema definition for the ResourceFile entity.
type ResourceFile struct {
	ent.Schema
}

// 混入
func (ResourceFile) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.DeleteTime{},
		mixins.CreateTime{},
		mixins.CreateBy{},
	}
}

// Fields of the ResourceFile.
func (ResourceFile) Fields() []ent.Field {
	return []ent.Field{
		field.Bytes("uuid").GoType(&uuid.UUID{}).MaxLen(16).DefaultFunc(func() *uuid.UUID {
			u, _ := uuid.NewUUID()
			return &u
		}).Comment("UUID"),
		field.String("title").MaxLen(255).Default("").Comment("文件名"),
		field.String("ext").MaxLen(64).Default("").Comment("后缀名"),
		field.String("mime").MaxLen(128).Default("").Comment("MIME类型"),
		field.Uint64("size").Min(0).Default(0).Comment("文件大小Bit"),
		field.String("path").MaxLen(128).Default("").Comment("路径"),
		field.String("url").MaxLen(128).Default("").Comment("访问路径"),
		field.String("md5").MaxLen(32).Default("").Comment("MD5"),
		field.String("sha1").MaxLen(40).Default("").Comment("SHA1"),
		field.Uint("status").Min(0).Max(255).Default(0).Comment("状态").Annotations(
			cannotations.ConstAnnotation{
				Values: []cannotations.ConstData{
					{
						Name:       "StatusWaitUpload",
						Value:      0,
						Annotation: "待上传",
					},
					{
						Name:       "StatusWaitResolution",
						Value:      1,
						Annotation: "待解析",
					},
					{
						Name:       "StatusWaitRelease",
						Value:      2,
						Annotation: "待发布",
					},
					{
						Name:       "StatusPublished",
						Value:      99,
						Annotation: "已发布",
					},
				},
			},
		),
	}
}

// Edges of the ResourceFile.
func (ResourceFile) Edges() []ent.Edge {
	return nil
}
