package smsLogService

import "github.com/boshangad/v1/app/global"

var Sms = NewSmsGateway(&global.Config.Sms)
