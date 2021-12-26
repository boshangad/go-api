package config

import (
	"os"
	"path/filepath"
	"strings"

	zapLog "github.com/boshangad/v1/app/log"
	"github.com/boshangad/v1/app/validators"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

type Viper struct {
	viper *viper.Viper
	// 日志配置
	logger *zap.Logger
	// 配置监听观察者
	notifyerFuncs map[string]NotifyerViperFunc
	notifyersKeys []string
}

// 重新加载配置
func (that *Viper) Reload() {
	// 不允许崩溃，应写入日志，这个日志怎么处理呢
	defer func() {
		if err := recover(); err != nil {
			that.logger.DPanic("config reload panic", zap.Any("recover", err))
		}
	}()
	var err = that.viper.ReadInConfig()
	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// 配置文件未找到错误；如果需要可以忽略
			that.logger.Warn("config file not found", zap.Error(err))
			return
		} else {
			// 配置文件被找到，但产生了另外的错误
			that.logger.Error("Fatal error config file", zap.String("filename", that.viper.ConfigFileUsed()), zap.Error(err))
			return
		}
	}
	if that.notifyerFuncs != nil {
		for _, k := range that.notifyersKeys {
			if fn, ok := that.notifyerFuncs[k]; ok {
				fn(that)
			} else {
				that.DelObserver(k)
			}
		}
	}
}

// 添加观察者
func (that *Viper) AddObserver(k string, n NotifyerViperFunc) {
	if that.notifyerFuncs == nil {
		that.notifyerFuncs = make(map[string]NotifyerViperFunc)
	}
	if n == nil {
		that.DelObserver(k)
		return
	}
	if _, ok := that.notifyerFuncs[k]; !ok {
		that.notifyersKeys = append(that.notifyersKeys, k)
	}
	that.notifyerFuncs[k] = n
}

// 删除监听者
func (that *Viper) DelObserver(k string) {
	delete(that.notifyerFuncs, k)
	var keys = []string{}
	for _, k := range that.notifyersKeys {
		if _, ok := that.notifyerFuncs[k]; ok {
			keys = append(keys, k)
		}
	}
	that.notifyersKeys = keys
}

// 获取配置器
func (that *Viper) Viper() *viper.Viper {
	return that.viper
}

// 实例化配置器
func NewViper(filename string) *Viper {
	var (
		vp = Viper{
			viper:         viper.New(),
			notifyerFuncs: make(map[string]NotifyerViperFunc),
			notifyersKeys: []string{},
			logger:        zap.NewExample(),
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
			vp.logger.Warn("config file not found", zap.Error(err))
		} else {
			// 配置文件被找到，但产生了另外的错误
			vp.logger.Fatal("Fatal error config file", zap.Error(err))
		}
	}
	// 观察者监听
	vp.viper.OnConfigChange(func(e fsnotify.Event) {
		vp.Reload()
	})
	vp.viper.WatchConfig()
	return &vp
}
