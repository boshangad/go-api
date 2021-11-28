package middlewares

import (
	"net/http"

	"github.com/boshangad/v1/app/controller"
	"github.com/boshangad/v1/app/global"
	"github.com/boshangad/v1/ent"
	"github.com/boshangad/v1/services/appUserTokenService"
	"go.uber.org/zap"
)

// AppUserMiddleware 检查应用用户是否有效
func AppUserMiddleware(c *controller.Context) {
	var (
		authorization = c.GetHeader("authorization")
		accessToken   appUserTokenService.AccessToken
		appUserToken  *ent.AppUserToken
		err           error
	)
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
	accessToken.AccessToken = authorization
	appUserToken, err = accessToken.Login()
	if err != nil {
		global.Log.Info("token invalid", zap.String("token", authorization))
		return
	}
	c.SetAppUserToken(appUserToken)
	c.SetApp(appUserToken.Edges.App)
	c.SetAppUser(appUserToken.Edges.AppUser)
	c.SetUser(appUserToken.Edges.User)
	return
}
