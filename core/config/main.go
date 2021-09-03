package config

import (
	"encoding/json"
	"errors"
	"github.com/tend/wechatServer/core/global"
	"github.com/tend/wechatServer/utils"
	"io/ioutil"
	"os"
	"strings"
)

var globalConfig Config

type Config struct {
	Mode string `json:"mode"`
	App appConfig `json:"app"`
	// 数据库配置
	Db dbConfig `json:"db,omitempty"`
	// PASETO加密格式
	Paseto PasetoConfig `json:"paseto,omitempty"`
	// 无需登录的接口地址
	NoAccess map[string]interface{} `json:"no_access,omitempty"`
}

// 初始化数据
func init() {
	globalConfig = loadConfig("")
}

// LoadConfig 加载配置文件
func loadConfig(configFile string) Config {
	var requiredCreateFile bool = false
	if configFile == "" {
		configFile= global.GetPathWithExec() + "/config.json"
		requiredCreateFile = true
	} else if !utils.IsFile(configFile) {
		fp := global.GetPathWithExec() + "/" + configFile
		if !utils.IsFile(fp) {
			panic(errors.New("config file not found"))
		}
		configFile = fp
	}
	loadConfig := Config{
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
	// 新增文件
	if requiredCreateFile {
		var d1, _ = json.MarshalIndent(loadConfig, "", "\t")
		if d1 != nil {
			// 写入文件(字节数组)
			_ = ioutil.WriteFile(configFile, d1, 0666)
		}
	}
	return loadConfig
}

// ReloadConfig 重新载入配置数据
func ReloadConfig()  {
	globalConfig = loadConfig("")
}

// Get 返回配置数据
func Get() Config {
	return globalConfig
}