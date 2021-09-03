package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/tend/wechatServer/app/services"
	"github.com/tend/wechatServer/core/global"
	"github.com/tend/wechatServer/ent"
	"net/http"
)

// LoadApp 加载应用模型
func LoadApp(c *gin.Context)  {
	appAlias := ""
	if c.Request.Method == http.MethodPost {
		appAlias = c.DefaultPostForm("app_alias", "")
	}
	if appAlias == "" {
		appAlias = c.DefaultQuery("app_alias", "")
	}
	if appAlias == "" {
		app, ok := c.Get("App")
		if !ok {
			c.AbortWithStatusJSON(http.StatusBadRequest, global.JsonResponse{
				Error: http.StatusBadRequest,
				Msg: "Access failed, please check whether the application authentication parameters exist",
			})
		} else {
			_, ok = app.(*ent.App)
			if !ok {
				c.AbortWithStatusJSON(http.StatusInternalServerError, global.JsonResponse{
					Error: http.StatusInternalServerError,
					Msg: "Access failed, unable to verify that the application is valid",
				})
			}
		}
		return
	}
	// 检查应用是否匹配
	app, ok := c.Get("App")
	if ok {
		appModel, ok := app.(*ent.App)
		if !ok {
			c.AbortWithStatusJSON(http.StatusInternalServerError, global.JsonResponse{
				Error: http.StatusInternalServerError,
				Msg: "Access failed, unable to verify that the application is valid",
			})
			return
		}
		// 应用别名 和 token的别名不一致
		if appModel.Alias != appAlias {
			c.AbortWithStatusJSON(http.StatusNotAcceptable, global.JsonResponse{
				Error: http.StatusNotAcceptable,
				Msg: "Access failed, application authentication match failed",
			})
		}
		return
	}
	// 如果应用不存在则需要查找应用
	appModel := services.GetAppModelByAlias(appAlias)
	if appModel == nil {
		c.AbortWithStatusJSON(http.StatusNotAcceptable, global.JsonResponse{
			Error: http.StatusNotAcceptable,
			Msg: "Access failed, application not found",
		})
		return
	}
	c.Set("App", appModel)
}