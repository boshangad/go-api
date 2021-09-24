package core

import (
	"flag"
	"fmt"
	viper2 "github.com/boshangad/go-api/core/viper"
	"github.com/boshangad/go-api/global"
	"github.com/fsnotify/fsnotify"
	jww "github.com/spf13/jwalterweatherman"
	"github.com/spf13/viper"
	"os"
	"path/filepath"
)

// Viper 初始化Viper配置加载
func Viper(path ...string) *viper.Viper {
	var config string
	if len(path) == 0 {
		flag.StringVar(&config, "c", "", "choose config file.")
		flag.Parse()
		// 优先级: 命令行 > 环境变量 > 默认值
		if config == "" {
			if configEnv := os.Getenv(viper2.ConfigEnv); configEnv == "" {
				config = viper2.ConfigFile
				fmt.Printf("您正在使用config的默认值,config的路径为%v\n", viper2.ConfigFile)
			} else {
				config = configEnv
				fmt.Printf("您正在使用G_CONFIG环境变量,config的路径为%v\n", config)
			}
		} else {
			fmt.Printf("您正在使用命令行的-c参数传递的值,config的路径为%v\n", config)
		}
	} else {
		config = path[0]
		fmt.Printf("您正在使用func Viper()传递的值,config的路径为%v\n", config)
	}
	jww.SetStdoutThreshold(jww.LevelTrace)
	v := viper.New()
	v.AddConfigPath(filepath.Dir(os.Args[0]))
	//v.SetConfigFile(config)
	v.SetConfigType("yaml")
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	// 监听配置文件修改
	v.WatchConfig()
	// 配置文件修改
	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config file changed:", e.Name)
		if err := v.Unmarshal(&global.G_CONFIG); err != nil {
			fmt.Println(err)
		}
	})
	if err := v.Unmarshal(&global.G_CONFIG); err != nil {
		fmt.Println(err)
	}
	global.G_CONFIG.App.Root, _ = filepath.Abs("..")
	return v
}
