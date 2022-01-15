package config

import (
	"fmt"

	smsConfig "github.com/boshangad/v1/app/sms/config"
	"go.uber.org/zap"

	zapLog "github.com/boshangad/v1/app/log"
)

type Config struct {
	logger *zap.Logger `json:"-" yaml:"-"`
	// 配置监听观察者
	notifyers map[string]NotifyerFunc `json:"-" yaml:"-"`
	// 应用配置
	App App `json:"app,omitempty" yaml:"app,omitempty"`
	// 日志
	Log zapLog.Zap `json:"log,omitempty" yaml:"log"`
	// 数据库连接配置
	Db map[string]interface{} `json:"db,omitempty" yaml:"db"`
	// 磁盘
	Disk *Disk `json:"disk,omitempty" yaml:"disk"`
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

// 添加观察者
func (that *Config) AddObserver(k string, n NotifyerFunc) {
	that.notifyers[k] = n
}

// 刷新回调
func (that *Config) Reload(v *Viper) {
	// map 数据不会被清除，需要清除
	that.Db = make(map[string]interface{})
	that.Cache = make(map[string]interface{})
	that.Redis = make(map[string]interface{})
	that.Email = make(map[string]interface{})
	that.Captcha = make(map[string]Captcha)
	err := v.viper.Unmarshal(&that)
	if err != nil {
		v.logger.Error("config file not found", zap.Error(err))
		return
	}
	that.logger = v.GetLogger()
	// 重新调用回调
	if that.notifyers != nil {
		for _, fn := range that.notifyers {
			fn(that)
		}
	}
}

// 实例化配置
func NewConfig(v *Viper) *Config {
	config := Config{
		logger:    v.logger,
		notifyers: make(map[string]NotifyerFunc),
	}
	err := v.viper.Unmarshal(&config)
	if err != nil {
		v.logger.Fatal("config file not found", zap.Error(err))
		return &config
	}
	v.AddObserver("config:"+fmt.Sprintf("%p", &config), func(v *Viper) {
		config.Reload(v)
	})
	return &config
}
