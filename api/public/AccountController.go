package public

import (
	"github.com/boshangad/v1/app/controller"
	"github.com/boshangad/v1/global"
	"github.com/boshangad/v1/services/appUserTokenService"
	"github.com/boshangad/v1/services/userService"
)

type AccountController struct {
}

// Login 账号密码登录
// @route login [POST]
func (that AccountController) Login(c *controller.Context) {
	var (
		userLogin = userService.Login{}
		// 错误信息
		err error
		// 登录成功的令牌
		accessToken *appUserTokenService.AccessToken
	)
	if err = c.ShouldBindValue(&userLogin); err != nil {
		c.JsonOutError(err)
		return
	}
	// 开启登录功能
	accessToken, err = userLogin.Login(c)
	if err != nil {
		c.JsonOutError(err)
		return
	}
	c.JsonOut(global.ErrSuccess, "success", accessToken)
}

// Register 账号注册
// @route register [POST]
func (that AccountController) Register(c *controller.Context) {
	var register = userService.Register{}
	err := c.ShouldBind(&register)
	if err != nil {
		c.JsonOutError(err)
		return
	}
	userModel, err := register.Register(c)
	if err != nil {
		c.JsonOutError(err)
		return
	}
	c.JsonOut(global.ErrSuccess, "success", userModel)
}
