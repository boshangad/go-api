package middlewares

import (
	"net/http"
	"strings"

	"github.com/boshangad/v1/app/global"
	"github.com/boshangad/v1/app/helpers"
	"github.com/gin-gonic/gin"
)

func AppMiddleware(c *gin.Context) {
	appModel := helpers.GetGinApp(c)
	if appModel == nil {
		var data map[string]string
		var appAlias string = strings.TrimSpace(c.GetHeader(global.Config.App.Name))
		if appAlias == "" && c.Request.Method != http.MethodGet {
			_ = c.ShouldBind(&data)
			if v, ok := data[global.Config.App.Name]; ok {
				appAlias = strings.TrimSpace(v)
			}
		}
		if appAlias == "" {
			appAlias = strings.TrimSpace(c.DefaultQuery(global.Config.App.Name, ""))
		}
		// 如果用户没有传入token 或 没有 传入参数
		if appAlias == "" {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error": http.StatusBadRequest,
				"msg":   "Access failed, please check whether the application authentication parameters exist",
			})
			return
		}
		// 如果应用不存在则需要查找应用
		// appModel = services.GetAppModelByAlias(appAlias)
		// if appModel == nil {
		// 	c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{
		// 		"error": http.StatusNotAcceptable,
		// 		"msg":   fmt.Sprintf("Access failed, #%s application not found.", appAlias),
		// 	})
		// 	return
		// }
		helpers.SetGinApp(c, appModel)
		return
	}
}
