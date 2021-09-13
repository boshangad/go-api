package smsService

// GatewayInterface sms服务应实现的接口
type GatewayInterface interface {
	SetAppId(appId uint64) GatewayInterface
	SetDialCode(dialCode string) GatewayInterface
	SetMobile(mobile string) GatewayInterface
	SetScope(scope string) GatewayInterface
	SetTypeId(typeId uint64) GatewayInterface
	SetGateway(gateway string) GatewayInterface
	SetIp(ip string) GatewayInterface
	SetContent(content map[string]interface{}) GatewayInterface
	AppendContent(key string, data interface{}) GatewayInterface
	SetCode(code string) GatewayInterface
	SetExpiredTime(expiredTime int64) GatewayInterface
	// Send 发送邮件
	Send() (err error)
	// Check 检查发出的代码
	Check(code string) (err error)
}

// Config 短信配置服务
type Config struct {
	client interface{}
	TemplateID string `json:"template_id,omitempty"`
	TemplateText string `json:"template_text,omitempty"`
}

// Data 短信网关数据
type Data struct {
	Id uint64 `json:"id,omitempty"`
	AppID uint64 `json:"app_id,omitempty"`
	DialCode string `json:"dial_code,omitempty"`
	Mobile string `json:"mobile,omitempty"`
	Scope string `json:"scope,omitempty"`
	TypeID uint64 `json:"type_id,omitempty"`
	Gateway string `json:"gateway,omitempty"`
	IP string `json:"ip,omitempty"`
	Content map[string]interface{} `json:"content,omitempty"`
	Code string `json:"code,omitempty"`
	ExpiredTime int64 `json:"expired_time,omitempty"`
}

func (that *Data) SetAppId(appId uint64) *Data {
	that.AppID = appId
	return that
}

func (that *Data) SetDialCode(dialCode string) *Data {
	that.DialCode = dialCode
	return that
}

func (that *Data) SetMobile(mobile string) *Data {
	that.Mobile = mobile
	return that
}

func (that *Data) SetScope(scope string) *Data {
	that.Scope = scope
	return that
}

func (that *Data) SetTypeId(typeId uint64) *Data {
	that.TypeID = typeId
	return that
}

func (that *Data) SetGateway(gateway string) *Data {
	that.Gateway = gateway
	return that
}

func (that *Data) SetIp(ip string) *Data {
	that.IP = ip
	return that
}

func (that *Data) SetCode(code string) *Data {
	that.Code = code
	that.Content["code"] = code
	return that
}

func (that *Data) SetExpiredTime(expiredTime int64) *Data {
	that.ExpiredTime = expiredTime
	that.Content["expiredTime"] = expiredTime
	return that
}

func (that *Data) AppendContent(key string, data interface{}) *Data {
	that.Content[key] = data
	return that
}

// NewDefaultGateWay 初始化默认推送网关
func NewDefaultGateWay(key string) GatewayInterface {
	return nil
}