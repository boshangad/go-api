package public

import (
	"github.com/boshangad/go-api/services/emailService"
	"github.com/boshangad/go-api/services/mpService"
	"github.com/boshangad/go-api/services/userService"
	"github.com/boshangad/go-api/core/mvvc"
	"github.com/boshangad/go-api/global"
	"github.com/boshangad/go-api/utils"
	"github.com/gin-gonic/gin"
)

type EmailController struct {
	mvvc.Controller
}

// Send 发出短信
// @router send [POST]
func (that EmailController) Send(c *gin.Context) {
	that.Init(c)
	var data = emailService.Data{}
	err := that.ShouldBind(&data)
	if err != nil {
		that.JsonOut(global.ErrNotice, err.Error(), nil)
		return
	}

	if email == "" {
		that.JsonOut(global.ErrNotice, "Mobile phone number cannot be empty.", nil)
		return
	}
	if !utils.ValidateEmail(email) {
		that.JsonOut(global.ErrNotice, "Inaccurate mobile phone number format.", nil)
		return
	}
	// Check is need captcha
	if captcha == "" {
		that.JsonOut(global.ErrNotice, "Miss captcha", nil)
		return
	}
	// send mobile
	if useType == "login" || useType == "forget" {
		exist := userService.CheckAllowEmailLogin()
		if !exist {
			that.JsonOut(global.ErrNotice, "Sending failed, please try again later", nil)
			return
		}
	}
	err := emailService.SendCode(that.App.ID, useType, email, that.Context.ClientIP())
	if err != nil {
		that.JsonOutByError(global.ErrNotice, err, nil)
		return
	}
	that.JsonOut(global.ErrSuccess, "Success", nil)
}
