package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"strings"
)

func AddHeaderRequest(c *gin.Context)  {
	c.Header("Request-Id", strings.ReplaceAll(uuid.New().String(), "-", ""))
	c.Next()
	// 后置执行
}
