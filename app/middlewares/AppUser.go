package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/boshangad/go-api/app/services"
	"github.com/boshangad/go-api/core/global"
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
	authData, err := services.CheckTokenValid(authorization)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusForbidden, global.JsonResponse{
			Error: http.StatusForbidden,
			Msg: err.Error(),
		})
		return
	}
	c.Set("App", authData.App)
	c.Set("AppUser", authData.AppUser)
	c.Set("AppUserToken", authData.AppUserToken)
}