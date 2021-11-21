package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type CorsConfig struct {
	// 是否允许跨域
	AllowCors bool `mapstructure:"allowCors" json:"allowCors,omitempty" yaml:"allowCors"`
	// 允许跨域的地址
	AllowOrigins []string `mapstructure:"allowOrigins" json:"allowOrigins,omitempty" yaml:"allowOrigins"`
	// 允许跨域的请求方法
	AllowMethods []string `mapstructure:"allowMethods" json:"allowMethods,omitempty" yaml:"allowMethods"`
	// 响应报头指示的请求的响应是否可以暴露于该页面
	AllowCredentials bool `mapstructure:"allowCredentials" json:"allowCredentials,omitempty" yaml:"allowCredentials"`
	// 暴露的请求头
	ExposeHeaders []string `mapstructure:"exposeHeaders" json:"exposeHeaders,omitempty" yaml:"exposeHeaders"`
	// 允许缓存的时长
	MaxAge int64 `mapstructure:"maxAge" json:"maxAge,omitempty" yaml:"maxAge"`
}

// Cors 处理跨域请求,支持options访问
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		origin := c.Request.Header.Get("Origin")
		c.Header("Access-Control-Allow-Origin", origin)
		// 允许访问的请求头
		c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token,X-Token,X-User-Id")
		c.Header("Access-Control-Allow-Methods", "POST,GET,OPTIONS,DELETE,PUT")
		// 允许暴露的请求头
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Max-Age", "3000")
		// 放行所有OPTIONS方法
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}
	}
}
