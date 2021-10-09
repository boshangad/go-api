package public

import (
	"github.com/boshangad/go-api/core/mvvc"
	"github.com/boshangad/go-api/global"
	"github.com/boshangad/go-api/services/emailService"
	"github.com/boshangad/go-api/services/userService"
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
	var sendData = emailService.SendData{}
	err := that.ShouldBind(&sendData)
	if err != nil {
		that.JsonOut(global.ErrNotice, err.Error(), nil)
		return
	}
	if sendData.Email == "" {
		that.JsonOut(global.ErrNotice, "Mobile phone number cannot be empty.", nil)
		return
	}
	if !utils.ValidateEmail(sendData.Email) {
		that.JsonOut(global.ErrNotice, "Inaccurate mobile phone number format.", nil)
		return
	}
	// Check is need captcha
	if sendData.Captcha == "" {
		that.JsonOut(global.ErrNotice, "Miss captcha", nil)
		return
	}
	// send mobile
	if utils.InArrayWithString(sendData.Scope, []string{"login", "forget"}) {
		exist := userService.CheckAllowEmailLogin()
		if !exist {
			that.JsonOut(global.ErrNotice, "Sending failed, please try again later", nil)
			return
		}
	}
	err = emailService.SendCode(
		global.G_CONFIG.Email.Default,
		sendData.Email,
		that.Context.ClientIP(),
		sendData.Scope,
		that.App.ID,
	)
	if err != nil {
		that.JsonOutByError(global.ErrNotice, err, nil)
		return
	}
	that.JsonOut(global.ErrSuccess, "Success", nil)
}
