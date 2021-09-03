package router

import (
	"github.com/boshangad/go-api/app/middlewares"
	"github.com/boshangad/go-api/app/routers"
	"github.com/boshangad/go-api/core/gin"
	"github.com/boshangad/go-api/core/global"
)

// 初始化加载路由
func init()  {
	server := gin.Get()
	server.Static("/static", global.GetPathWithStatic())
	rootGroup := server.Group("/")
	rootGroup.Use(
		middlewares.AddHeaderRequest,
		middlewares.LoadAppUser,
		middlewares.LoadApp,
		middlewares.CheckAuth,
		)
	routers.New(server, rootGroup)
}
