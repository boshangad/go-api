package sms

import "github.com/boshangad/go-api/core/config/sms/gateways"

type Config struct {
	// 默认使用的配置项
	Default string `json:"default"`
	// 验证码最大验证次数
	CodeMaxCheckNumber int64 `json:"code_max_check_number,omitempty"`
	// 配置项
	Gateways map[string]*gateways.Config `json:"gateways,omitempty"`
}

func (that *Config) Init() {
	if that.Gateways == nil {
		that.Gateways = make(map[string]*gateways.Config)
	}
	for _, gateway := range that.Gateways {
		gateway.Init()
	}
}