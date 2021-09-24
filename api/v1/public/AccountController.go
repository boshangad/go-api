package public

import (
	"github.com/boshangad/go-api/app/services"
	"github.com/boshangad/go-api/app/services/appUserTokenService"
	"github.com/boshangad/go-api/app/services/userService"
	"github.com/boshangad/go-api/core/mvvc"
	"github.com/boshangad/go-api/ent"
	"github.com/boshangad/go-api/global"
	"github.com/boshangad/go-api/utils"
)

type AccountController struct {
	mvvc.Controller
}

// Login 账号密码登录
// @route login [POST]
func (that *AccountController) Login()  {
	username := that.GetParamWithString("username")
	dailCode := that.GetParamWithString("dail_code")
	mobile := that.GetParamWithString("mobile")
	email := that.GetParamWithString("email")
	password := that.GetParamWithString("password")
	code := that.GetParamWithString("code")
	captcha := that.GetParamWithString("captcha")
	if password == "" && code == "" {
		that.JsonOut(global.ErrNotice, "Password cannot be empty.", nil)
		return
	}
	if captcha == "" {
		that.JsonOut(global.ErrNotice, "Captcha cannot be empty.", nil)
		return
	}
	var (
		err error
		userModel *ent.User
	)
	if mobile != "" {
		if dailCode == "" {
			dailCode = "86"
		}
		if dailCode == "86" && !utils.ValidateMobile(mobile) {
			that.JsonOut(global.ErrNotice, "Inaccurate mobile phone number format.", nil)
			return
		}
		if password != "" && code != "" {
			userService.LoginByMobileWithPasswordAndCode(dailCode, mobile, password, code)
		} else if code != "" {
			userModel, err = userService.LoginByMobileWithCode(dailCode, mobile, code)
		} else if password != "" {
			userModel, err = userService.LoginByMobileWithPassword(dailCode, mobile, password)
		}
	} else if email != "" {
		if !utils.ValidateEmail(email) {
			that.JsonOut(global.ErrNotice, "Inaccurate email format.", nil)
			return
		}
		if password != "" && code != "" {
			userService.LoginByEmailWithPasswordAndCode(email, password, code)
		} else if code != "" {
			userService.LoginByEmailWithCode(email, code)
		} else if password != "" {
			userService.LoginByEmailWithPassword(email, password)
		}
	} else if username != "" {
		userModel, err = userService.LoginByUsername(username, password)
	} else {
		that.JsonOut(global.ErrMissLoginParams, "Login params miss.", nil)
		return
	}
	// 登录失败报错
	if err != nil {
		that.JsonOutByError(global.ErrSuccess, err, nil)
		return
	}
	// 登录成功，创建token等信息
	err = userService.EventByLoginWithUser(userModel, that.Context)
	if err != nil {
		that.JsonOutByError(global.ErrNotice, err, nil)
		return
	}
	AppUserToken, ok := that.Context.Get("AppUserToken")
	if !ok {
		that.JsonOut(global.ErrNotice, "service exception, please try again", nil)
		return
	}
	AppUserTokenModel, ok := AppUserToken.(*ent.AppUserToken)
	if !ok {
		that.JsonOut(global.ErrNotice, "service exception, please try again", nil)
		return
	}
	token := appUserTokenService.CreateTokenWithModel(AppUserTokenModel)
	that.JsonOut(global.ErrSuccess, "success", services.StructLoginSuccess{
		AccessToken: token,
		ExpiredTime: AppUserTokenModel.ExpireTime,
		IsBindUser: true,
		UserAlias: userModel.UUID.String(),
	})
}

// Register 账号注册
// @route register [POST]
func (that AccountController) Register() {
	var serviceWithUser = userService.UserRegisterParams{}
	err := that.Context.ShouldBind(&serviceWithUser)
	if err != nil {
		that.JsonOutByError(global.ErrNotice, err, nil)
		return
	}
	serviceWithUser.Filter()
	if serviceWithUser.Username == "" && serviceWithUser.Mobile == "" && serviceWithUser.Email == "" {
		that.JsonOut(global.ErrNotice, "用户名不能为空", nil)
		return
	}
	userModel, err := serviceWithUser.Register(that.Controller)
	if err != nil {
		that.JsonOutByError(global.ErrNotice, err, nil)
		return
	}
	that.JsonOut(global.ErrSuccess, "success", userModel)
}