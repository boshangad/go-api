package gateways

import (
	"errors"
	"log"
	"strings"
)

type ConfigInterface interface {
	Name() string
	Send(data Data) (returnMsg string, err error)
}

type LocalConfig struct {
	Gateway string `json:"gateway,omitempty"`
}

type Data struct {
	Email       string `json:"email,omitempty"`
	Title       string `json:"title,omitempty"`
	Content     string `json:"content,omitempty"`
	Data        map[string]interface{} `json:"data,omitempty"`
	// 可选参数
	FromName    string `json:"from_name,omitempty"`
	FromAddress string `json:"from_address,omitempty"`
}

func (LocalConfig) Name() string {
	return "local"
}

func (LocalConfig) Send(data Data) (returnMsg string, err error) {
	return "", errors.New("local mail delivery")
}

// NewGateWay 初始化默认推送网关
func NewGateWay(config map[string]interface{}) ConfigInterface {
	switch strings.ToLower(config["gateway"].(string)) {
	case (LocalConfig{}).Name():
		return LocalConfig{}
	case (AliyunConfig{}).Name():
		return NewAliyunGateway(config)
	default:
		log.Panicln("invalid configuration parameter, no corresponding gateway found")
	}
	return nil
}