package config

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/boshangad/go-api/core/config/email"
	"github.com/boshangad/go-api/core/config/sms"
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
	Mode string `json:"mode,omitempty"`
	App *appConfig `json:"app,omitempty"`
	// 数据库配置
	Db *dbConfig `json:"db,omitempty"`
	// 短信网关
	Sms *sms.Config `json:"sms,omitempty"`
	// 邮件推送
	EmailPush *email.Push `json:"email_push,omitempty"`
	// PASETO加密格式
	Paseto *PasetoConfig `json:"paseto,omitempty"`
	// 关于登录
	AsAccess *asAccess `json:"as access"`
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
}

// 初始化文件路径监听器
func initWatcher()  {
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

// 获取配置文件路径
func getConfigPath(filename string) (filepath string) {
	if filename == "" {
		filepath = global.GetPathWithExec() + "/config.json"
	} else {
		filepath = filename
		if !utils.IsFile(filepath) {
			filepath = global.GetPathWithExec() + "/" + strings.TrimLeft(filename, "/\\")
			if !utils.IsFile(filepath) {
				panic(errors.New("config file not found"))
			}
		}
	}
	return
}

func readConfigByFile(filepath string) (data []byte) {
	f, err := os.Open(filepath)
	if err != nil {
		panic(err)
	}
	defer func(oFile *os.File) { _ = oFile.Close() }(f)
	// 读取配置文件
	data, err = ioutil.ReadAll(f)
	if err != nil {

	}
	return
}

// LoadConfig 加载配置文件
func loadConfig(filename string) (config Config) {
	var (
		err error
		filepath string
	)
	if watcherFile != "" {
		_ = watcher.Remove(watcherFile)
	}
	filepath = getConfigPath(filename)
	watcherFile = filepath
	err = watcher.Add(filepath)
	if err != nil {
		log.Println("watcher add file fail", err)
	}
	config = Config{
		FilePath: filepath,
		Mode: "release",
		App: &appConfig{},
		Db: &dbConfig{},
		Paseto: &PasetoConfig{},
		Sms: &sms.Config{},
		EmailPush: &email.Push{},
		AsAccess: &asAccess{},
	}
	// 如果配置文件存在,则加载配置文件
	if utils.IsFile(filepath) {
		data := readConfigByFile(filepath)
		err = json.Unmarshal(data, &config)
		if err != nil {
			log.Panicln("json unmarshal fail", err)
		}
	}
	if strings.ToLower(config.Mode) != "release" {
		config.Mode = "debug"
	}

	config.App.Init()
	config.Db.Init()
	config.Paseto.Init()
	config.AsAccess.Init().Load()
	config.Sms.Init()
	return
}

// ReloadConfig 重新载入配置数据
func ReloadConfig(f string) {
	ch := make(chan bool)
	go func() {
		oldConfig := globalConfig
		globalConfig = loadConfig(f)
		// 关闭相关的数据库连接
		oldConfig.Db.CloseAllClient()
		ch <- true
	}()
	_ = <-ch
}

// SaveToDisk 保存配置到问卷
func SaveToDisk(filename string, config Config) {
	var d1, _ = json.MarshalIndent(config, "", "\t")
	if d1 != nil {
		// 写入文件(字节数组)
		err := ioutil.WriteFile(filename, d1, 0664)
		if err != nil {
			log.Println("write file fail", err)
		}
	}
}

// Get 返回配置数据
func Get() Config {
	return globalConfig
}