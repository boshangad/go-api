package Annotations

import "entgo.io/ent/schema"

// 默认注释
var DefaultStatusAnnotation schema.Annotation = ConstAnnotation{
	Values: []ConstData{
		{
			Name:       "StatusDraft",
			Value:      0,
			Annotation: "草稿",
		},
		{
			Name:       "StatusRelease",
			Value:      1,
			Annotation: "发布",
		},
	},
}

type ConstData struct {
	// 名称
	Name string
	// 值
	Value interface{}
	// 注释
	Annotation string
}

// 常量注解
type ConstAnnotation struct {
	Values []ConstData
}

// 实现注解方法
func (that ConstAnnotation) Name() string {
	return "ConstAnnotation"
}
