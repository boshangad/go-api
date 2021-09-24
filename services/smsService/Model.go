package smsService

const (
	SmsTypeSystem = 1
	SmsTypeNotify = 2
	SmsTypeLogin    = 3
	SmsTypeRegister = 3
	SmsTypeForget = 4
	SmsTypeSafe   = 5
)

const (
	// StatusDraft 未发布
	StatusDraft = 0
	// StatusUsed 已核销
	StatusUsed = 1
	// StatusPublished 已发布
	StatusPublished = 2
	// StatusExpire 已失效
	StatusExpire = 3
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