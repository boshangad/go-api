package middlewares

import (
	"fmt"
	"github.com/boshangad/go-api/app/services"
	"github.com/boshangad/go-api/core/global"
	"github.com/boshangad/go-api/ent"
	"github.com/gin-gonic/gin"
	"net/http"
)

// LoadApp 加载应用模型
func LoadApp(c *gin.Context) {
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
			Msg: fmt.Sprintf("Access failed, #%s application not found.", appAlias),
		})
		return
	}
	c.Set("App", appModel)
}