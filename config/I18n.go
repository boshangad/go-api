package config

type I18n struct {
	// 默认语言
	Lang string
	// 入参方式，有 Query|Body|Header|Cookie 多个以逗号分隔，默认优先级 header,cookie,body,query
	ParamPosition string
	// 参数名称
	ParamName string
	// 语言包路径
	LangDir string
}
