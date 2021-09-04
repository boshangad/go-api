package controllers

import (
	"github.com/boshangad/go-api/app/services/smsService"
	"github.com/boshangad/go-api/app/services/userService"
	"github.com/boshangad/go-api/core/controller"
	"github.com/boshangad/go-api/core/global"
	"github.com/boshangad/go-api/utils"
)

type SmsController struct {
	controller.Controller
}

// Send 发出短信
// @router send [POST]
func (that SmsController) Send() {
	mobile := that.GetParamWithString("mobile")
	dialCode := that.GetParamWithString("dial_code")
	useType := that.GetParamWithString("use_type")
	captcha := that.GetParamWithString("captcha")

	if dialCode == "" {
		dialCode = "86"
	}
	if mobile == "" {
		that.JsonOut(global.ErrNotice, "Mobile phone number cannot be empty.", nil)
		return
	}
	if dialCode == "86" {
		if !utils.ValidateMobile(mobile) {
			that.JsonOut(global.ErrNotice, "Inaccurate mobile phone number format.", nil)
			return
		}
	}
	// Check is need captcha
	if captcha == "" {
		that.JsonOut(global.ErrNotice, "Miss captcha", nil)
		return
	}
	var (
		typeId uint64
		ok bool
	)
	if typeId, ok = smsService.TypeCorrespondId[useType]; !ok {
		that.JsonOut(global.ErrNotice, "Sending failed, undeclared type", nil)
		return
	}
	// send mobile
	if useType == "login" || useType == "forget" {
		exist := userService.CheckIsExistByMobile(dialCode, mobile)
		if !exist {
			that.JsonOut(global.ErrNotice, "Sending failed, please try again later", nil)
			return
		}
	}
	err := smsService.Send(mobile, dialCode, typeId)
	if err != nil {
		that.JsonOutByError(global.ErrNotice, err, nil)
	}

	that.JsonOut(global.ErrSuccess, "Success", nil)
}
