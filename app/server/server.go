package server

import (
	"context"
	"time"

	"github.com/boshangad/v1/app/log"
	"github.com/gin-contrib/pprof"
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
		log.Logger.Info(
			"The page is denied access, the method is not allowed",
			zap.String("path", c.Request.RequestURI),
			zap.String("method", c.Request.Method),
		)
	}
}

// 路由未定义
func pageNotFound() gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Logger.Debug(
			"page not found.",
			zap.String("path", c.Request.RequestURI),
		)
	}
}

// 实例化一个服务
func NewServer(listen string, initEnginefun func(*gin.Engine)) Server {
	var (
		e      = gin.New()
		server Server
	)
	// 注册pprof工具
	pprof.Register(e)
	// 定义未配置路由的错误
	e.NoRoute(pageNotFound())
	e.NoMethod(pageMethodNotAllow())
	e.Use(Logger(), Recovery(true))
	if initEnginefun != nil {
		initEnginefun(e)
	}
	server = InitServer(listen, e)
	return server
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
	log.Logger.Info("shutting down gracefully, press Ctrl+C again to force")
	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	serverCtx, serverCancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer serverCancel()
	return s.Shutdown(serverCtx)
}
