package smsService

// sms服务应实现的接口
type smsInterface interface {
	SetAppId(appId uint64) smsInterface
	SetDialCode(dialCode string) smsInterface
	SetMobile(mobile string) smsInterface
	SetScope(scope string) smsInterface
	SetTypeId(typeId uint64) smsInterface
	SetGateway(gateway string) smsInterface
	SetIp(ip string) smsInterface
	SetTemplateId(templateId string) smsInterface
	SetTemplateText(templateText string) smsInterface
	SetContent(content map[string]interface{}) smsInterface
	SetCode(code string) smsInterface
	SetExpiredTime(expiredTime int64) smsInterface
	// Send 发送邮件
	Send() (err error)
	// Check 检查发出的代码
	Check(code string) (err error)
}

// SmsService 短信服务
type SmsService struct {
	smsInterface
	AppID uint64 `json:"app_id,omitempty"`
	DialCode string `json:"dial_code,omitempty"`
	Mobile string `json:"mobile,omitempty"`
	Scope string `json:"scope,omitempty"`
	TypeID uint64 `json:"type_id,omitempty"`
	Gateway string `json:"gateway,omitempty"`
	IP string `json:"ip,omitempty"`
	TemplateID string `json:"template_id,omitempty"`
	TemplateText string `json:"template_text,omitempty"`
	Content map[string]interface{} `json:"content,omitempty"`
	Code string `json:"code,omitempty"`
	ExpiredTime int64 `json:"expired_time,omitempty"`
}

func (that *SmsService) SetAppId(appId uint64) *SmsService {
	that.AppID = appId
	return that
}

func (that *SmsService) SetDialCode(dialCode string) *SmsService {
	that.DialCode = dialCode
	return that
}

func (that *SmsService) SetMobile(mobile string) *SmsService {
	that.Mobile = mobile
	return that
}

func (that *SmsService) SetScope(scope string) *SmsService {
	that.Scope = scope
	return that
}

func (that *SmsService) SetTypeId(typeId uint64) *SmsService {
	that.TypeID = typeId
	return that
}

func (that *SmsService) SetGateway(gateway string) *SmsService {
	that.Gateway = gateway
	return that
}

func (that *SmsService) SetIp(ip string) *SmsService {
	that.IP = ip
	return that
}

func (that *SmsService) SetTemplateId(templateId string) *SmsService {
	that.TemplateID = templateId
	return that
}

func (that *SmsService) SetTemplateText(templateText string) *SmsService {
	that.TemplateText = templateText
	return that
}

func (that *SmsService) SetContent(content map[string]interface{}) *SmsService {
	that.Content = content
	return that
}

func (that *SmsService) SetCode(code string) *SmsService {
	that.Code = code
	that.Content["code"] = code
	return that
}

func (that *SmsService) SetExpiredTime(expiredTime int64) *SmsService {
	that.ExpiredTime = expiredTime
	that.Content["expiredTime"] = expiredTime
	return that
}

//
//// NewDefaultGateWay 初始化默认推送网关
//func NewDefaultGateWay(key string) smsInterface {
//	sms := config.Get().Sms
//	if sms == nil {
//		panic("No mail service configuration was found#1")
//	}
//	if sms.Gateways == nil || len(sms.Gateways) < 1 {
//		panic("No mail service configuration was found#2")
//	}
//	if key == "" {
//		key = sms.Default
//	}
//	gatewayConfig, ok := sms.Gateways[key]
//	if !ok {
//		panic("")
//	}
//	f, ok := defaultContainerFuncs[gatewayConfig.Gateway]
//	if !ok {
//		panic("")
//	}
//	return f(key)
//}