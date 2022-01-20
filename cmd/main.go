package main

import (
	"path/filepath"

	"github.com/boshangad/v1/app"
	"github.com/boshangad/v1/global"
	"github.com/boshangad/v1/routers"
	"github.com/gin-gonic/gin"
)

func main() {
	// 初始化配置
	// 执行应用
	app.Run(global.Config, func(e *gin.Engine) {
		// 为用户头像和文件提供静态地址
		e.StaticFS(filepath.Join("static"), gin.Dir(filepath.Join(global.Config.App.RootPath, "static"), false))
		// 注册根组件
		rootGroup := e.Group("")
		{
			// // 跨域,如需跨域可以打开
			// rootGroup.Use(
			// 	middlewares.Cors(),
			// 	// controller.BindingFunc(middlewares.AppUserMiddleware),
			// 	// controller.BindingFunc(middlewares.AppMiddleware),
			// )
			// 获取路由组实例
			publicGroup := rootGroup.Group("")
			{
				routers.V1.Init(publicGroup)
			}
		}
	})
}
