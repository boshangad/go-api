package mvvc

import (
	"github.com/boshangad/go-api/core/global"
	"github.com/boshangad/go-api/ent"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

// ControllerInterface 接口
type ControllerInterface interface {
}

type Controller struct {
	ControllerInterface
	// 请求对象
	Context *gin.Context
	RequestData map[string]interface{}
	// 应用用户
	App *ent.App
	// 应用用户
	AppUser *ent.AppUser
	// 用户
	User *ent.User
	// 应用用户登录token
	AppUserToken *ent.AppUserToken
}

// Init 初始化控制器数据
func (that *Controller) Init(c *gin.Context) *Controller {
	that.Context = c
	var (
		value interface{}
		ok bool
		requestData = make(map[string]interface{})
	)
	if value, ok = c.Get("App"); ok {
		that.App, ok = value.(*ent.App)
	}
	if value, ok = c.Get("AppUserToken"); ok {
		that.AppUserToken, ok = value.(*ent.AppUserToken)
	}
	if that.AppUserToken != nil {
		if value, ok = c.Get("AppUser"); ok {
			that.AppUser, ok = value.(*ent.AppUser)
		}
		if value, ok = c.Get("User"); ok {
			that.User, ok = value.(*ent.User)
		}
	}
	// 初始化用户入参
	that.RequestData = requestData
	return that
}

// GetParamWithString 获取用户入参
func (that Controller) GetParamWithString(key string) string {
	var (
		str string = ""
		method string = ""
	)
	method = strings.ToUpper(that.Context.Request.Method)
	if method == http.MethodGet {
		str = that.Context.DefaultQuery(key, "")
	} else if method != http.MethodOptions {
		str = that.Context.PostForm(key)
	}
	if str == "" {
		return that.Context.DefaultQuery(key, "")
	}
	return str
}

// JsonOut 输出json数据
func (that Controller) JsonOut(error int64, msg string, data interface{})  {
	var response global.JsonResponse
	response.Error = error
	if msg == "" {
		msg = "success"
	}
	response.Msg = msg
	if data != nil {
		response.Data = &data
	}
	that.Context.AbortWithStatusJSON(http.StatusOK, response)
	return
}

func (that Controller) JsonOutByError(error int64, msg error, data interface{})  {
	that.JsonOut(error, msg.Error(), data)
	return
}