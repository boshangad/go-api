package config

type Server struct {
	App App
	// 系统配置
	System System
	// 数据库参数配置
	DB DB
	// zap日志
	Zap Zap
	// Redis配置
	Redis map[string]Redis
	// 短信服务
	Sms Sms
	// 邮件推送服务
	Email map[string]Email
	// 允许访问方式
	AsAccess AsAccess
}
