package config

import (
	"log"
	"os"
	"path/filepath"

	"github.com/boshangad/v1/app/helpers"
	zapLog "github.com/boshangad/v1/app/log"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type Config struct {
	// viper 配置器
	viper *viper.Viper
	// 配置监听观察者
	notifyers map[string]Notifyer
	// 配置文件完整名称
	ConfigFile string
	// 配置文件类型
	ConfigType string
	// 配置文件名称，不带后缀名和路径
	ConfigName string
	// ----------------------------------------------------------------
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
	Sms map[string]interface{} `json:"sms,omitempty" yaml:"sms"`
	// 邮件网关
	Email map[string]interface{} `json:"email,omitempty" yaml:"email"`
	// 验证码配置
	Captcha map[string]Captcha `json:"captcha,omitempty" yaml:"captcha"`
}

// 加载配置
func (that *Config) Load() *Config {
	// 加载配置文件
	viperClient := viper.New()
	if that.ConfigName != "" {
		viperClient.SetConfigName(that.ConfigName)
	}
	if that.ConfigType != "" {
		viperClient.SetConfigType(that.ConfigType)
	}
	if that.ConfigFile != "" {
		viperClient.SetConfigFile(that.ConfigFile)
	}
	// 配置器
	viperClient.AddConfigPath(".")
	// 读取配置文件
	err := viperClient.ReadInConfig()
	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// 配置文件未找到错误；如果需要可以忽略
			log.Println("config file not found: " + err.Error() + "\n")
			// 设置观察者对象
			that.viper = viperClient
			return that
		}
		// 配置文件被找到，但产生了另外的错误
		log.Panicln("Fatal error config file: " + err.Error() + "\n")
	}
	err = viperClient.Unmarshal(&that)
	if err != nil {
		log.Println("config file not found: " + err.Error() + "\n")
		// 设置观察者对象
		that.viper = viperClient
		return that
	}
	// 移除字符串左右的空格
	helpers.TrimSpace(that)
	// 观察者监听
	viperClient.OnConfigChange(func(e fsnotify.Event) {
		log.Println("Config file changed:", e.Name)
		if that.notifyers != nil {
			for _, v := range that.notifyers {
				v.Callback(that)
			}
		}
	})
	viperClient.WatchConfig()
	// 设置观察者对象
	that.viper = viperClient
	return that
}

// 重新加载配置
func (that *Config) Reload() {

}

// 添加观察者
func (that *Config) AddObserver(k string, n Notifyer) {
	that.notifyers[k] = n
}

// 获取配置器
func (that Config) Viper() *viper.Viper {
	return that.viper
}

// 默认配置
func DefaultConfig() *Config {
	return &Config{
		notifyers: make(map[string]Notifyer),
		App: App{
			Debug:    false,
			Name:     "",
			Listen:   ":80",
			RootPath: filepath.Dir(os.Args[0]),
			BaseUrl:  "http://127.0.0.1",
			Cors:     Cors{},
		},
		Log: zapLog.DefaultZapConfig(),
	}
}

// 实例化配置
func New() *Config {
	return &Config{}
}
