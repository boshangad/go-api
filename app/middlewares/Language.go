package middlewares

import "github.com/gin-gonic/gin"

func lang(c *gin.Context) *gin.Context {
	lang := c.DefaultQuery("lang", "")
	if lang == "" {
		lang = c.GetHeader("accept-language")
	}

	return c
}