package gateways

import (
	"errors"
	"log"
	"strings"
)

type ConfigInterface interface {
	Name() string
	Send(data Data) (isSuccess bool, err error)
}

type LocalConfig struct {
	Gateway string `json:"gateway,omitempty"`
}

type Data struct {
	AppId       uint64 `json:"app_id,omitempty"`
	Email       string `json:"email,omitempty"`
	Scope       string `json:"scope,omitempty"`
	TypeId      uint64 `json:"type_id,omitempty"`
	Ip          string `json:"ip,omitempty"`
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

func (LocalConfig) Send(data Data) (isSuccess bool, err error) {
	return false, errors.New("local mail delivery")
}

// NewGateWay 初始化默认推送网关
func NewGateWay(config ConfigInterface) ConfigInterface {
	if data, ok := (config).(LocalConfig); ok {
		switch strings.ToLower(data.Gateway) {
		case data.Name():
			return data
		case (AliyunConfig{}).Name():
			if d, ok := (config).(*AliyunConfig); ok {
				return NewAliyunGateway(*d)
			}
		default:
			log.Panicln("invalid configuration parameter, no corresponding gateway found")
		}
	}
	log.Panicln("invalid configuration parameter, gateway required")
	return nil
}