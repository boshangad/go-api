package controllers

import (
	"github.com/tend/wechatServer/app/services"
	"github.com/tend/wechatServer/app/services/appUserTokenService"
	"github.com/tend/wechatServer/app/services/userService"
	"github.com/tend/wechatServer/core/global"
	"github.com/tend/wechatServer/ent"
	"github.com/tend/wechatServer/utils"
)

type AccountController struct {
	Controller
}

// Login 账号密码登录
// @route login [POST]
func (c *AccountController) Login()  {
	username := c.getParamWithString("username")
	dailCode := c.getParamWithString("dail_code")
	mobile := c.getParamWithString("mobile")
	email := c.getParamWithString("email")
	password := c.getParamWithString("password")
	code := c.getParamWithString("code")
	captcha := c.getParamWithString("captcha")
	if password == "" && code == "" {
		c.jsonOut(global.ErrNotice, "Password cannot be empty.", nil)
		return
	}
	if captcha == "" {
		c.jsonOut(global.ErrNotice, "Captcha cannot be empty.", nil)
		return
	}
	var err error
	var userModel *ent.User
	if mobile != "" {
		if dailCode == "" {
			dailCode = "86"
		}
		if dailCode == "86" && !utils.ValidateMobile(mobile) {
			c.jsonOut(global.ErrNotice, "Inaccurate mobile phone number format.", nil)
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
			c.jsonOut(global.ErrNotice, "Inaccurate email format.", nil)
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
		c.jsonOut(global.ErrMissLoginParams, "Login params miss.", nil)
		return
	}
	// 登录失败报错
	if err != nil {
		c.jsonOutByError(global.ErrSuccess, err, nil)
		return
	}
	// 登录成功，创建token等信息
	err = userService.EventByLoginWithUser(userModel, c.Context)
	if err != nil {
		c.jsonOutByError(global.ErrNotice, err, nil)
		return
	}
	AppUserToken, ok := c.Context.Get("AppUserToken")
	if !ok {
		c.jsonOut(global.ErrNotice, "service exception, please try again", nil)
		return
	}
	AppUserTokenModel, ok := AppUserToken.(*ent.AppUserToken)
	if !ok {
		c.jsonOut(global.ErrNotice, "service exception, please try again", nil)
		return
	}
	token := appUserTokenService.CreateTokenWithModel(AppUserTokenModel)
	c.jsonOut(global.ErrSuccess, "success", services.StructLoginSuccess{
		AccessToken: token,
		ExpiredTime: AppUserTokenModel.ExpireTime,
		IsBindUser: true,
		UserAlias: userModel.UUID,
	})
}

// Register 账号注册
// @route register [POST]
func (c *AccountController) Register() {
	//username := c.getParamWithString("username")
	//mobile := c.getParamWithString("mobile")
	//email := c.getParamWithString("email")

}