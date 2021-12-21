package config

import (
	"log"
	"os"
	"path/filepath"
	"strings"

	zapLog "github.com/boshangad/v1/app/log"
	"github.com/boshangad/v1/app/validators"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type Viper struct {
	viper *viper.Viper
	// 配置监听观察者
	notifyers map[string]Notifyer
}

// 重新加载配置
func (that *Viper) Reload() {
	err := that.viper.ReadInConfig()
	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// 配置文件未找到错误；如果需要可以忽略
			log.Println("config file not found: " + err.Error() + "\n")
			return
		} else {
			// 配置文件被找到，但产生了另外的错误
			log.Println("Fatal error config file: " + err.Error() + "\n")
			return
		}
	}
}

// 添加观察者
func (that *Viper) AddObserver(k string, n Notifyer) {
	that.notifyers[k] = n
}

// 获取配置器
func (that *Viper) Viper() *viper.Viper {
	return that.viper
}

// 实例化配置器
func NewViper(filename string) *Viper {
	var (
		vp = Viper{
			viper:     viper.New(),
			notifyers: make(map[string]Notifyer),
		}
	)
	// 设置默认值
	vp.viper.SetDefault("app", App{
		Debug:    false,
		Name:     "adboshang",
		Listen:   ":80",
		RootPath: filepath.Dir(os.Args[0]),
		BaseUrl:  "http://127.0.0.1",
		Cors:     Cors{},
	})
	vp.viper.SetDefault("log", zapLog.DefaultZapConfig())

	// 开始加载读取配置
	filename = strings.TrimRight(strings.TrimSpace(filename), "/\\")
	if validators.IsUrl(filename) {
		// 远端配置文件
	} else if strings.ContainsAny(filename, "/\\") {
		// 文件路径
		vp.viper.AddConfigPath(filepath.Dir(filename))
		vp.viper.SetConfigFile(filepath.Base(filename))
	} else {
		// 文件名称
		vp.viper.AddConfigPath(".")
		vp.viper.SetConfigName(filename)
	}
	err := vp.viper.ReadInConfig()
	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// 配置文件未找到错误；如果需要可以忽略
			log.Println("config file not found: " + err.Error() + "\n")
		} else {
			// 配置文件被找到，但产生了另外的错误
			log.Panicln("Fatal error config file: " + err.Error() + "\n")
		}
	}
	// 观察者监听
	vp.viper.OnConfigChange(func(e fsnotify.Event) {
		// 不允许崩溃，应写入日志，这个日志怎么处理呢
		defer func() {
			if err := recover(); err != nil {
				log.Println(err)
			}
		}()
		vp.Reload()
		if vp.notifyers != nil {
			for _, v := range vp.notifyers {
				v.Callback(&vp)
			}
		}
	})
	vp.viper.WatchConfig()
	return &vp
}
