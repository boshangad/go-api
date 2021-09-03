package main

type templateStruct struct {
	// 包名称，默认为routers
	PackageName string
	// 控制器相对路径
	ControllerPath string
	// 控制器的名称
	QuoteName string
	// 相关的函数数据
	Functions map[string]routeRegisterFun
}

func getTemplateStr() string {
	return `{{$ = .}}package {{$.PackageName}}

import (
	"github.com/gin-gonic/gin"
	"github.com/tend/wechatServer/{{$.ControllerPath}}"
)

type {{$.QuoteName}}Controller struct {
	routerStruct
}

func (c *{{$.QuoteName}}Controller) Init(e *gin.Engine, g *gin.RouterGroup) *{{$.QuoteName}}Controller {
	c.SetEngine(e).SetGroup(g)
	{{range $.Functions}}
		c.{{.Name}}()
	{{end}}
	return c
}

{{range $.Functions}}
func (c {{$.QuoteName}}Controller) {{.Name}}() {{$.QuoteName}}Controller {
	g := controllers.{{.Name}}{}
	groups := e.Group.Group("/{{.Path}}")
	{
		{{range $item := .Items}}
			{{range $item.Methods}}
				groups.{{.}}("/{{$item.Path}}", func(c *gin.Context)) {
					g.Init(c)
					g.{{$item.Name}}()
				}
			{{end}}
		{{end}}
	}
	return c
}
{{end}}
`
}
