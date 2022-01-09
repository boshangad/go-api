package config

import (
	"github.com/boshangad/v1/app/log"
	"go.uber.org/zap"
)

// 实例化日志配置
func NewLog(v *Viper, c *Config) *zap.Logger {
	var logger = log.NewLogger(c.Log)
	v.SetLogger(logger)
	c.AddObserver("log", func(c *Config) {

	})
	return logger
}
