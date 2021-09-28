package middlewares

import (
	"fmt"
	"github.com/boshangad/go-api/global"
	"github.com/boshangad/go-api/services"
	"github.com/boshangad/go-api/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

// LoadApp 加载应用模型
func LoadApp(c *gin.Context) {
	appModel := utils.GetGinApp(c)
	if appModel == nil {
		var data map[string]string
		var appAlias string = strings.TrimSpace(c.GetHeader(global.G_CONFIG.App.AppParamName))
		if appAlias == "" && c.Request.Method != http.MethodGet {
			_ = c.ShouldBind(&data)
			if v, ok := data[global.G_CONFIG.App.AppParamName]; ok {
				appAlias = strings.TrimSpace(v)
			}
		}
		if appAlias == "" {
			appAlias = strings.TrimSpace(c.DefaultQuery(global.G_CONFIG.App.AppParamName, ""))
		}
		// 如果用户没有传入token 或 没有 传入参数
		if appAlias == "" {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error": http.StatusBadRequest,
				"msg": "Access failed, please check whether the application authentication parameters exist",
			})
			return
		}
		// 如果应用不存在则需要查找应用
		appModel = services.GetAppModelByAlias(appAlias)
		if appModel == nil {
			c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{
				"error": http.StatusNotAcceptable,
				"msg": fmt.Sprintf("Access failed, #%s application not found.", appAlias),
			})
			return
		}
		utils.SetGinApp(c, appModel)
		return
	}
}