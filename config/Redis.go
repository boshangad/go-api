package config

// Redis Redis服务
type Redis struct {
	Default string `mapstructure:"default" json:"default" yaml:"default"`
	Connections map[string]RedisParams `mapstructure:"connections" json:"connections" yaml:"connections"`
}

// RedisParams redis服务参数
type RedisParams struct {
	// redis的哪个数据库
	Db int `mapstructure:"db" json:"db" yaml:"db"`
	// 数据前缀
	Prefix string `mapstructure:"prefix" json:"prefix" yaml:"prefix"`
	// 服务器地址:端口
	Addr string `mapstructure:"addr" json:"addr" yaml:"addr"`
	// 密码
	Password string `mapstructure:"password" json:"password" yaml:"password"`
}