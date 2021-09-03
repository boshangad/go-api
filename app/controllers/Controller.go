package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/tend/wechatServer/core/global"
	"github.com/tend/wechatServer/ent"
	"net/http"
	"strings"
)

type ControllerInterface interface {
	// Init 前置执行初始化
	Init(ctx *gin.Context) *ControllerInterface
	// 通过键 获取字符串入参
	getParamWithString(string) string
	// 输出json字符串
	jsonOut(int, string, interface{})
}

type Controller struct {
	ControllerInterface
	Context *gin.Context
	App *ent.App
	AppUser *ent.AppUser
	AppUserToken *ent.AppUserToken
}

// Init 初始化
//@ROUTE -
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

// 获取字符串变量
func (gh Controller) getParamWithString(key string) string {
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
func (gh *Controller) jsonOut(error int64, msg string, data interface{})  {
	var response global.JsonResponse
	response.Error = error
	if msg == "" {
		msg = "Success"
	}
	response.Msg = msg
	response.Data = data
	gh.Context.AbortWithStatusJSON(http.StatusOK, response)
}

func (gh *Controller) jsonOutByError(error int64, msg error, data interface{})  {
	gh.jsonOut(error, msg.Error(), data)
}