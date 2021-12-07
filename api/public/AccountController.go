package public

import (
	"github.com/boshangad/v1/app/controller"
	"github.com/boshangad/v1/app/global"
	"github.com/boshangad/v1/services/appUserLoginLogService"
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
	// 记录用户登录日志
	appUserLoginLog, err := appUserLoginLogService.CheckAndCreateLoginLogByConfirm(
		c.GetApp().ID,
		c.ClientIP(),
		c.Request,
	)
	if err != nil {
		c.JsonOutError(err)
		return
	}
	// 开启登录功能
	accessToken, err = userLogin.Login(appUserLoginLog)
	if err != nil {
		c.JsonOutError(err)
		return
	}
	c.JsonOut(global.ErrSuccess, "success", accessToken)
}

// Register 账号注册
// @route register [POST]
func (that AccountController) Register() {
	// var serviceWithUser = userService.UserRegisterParams{}
	// err := that.Context.ShouldBind(&serviceWithUser)
	// if err != nil {
	// 	that.JsonOutByError(global.ErrNotice, err, nil)
	// 	return
	// }
	// serviceWithUser.Filter()
	// if serviceWithUser.Username == "" && serviceWithUser.Mobile == "" && serviceWithUser.Email == "" {
	// 	that.JsonOut(global.ErrNotice, "用户名不能为空", nil)
	// 	return
	// }
	// userModel, err := serviceWithUser.Register(that.Controller)
	// if err != nil {
	// 	that.JsonOutByError(global.ErrNotice, err, nil)
	// 	return
	// }
	// that.JsonOut(global.ErrSuccess, "success", userModel)
}
