package config

// 短信网关配置
type Sms struct {
	// 默认使用网关
	Defaults []string `json:"defaults,omitempty" yaml:"defaults"`
	// 网关配置项
	Gateways map[string]map[string]interface{} `json:"gateways,omitempty" yaml:"gateways"`
}
