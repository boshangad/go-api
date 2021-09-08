package config

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/boshangad/go-api/core/global"
	"github.com/boshangad/go-api/utils"
	"github.com/fsnotify/fsnotify"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

var (
	// 全局配置
	globalConfig Config
	// 监听的配置文件
	watcherFile string
	// 监听器
	watcher *fsnotify.Watcher
)

type Config struct {
	// 配置文件路径
	FilePath string `json:"-"`
	// 应用模式
	Mode string `json:"mode"`
	App appConfig `json:"app"`
	// 数据库配置
	Db dbConfig `json:"db,omitempty"`
	// 短信网关
	Sms *sms `json:"sms,omitempty"`
	// 邮件推送
	EmailPush *emailPush `json:"email_push,omitempty"`
	// PASETO加密格式
	Paseto PasetoConfig `json:"paseto,omitempty"`
	// 无需登录的接口地址
	NoAccess map[string]interface{} `json:"no access"`
}

// LoadConfig 加载配置文件
func loadConfig(configFile string) Config {
	if watcherFile != "" {
		_ = watcher.Remove(watcherFile)
	}
	var requiredCreateFile bool = false
	if configFile == "" {
		configFile = global.GetPathWithExec() + "/config.json"
		if !utils.IsFile(configFile) {
			requiredCreateFile = true
		}
	} else if !utils.IsFile(configFile) {
		fp := global.GetPathWithExec() + "/" + configFile
		if !utils.IsFile(fp) {
			panic(errors.New("config file not found"))
		}
		configFile = fp
	}
	watcherFile = configFile
	err := watcher.Add(configFile)
	if err != nil {
		log.Println("watcher add file fail", err)
	}
	loadConfig := Config{
		FilePath: configFile,
		Mode: "release",
	}
	// 如果配置文件存在,则加载配置文件
	if utils.IsFile(configFile) {
		//requiredCreateFile = false
		f, err := os.Open(configFile)
		if err != nil {
			panic(err)
		}
		defer func(oFile *os.File) { _ = oFile.Close() }(f)
		// 读取配置文件
		byteValue, _ := ioutil.ReadAll(f)
		err = json.Unmarshal(byteValue, &loadConfig)
		if err != nil {
			panic(err)
		}
	}
	if strings.ToLower(loadConfig.Mode) != "debug" {
		loadConfig.Mode = "release"
	}
	loadConfig.App.initDefaultData()
	loadConfig.Db.initDefaultData()
	loadConfig.Paseto.initDefaultData()
	if loadConfig.NoAccess == nil {
		loadConfig.NoAccess = make(map[string]interface{})
	}
	// 新增文件
	if requiredCreateFile {
		var d1, _ = json.MarshalIndent(loadConfig, "", "\t")
		if d1 != nil {
			// 写入文件(字节数组)
			_ = ioutil.WriteFile(configFile, d1, 0664)
		}
	}
	return loadConfig
}

// ReloadConfig 重新载入配置数据
func ReloadConfig(f string) {
	if f == "" {}
	globalConfig = loadConfig(f)
	global.CasbinAuthRequiredLogin.LoadNoAccess(Get().NoAccess)
}

// Get 返回配置数据
func Get() Config {
	return globalConfig
}

// 初始化数据
func init() {
	var err error
	watcher, err = fsnotify.NewWatcher()
	if err != nil {
		log.Fatalf("error: fsnotify watcher create fail %s", err)
	}
	// 加载文件
	globalConfig = loadConfig("")
	// 执行监听的任务
	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				if event.Op & fsnotify.Write == fsnotify.Write {
					fmt.Println("重新加载", event.Name, "配置")
					ReloadConfig(event.Name)
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				log.Println("config watcher error:", err)
			}
		}
	}()
}