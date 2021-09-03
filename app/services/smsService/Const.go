package smsService

const (
	SmsTypeSystem = 1
	SmsTypeNotify = 2
	SmsTypeLogin    = 3
	SmsTypeRegister = 3
	SmsTypeForget = 4
	SmsTypeSafe   = 5
)

// TypeCorrespondId 类型与id的对应关系
var TypeCorrespondId map[string]uint64

func init()  {
	TypeCorrespondId = map[string]uint64{
		"default":  SmsTypeSystem,
		"common":   SmsTypeNotify,
		"login":    SmsTypeLogin,
		"register": SmsTypeRegister,
		"forget":   SmsTypeForget,
		"safe":     SmsTypeSafe,
	}
}