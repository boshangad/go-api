package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
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
	// model := appUserTokenService.NewModel()
	// err := model.LoginByAccessToken(authorization)
	// if err != nil {
	// 	c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
	// 		"error": http.StatusForbidden,
	// 		"msg":   err.Error(),
	// 	})
	// 	return
	// }
	// helpers.SetGinAppUserToken(c, model.AppUserToken)
	// helpers.SetGinApp(c, model.AppUserToken.Edges.App)
	// helpers.SetGinAppUser(c, model.AppUserToken.Edges.AppUser)
}
