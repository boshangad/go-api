package config

type Email struct {
	// 收件人:多个以英文逗号分隔
	To       string `mapstructure:"to" json:"to" yaml:"to"`
	// 端口
	Port     int    `mapstructure:"port" json:"port" yaml:"port"`
	// 收件人
	From     string `mapstructure:"from" json:"from" yaml:"from"`
	// 服务器地址
	Host     string `mapstructure:"host" json:"host" yaml:"host"`
	// 是否SSL
	IsSSL    bool   `mapstructure:"is-ssl" json:"isSSL" yaml:"is-ssl"`
	// 密钥
	Secret   string `mapstructure:"secret" json:"secret" yaml:"secret"`
	// 昵称
	Nickname string `mapstructure:"nickname" json:"nickname" yaml:"nickname"`
}
