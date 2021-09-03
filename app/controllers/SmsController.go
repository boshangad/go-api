package controllers

import (
	"github.com/tend/wechatServer/app/services/smsService"
	"github.com/tend/wechatServer/app/services/userService"
	"github.com/tend/wechatServer/core/global"
	"github.com/tend/wechatServer/utils"
)

type smsController struct {
	Controller
}

// Send 发出短信
// @router send [POST]
func (c smsController) Send() {
	mobile := c.getParamWithString("mobile")
	dialCode := c.getParamWithString("dial_code")
	useType := c.getParamWithString("use_type")
	captcha := c.getParamWithString("captcha")

	if dialCode == "" {
		dialCode = "86"
	}
	if mobile == "" {
		c.jsonOut(global.ErrNotice, "Mobile phone number cannot be empty.", nil)
		return
	}
	if dialCode == "86" {
		if !utils.ValidateMobile(mobile) {
			c.jsonOut(global.ErrNotice, "Inaccurate mobile phone number format.", nil)
			return
		}
	}
	// Check is need captcha
	if captcha == "" {
		c.jsonOut(global.ErrNotice, "Miss captcha", nil)
		return
	}
	var (
		typeId uint64
		ok bool
	)
	if typeId, ok = smsService.TypeCorrespondId[useType]; !ok {
		c.jsonOut(global.ErrNotice, "Sending failed, undeclared type", nil)
		return
	}
	// send mobile
	if useType == "login" || useType == "forget" {
		exist := userService.CheckIsExistByMobile(dialCode, mobile)
		if !exist {
			c.jsonOut(global.ErrNotice, "Sending failed, please try again later", nil)
			return
		}
	}
	err := smsService.Send(mobile, dialCode, typeId)
	if err != nil {
		c.jsonOutByError(global.ErrNotice, err, nil)
	}

	c.jsonOut(global.ErrSuccess, "Success", nil)
}
