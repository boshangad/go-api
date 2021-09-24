package config

type System struct {
	// 环境值
	Env string `mapstructure:"env" json:"env" yaml:"env"`
	// 端口值
	Addr int `mapstructure:"addr" json:"addr" yaml:"addr"`
	// 多点登录拦截
	UseMultipoint bool `mapstructure:"use-multipoint" json:"useMultipoint" yaml:"use-multipoint"`
	// Token生成方式
	UseTokenType string `mapstructure:"use-token-type" json:"useTokenType" yaml:"use-token-type"`
}