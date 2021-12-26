package smsLogService

import (
	"github.com/boshangad/v1/app/config"
	"github.com/boshangad/v1/app/global"
	"github.com/boshangad/v1/app/sms"
)

var smsGateway = sms.NewSmsGateway(global.Config.Sms)

func NewGateway(c config.Config) {
	
	smsGateway := sms.NewSmsGateway(c.Sms)
	c.AddObserver("smsGateway", smsGateway)

}

// 验证码
type ValidateCode struct {
	// 手机号
	Mobile string `json:"mobile" binding:"required"`
	// 国际区号
	DialCode string `json:"dial_code,omitempty" binding:""`
	// 可用范围
	Scope string `json:"scope,omitempty" binding:""`
}

// 发出验证码
func (ValidateCode) SendRandomCode() {

}
