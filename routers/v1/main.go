package v1

import "github.com/gin-gonic/gin"

type routerInterface interface {
	Init(e *gin.Engine, g *gin.RouterGroup)
}

type routerStruct struct {
	routerInterface
	Gin *gin.Engine
	Group *gin.RouterGroup
}

func (r *routerStruct) SetEngine(e *gin.Engine) *routerStruct {
	r.Gin = e
	return r
}

func (r *routerStruct) SetGroup(g *gin.RouterGroup) *routerStruct {
	r.Group = g
	return r
}

func New(e *gin.Engine, g *gin.RouterGroup) {
	new(mainController).Init(e, g)
}
