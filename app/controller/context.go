package controller

import (
	"net/http"

	"github.com/boshangad/v1/ent"
	"github.com/gin-gonic/gin"
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
