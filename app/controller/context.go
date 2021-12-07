package controller

import (
	"fmt"
	"net/http"

	"github.com/boshangad/v1/app/errors"
	"github.com/boshangad/v1/app/global"
	"github.com/boshangad/v1/ent"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

const (
	// 应用key值
	ContextKeyByApp string = "$app"
	// 用户key值
	ContextKeyByUser string = "$user"
	// 应用用户key值
	ContextKeyByAppUser string = "$appUser"
	// 应用用户token的key值
	ContextKeyByAppUserToken string = "$appUserToken"
)

// 请求context
type Context struct {
	*gin.Context
}

// 获取登录的应用用户
func (that *Context) GetApp() (a *ent.App) {
	if val, ok := that.Get(ContextKeyByApp); ok && val != nil {
		a, _ = val.(*ent.App)
	}
	return
}

// 获取登录用户
func (that *Context) GetUser() (u *ent.User) {
	if val, ok := that.Get(ContextKeyByUser); ok && val != nil {
		u, _ = val.(*ent.User)
	}
	return
}

// 获取登录的应用用户
func (that *Context) GetAppUser() (au *ent.AppUser) {
	if val, ok := that.Get(ContextKeyByAppUser); ok && val != nil {
		au, _ = val.(*ent.AppUser)
	}
	return
}

// 获取登录的应用用户
func (that *Context) GetAppUserToken() (aut *ent.AppUserToken) {
	if val, ok := that.Get(ContextKeyByAppUserToken); ok && val != nil {
		aut, _ = val.(*ent.AppUserToken)
	}
	return
}

// 获取登录的应用用户
func (that *Context) SetApp(app *ent.App) {
	that.Set(ContextKeyByApp, app)
}

// 获取登录用户
func (that *Context) SetUser(user *ent.User) {
	that.Set(ContextKeyByUser, user)
}

// 获取登录的应用用户
func (that *Context) SetAppUser(appUser *ent.AppUser) {
	that.Set(ContextKeyByAppUser, appUser)
}

// 获取登录的应用用户
func (that *Context) SetAppUserToken(appUserToken *ent.AppUserToken) {
	that.Set(ContextKeyByAppUserToken, appUserToken)
}

// 绑定数据，允许多次调用
func (that *Context) ShouldBindValue(obj interface{}) (err error) {
	// 如果请求类型是JSOn,Xml 存在数据
	var (
		contentType = that.ContentType()
	)
	if that.Request.Method != http.MethodGet {
		switch contentType {
		case binding.MIMEJSON:
			return that.Context.ShouldBindBodyWith(obj, binding.JSON)
		case binding.MIMEXML, binding.MIMEXML2:
			return that.Context.ShouldBindBodyWith(obj, binding.XML)
		case binding.MIMEPROTOBUF:
			return that.Context.ShouldBindBodyWith(obj, binding.ProtoBuf)
		case binding.MIMEMSGPACK, binding.MIMEMSGPACK2:
			return that.Context.ShouldBindBodyWith(obj, binding.MsgPack)
		}
	}
	err = that.Context.ShouldBind(obj)
	return
}

// 输出json格式数据
func (that Context) JsonOut(error int64, msg string, data interface{}) {
	response := gin.H{}
	response["error"] = error
	if msg == "" {
		msg = "success"
	}
	response["msg"] = msg
	if data != nil {
		response["data"] = data
	}
	that.Context.AbortWithStatusJSON(http.StatusOK, response)
}

// 输出错误信息
func (that Context) JsonOutError(err error) {
	if err == nil {
		return
	}
	switch e := err.(type) {
	// case ContextError:
	// 	that.JsonOut(e.Code, e.Msg, e.Data)
	case validator.ValidationErrors:
		invalidParam := errors.NewInvalidParam()
		invalidParamData := errors.InvalidParamData{}
		for _, validationErr := range e {
			msg := ""
			if validationErr.Param() != "" {
				msg += ",the value is restricted to " + validationErr.Param() + "."
			}
			invalidParamData.Errors = append(invalidParamData.Errors, errors.InvalidError{
				Field: validationErr.Field(),
				Msg:   fmt.Sprintf("Field validation for '%s' failed on the '%s'%s", validationErr.Field(), validationErr.Tag(), msg),
			})
		}
		invalidParam.Data = &invalidParamData
		that.Context.AbortWithStatusJSON(invalidParam.GetStatus(), invalidParam)
	default:
		that.JsonOut(global.ErrNotice, err.Error(), nil)
	}
}
