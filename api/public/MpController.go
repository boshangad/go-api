package public

import (
	"github.com/boshangad/v1/app/controller"
	"github.com/boshangad/v1/global"

	"github.com/boshangad/v1/services/appUserService"
)

type MpController struct {
}

// Login 微信code登录
func (that MpController) Login(c *controller.Context) {
	var data = appUserService.MiniProgramLogin{}
	if err := c.ShouldBindValue(&data); err != nil {
		c.JsonOutError(err)
		return
	}
	if err := data.Login(c); err != nil {
		c.JsonOut(global.ErrNotice, err.Error(), nil)
		return
	}
	c.JsonOut(global.ErrSuccess, "success", nil)
}

// 绑定用户
func (that MpController) BindUser(c *controller.Context) {

}

// 绑定微信手机号
func (that MpController) BindMpMobile(c *controller.Context) {

}

// SetUserProfile 设置用户信息
func (that MpController) Profile(c *controller.Context) {
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
