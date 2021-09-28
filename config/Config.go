package config

type Server struct {
	App App
	// 系统配置
	System System
	// 数据库参数配置
	DB DB
	// Redis配置
	Redis map[string]Redis
	// 磁盘
	Disk Disk
	// zap日志
	Zap Zap
	// 短信服务
	Sms Sms
	// 邮件推送服务
	Email map[string]Email
	// 允许访问方式
	AsAccess AsAccess
	// 加密
	Paseto Paseto
	Jwt Jwt
}
