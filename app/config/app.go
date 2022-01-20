package config

import "github.com/boshangad/v1/app/helpers"

// 应用结构
type App struct {
	// 是否调试模式
	Debug bool `mapstructure:"debug" json:"debug,omitempty" yaml:"debug"`
	// 应用名称
	Name string `mapstructure:"name" json:"name,omitempty" yaml:"name"`
	// 监听地址
	Listen string `mapstructure:"listen" json:"listen,omitempty" yaml:"listen"`
	// 应用根路径
	RootPath string `mapstructure:"rootPath" json:"rootPath,omitempty" yaml:"rootPath"`
	// 访问路径
	BaseUrl string `mapstructure:"baseUrl" json:"baseUrl,omitempty" yaml:"baseUrl"`
	// 密码hash等级 4~31之间，默认为10
	PasswdLevel helpers.Int `mapstructure:"passwdLevel" json:"passwdLevel,omitempty" yaml:"passwdLevel"`
	// Cors 跨域
	Cors Cors `mapstructure:"cors" json:"cors,omitempty" yaml:"cors"`
}
