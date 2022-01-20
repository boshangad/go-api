package server

import (
	"time"

	"github.com/boshangad/v1/app/log"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// Logger 接收gin框架默认的日志
func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()
		cost := time.Since(start)
		if log.Logger != nil {
			path := c.Request.URL.Path
			query := c.Request.URL.RawQuery
			log.Logger.Info(path,
				zap.Int("status", c.Writer.Status()),
				zap.String("method", c.Request.Method),
				zap.String("path", path),
				zap.String("query", query),
				zap.String("ip", c.ClientIP()),
				zap.String("user-agent", c.Request.UserAgent()),
				zap.String("errors", c.Errors.ByType(gin.ErrorTypePrivate).String()),
				zap.Duration("cost", cost),
			)
		}
	}
}
