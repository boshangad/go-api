package controller

import (
	"github.com/boshangad/go-api/core/global"
	"github.com/boshangad/go-api/ent"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

// ControllerInterface 接口
type ControllerInterface interface {
	// Init 前置执行初始化
	Init(ctx *gin.Context) *ControllerInterface
	// GetParamWithString 通过键 获取字符串入参
	GetParamWithString(string) string
	// JsonOut 输出json字符串
	JsonOut(int, string, interface{})
}

type Controller struct {
	ControllerInterface
	Context *gin.Context
	App *ent.App
	AppUser *ent.AppUser
	AppUserToken *ent.AppUserToken
}

func (gh *Controller) Init(c *gin.Context) *Controller {
	gh.Context = c
	keyData, ok := c.Get("App")
	if ok {
		gh.App, ok = keyData.(*ent.App)
	}
	keyData, ok = c.Get("AppUser")
	if ok {
		gh.AppUser, ok = keyData.(*ent.AppUser)
	}
	keyData, ok = c.Get("AppUserToken")
	if ok {
		gh.AppUserToken, ok = keyData.(*ent.AppUserToken)
	}
	return gh
}

func (gh Controller) GetParamWithString(key string) string {
	method := strings.ToUpper(gh.Context.Request.Method)
	if method == http.MethodGet {
		return gh.Context.DefaultQuery(key, "")
	}
	str := ""
	if method == http.MethodPost {
		str = gh.Context.PostForm(key)
	}
	if str == "" {
		return gh.Context.DefaultQuery(key, "")
	}
	return str
}

// JsonOut 输出json数据
func (gh Controller) JsonOut(error int64, msg string, data interface{})  {
	var response global.JsonResponse
	response.Error = error
	if msg == "" {
		msg = "Success"
	}
	response.Msg = msg
	response.Data = data
	gh.Context.AbortWithStatusJSON(http.StatusOK, response)
}

func (gh Controller) JsonOutByError(error int64, msg error, data interface{})  {
	gh.JsonOut(error, msg.Error(), data)
}