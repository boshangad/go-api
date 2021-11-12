package public

import (
	"reflect"
	"unsafe"

	"github.com/boshangad/go-api/core/mvvc"
	"github.com/boshangad/go-api/ent"
	"github.com/boshangad/go-api/global"
	"github.com/boshangad/go-api/services"
	"github.com/boshangad/go-api/services/appUserTokenService"
	"github.com/boshangad/go-api/services/mpService"
	"github.com/gin-gonic/gin"
)

type MpController struct {
	mvvc.Controller
}

// Login 微信code登录
// @route login [POST]
func (that MpController) Login(c *gin.Context) {
	that.Init(c)
	var data = mpService.Login{}
	if err := that.ShouldBind(&data); err != nil {
		that.JsonOut(global.ErrNotice, "missing parameter #code", nil)
		return
	}
	if data.Code == "" {
		that.JsonOut(global.ErrNotice, "missing parameter #code", nil)
		return
	}
	appUserModel, err := data.Login(that.App)
	if err != nil {
		return
	}
	model := appUserTokenService.NewModel()
	model.SetAppUser(appUserModel).CreateModel(that.Context)
	token := model.BuildAccessToken()

	that.JsonOut(global.ErrSuccess, "操作成功", services.StructLoginSuccess{
		AccessToken:       token,
		ExpireTime:        int64(that.AppUserToken.ExpireTime),
		IsLoadUserProfile: that.AppUser.IsLoadUserProfile,
		IsBindUser:        false,
	})
}

// SetUserProfile 设置用户信息
// @route user-profile [POST]
func (that MpController) SetUserProfile(c *gin.Context) {
	that.Init(c)
	var data = mpService.Profile{}
	err := that.ShouldBind(&data)
	if err != nil {
		that.JsonOut(global.ErrNotice, err.Error(), nil)
		return
	}
	err = data.SetAppUser(that.AppUser).Save()
	if err != nil {
		that.JsonOut(global.ErrNotice, err.Error(), nil)
		return
	}
	that.JsonOut(global.ErrSuccess, "success", data)
}

// Info 获取用户信息
// @route info [GET]
func (that MpController) Info() {
	that.JsonOut(global.ErrSuccess, "success", struct {
		*ent.AppUser
	}{AppUser: that.AppUser})
}