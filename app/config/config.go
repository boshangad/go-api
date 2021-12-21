package config

import (
	"log"

	smsConfig "github.com/boshangad/v1/app/sms/config"

	zapLog "github.com/boshangad/v1/app/log"
)

type Config struct {
	// 应用配置
	App App `json:"app,omitempty" yaml:"app,omitempty"`
	// 日志
	Log zapLog.Zap `json:"log,omitempty" yaml:"log"`
	// 数据库连接配置
	Db map[string]interface{} `json:"db,omitempty" yaml:"db"`
	// 缓存功能模块
	Cache map[string]interface{} `json:"cache,omitempty" yaml:"cache"`
	// Redis 服务器
	Redis map[string]interface{} `json:"redis,omitempty" yaml:"redis"`
	// 短信网关
	Sms *smsConfig.Config `json:"sms,omitempty" yaml:"sms"`
	// 邮件网关
	Email map[string]interface{} `json:"email,omitempty" yaml:"email"`
	// 验证码配置
	Captcha map[string]Captcha `json:"captcha,omitempty" yaml:"captcha"`
}

// 刷新回调
func (that *Config) Callback(v *Viper) {
	// map 数据不会被清除，需要清除
	that.Db = make(map[string]interface{})
	that.Cache = make(map[string]interface{})
	that.Redis = make(map[string]interface{})
	that.Email = make(map[string]interface{})
	that.Captcha = make(map[string]Captcha)
	err := v.viper.Unmarshal(&that)
	if err != nil {
		log.Println("config file not found: " + err.Error() + "\n")
	}
}

func NewConfig(v *Viper) *Config {
	config := Config{}
	err := v.viper.Unmarshal(&config)
	if err != nil {
		log.Println("config file not found: " + err.Error() + "\n")
	}
	v.AddObserver("config", &config)
	return &config
}
