package gin

import (
	"github.com/gin-gonic/gin"
	"github.com/tend/wechatServer/core/config"
	"github.com/tend/wechatServer/core/global"
	"github.com/tend/wechatServer/ent"
	"net/http"
)

var engine *gin.Engine

// New 实例化一个Gin引擎
func New() *gin.Engine {
	gin.SetMode(config.Get().Mode)
	engine = gin.New()
	engine.NoRoute(func(c *gin.Context) {
		c.Abort()
		c.SecureJSON(http.StatusNotFound, global.JsonResponse{
			Error: http.StatusNotFound,
			Msg: "Page not found.",
		})
	})
	engine.Use(gin.CustomRecovery(func(c *gin.Context, err interface{}) {
		c.AbortWithStatusJSON(http.StatusInternalServerError, global.JsonResponse{
			Error: http.StatusInternalServerError,
			Msg: "Service exception, please try again",
		})
	}))
	engine.Group("/")
	engine.GET("/entviz", gin.WrapH(ent.ServeEntviz()))
	return engine
}

// Get 获取Gin引擎
func Get() *gin.Engine {
	if engine == nil {
		return New()
	}
	return engine
}
