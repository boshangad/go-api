package public

import (
	"github.com/boshangad/v1/app/controller"
	"github.com/boshangad/v1/app/global"

	"github.com/boshangad/v1/services/appUserService"
)

type MpController struct {
}

// Login 微信code登录
// @route login [POST]
func (that MpController) Login(c *controller.Context) {
	var data = appUserService.MiniProgram{}
	if err := c.ShouldBind(&data); err != nil {
		c.JsonOut(global.ErrNotice, "missing parameter #code", nil)
		return
	}
	c.JsonOut(1000, "d", nil)

	return
	if err := data.Login(c); err != nil {
		c.JsonOut(global.ErrNotice, err.Error(), nil)
		return
	}
	c.JsonOut(global.ErrSuccess, "success", nil)
}

// SetUserProfile 设置用户信息
// @route user-profile [POST]
func (that MpController) SetUserProfile(c *controller.Context) {
	// var data = mpService.Profile{}
	// err := that.ShouldBind(&data)
	// if err != nil {
	// 	that.JsonOut(global.ErrNotice, err.Error(), nil)
	// 	return
	// }
	// err = data.SetAppUser(that.AppUser).Save()
	// if err != nil {
	// 	that.JsonOut(global.ErrNotice, err.Error(), nil)
	// 	return
	// }
	// that.JsonOut(global.ErrSuccess, "success", data)
}

// Info 获取用户信息
// @route info [GET]
func (that MpController) Info() {
	// that.JsonOut(global.ErrSuccess, "success", struct {
	// 	*ent.AppUser
	// }{AppUser: that.AppUser})
}
