package config

// Sms 短信
type Sms struct {
	// 默认网关
	Default string
	// 网关配置列表
	Gateways map[string]SmsParams
}

// SmsParams 网关配置参数
type SmsParams struct {
	// 网关名称有aliyun、ucloud等
	Gateway string
	// 签名名称
	SignName string
	// 区域
	RegionId string
	// 访问密钥
	AccessKey string
	// 访问公钥
	AccessSecret string
	// 短信模板
	Templates []SmsTemplate
}

// SmsTemplate 短信消息模板
type SmsTemplate struct {
	// 使用范围
	Scopes []string
	// 模板ID
	TemplateId string
	// 模板内容
	TemplateText string
}