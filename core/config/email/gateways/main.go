package gateways

import (
	"log"
)

type ConfigInterface interface {
	Send(data Data) (isSuccess bool, err error)
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

// NewGateWay 初始化默认推送网关
func NewGateWay(config ConfigInterface) ConfigInterface {
	switch config.(type) {
	case aliyunConfig:
		return NewAliyunGateway(config.(aliyunConfig))
	default:
		log.Panicln("invalid configuration parameter")
	}
	return nil
}