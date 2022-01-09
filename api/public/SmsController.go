package public

import (
	"github.com/boshangad/v1/app/controller"
	"github.com/boshangad/v1/app/global"
	"github.com/boshangad/v1/services/smsLogService"
)

type SmsController struct {
}

// Send 发出短信
func (that SmsController) Send(c *controller.Context) {
	var (
		validateCode = smsLogService.ValidateCode{}
	)
	if err := c.ShouldBindValue(&validateCode); err != nil {
		c.JsonOutError(err)
		return
	}
	if err := validateCode.SendRandomCode(c); err != nil {
		c.JsonOutError(err)
		return
	}
	c.JsonOut(global.ErrSuccess, "success", nil)
}
