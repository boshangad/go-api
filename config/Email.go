package config

// 邮箱推送配置结构
type Email struct {
	Defaults []string                          `mapstructure:"defaults" json:"defaults" yaml:"defaults"`
	Gateways map[string]map[string]interface{} `mapstructure:"gateways" json:"gateways" yaml:"gateways"`
}
