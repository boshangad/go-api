package router

import (
	"github.com/tend/wechatServer/app/middlewares"
	"github.com/tend/wechatServer/app/routers"
	"github.com/tend/wechatServer/core/gin"
	"github.com/tend/wechatServer/core/global"
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
