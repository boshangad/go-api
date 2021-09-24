package middlewares

import (
	"github.com/boshangad/go-api/services/appUserTokenService"
	"github.com/gin-gonic/gin"
	"net/http"
)

// LoadAppUser 检查应用用户是否有效
func LoadAppUser(c *gin.Context) {
	authorization := c.GetHeader("authorization")
	if authorization == "" {
		if c.Request.Method == http.MethodPost {
			authorization = c.DefaultPostForm("access_token", "")
		}
		if authorization == "" {
			authorization = c.DefaultQuery("access_token", "")
		}
	}
	if authorization == "" {
		return
	}
	authData, err := appUserTokenService.CheckTokenValid(authorization)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
			"error": http.StatusForbidden,
			"msg": err.Error(),
		})
		return
	}
	c.Set("App", authData.App)
	c.Set("AppUser", authData.AppUser)
	c.Set("AppUserToken", authData.AppUserToken)
}