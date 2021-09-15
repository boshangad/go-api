package gin

import (
	"fmt"
	"github.com/boshangad/go-api/core/config"
	"github.com/boshangad/go-api/core/global"
	"github.com/boshangad/go-api/ent"
	"github.com/gin-gonic/gin"
	"net/http"
)

var engine *gin.Engine

// New 实例化一个Gin引擎
func New() *gin.Engine {
	gin.SetMode(config.Get().Mode)
	engine = gin.New()
	// 初始化验证器
	engine.NoRoute(func(c *gin.Context) {
		c.Abort()
		c.SecureJSON(http.StatusNotFound, global.JsonResponse{
			Error: http.StatusNotFound,
			Msg: "Page not found.",
		})
	})
	// 自定义
	engine.Use(gin.CustomRecovery(func(c *gin.Context, err interface{}) {
		if gin.Mode() == gin.DebugMode {
			c.AbortWithStatusJSON(http.StatusInternalServerError, global.JsonResponse{
				Error: http.StatusInternalServerError,
				Msg: fmt.Sprintf("%s", err),
			})
			return
		}
		c.AbortWithStatusJSON(http.StatusInternalServerError, global.JsonResponse{
			Error: http.StatusInternalServerError,
			Msg: "Service exception, please try again",
		})
	}))
	// 默认的组
	engine.Group("/")
	// 测试环境下定义模型
	if gin.Mode() == gin.DebugMode {
		engine.GET("/entviz", gin.WrapH(ent.ServeEntviz()))
	}
	return engine
}

// Get 获取Gin引擎
func Get() *gin.Engine {
	if engine == nil {
		return New()
	}
	return engine
}
