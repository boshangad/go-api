package public

import (
	"github.com/boshangad/v1/app/controller"
	"github.com/boshangad/v1/app/global"
)

type SmsController struct {
}

// Send 发出短信
func (that SmsController) Send(c *controller.Context) {
	// var (
	// 	codeData = smsLogService.SmsLog{}
	// )
	// if err := c.ShouldBindValue(&codeData); err != nil {
	// 	c.JsonOutError(err)
	// 	return
	// }
	// if err := codeData.Send(); err != nil {
	// 	c.JsonOutError(err)
	// 	return
	// }
	c.JsonOut(global.ErrSuccess, "success", nil)
}
