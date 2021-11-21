package server

import (
	"context"
	"path/filepath"
	"time"

	"github.com/boshangad/v1/app/global"
	"github.com/boshangad/v1/app/middlewares"
	"github.com/boshangad/v1/routers"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// 服务接口
type Server interface {
	ListenAndServe() error
	Shutdown(ctx context.Context) error
}

// 路由存在，但方法不被允许报错
func pageMethodNotAllow() gin.HandlerFunc {
	return func(c *gin.Context) {
		global.Log.Info(
			"The page is denied access, the method is not allowed",
			zap.String("path", c.Request.RequestURI),
			zap.String("method", c.Request.Method),
		)
	}
}

// 路由未定义
func pageNotFound() gin.HandlerFunc {
	return func(c *gin.Context) {
		global.Log.Debug(
			"page not found.",
			zap.String("path", c.Request.RequestURI),
		)
	}
}

// 实例化一个服务
func NewServer(addr string) Server {
	e := gin.New()
	e.Use(
		// 记录日志
		Logger(),
		// 异常恢复
		Recovery(true),
	)
	// 定义未配置路由的错误
	e.NoRoute(pageNotFound())
	e.NoMethod(pageMethodNotAllow())

	// 为用户头像和文件提供静态地址
	e.StaticFS(filepath.Join("static"), gin.Dir(filepath.Join(global.Config.App.RootPath, "static"), false))
	// 打开就能玩https了
	// e.Use(middleware.LoadTls())
	// 跨域,如需跨域可以打开
	e.Use(middlewares.Cors())
	//e.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	// 获取路由组实例
	PublicGroup := e.Group("")
	{
		routers.V1.Init(PublicGroup)
	}
	s := InitServer(addr, e)
	return s
}

// 启动服务
func Run(s Server) error {
	return s.ListenAndServe()
}

// 重启服务
func Reload() error {
	return nil
}

// 停止关闭服务
func Close(s Server) error {
	global.Log.Info("shutting down gracefully, press Ctrl+C again to force")
	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	serverCtx, serverCancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer serverCancel()
	return s.Shutdown(serverCtx)
}
