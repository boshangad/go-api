package emailService

import (
	"github.com/boshangad/go-api/core/config"
)

var defaultContainerFuncs = map[string]func(string) EmailInterface {
	"aliyun": func(s string) EmailInterface {return NewAliyun(s)},
}

type EmailInterface interface {
	// Send 发送邮件
	Send(emailStruct EmailStruct) (isSuccess bool, err error)
	// CheckCode 检查发出的代码
	CheckCode(email, code string, typeId uint64, appId uint64) (err error)
}

type EmailStruct struct {
	AppId       uint64 `json:"app_id,omitempty"`
	Email       string `json:"email,omitempty"`
	Scope       string `json:"scope,omitempty"`
	TypeId      uint64 `json:"type_id,omitempty"`
	GateWay     string `json:"gate_way,omitempty"`
	Ip          string `json:"ip,omitempty"`
	FromName    string `json:"from_name,omitempty"`
	FromAddress string `json:"from_address,omitempty"`
	Title       string `json:"title,omitempty"`
	Content     string `json:"content,omitempty"`
	// 核销码
	Data        map[string]string `json:"data,omitempty"`
	ReturnMsg   string `json:"return_msg,omitempty"`
}

// NewDefaultGateWay 初始化默认推送网关
func NewDefaultGateWay(key string) EmailInterface {
	emailPush := config.Get().EmailPush
	if emailPush == nil {
		panic("No mail service configuration was found#1")
	}
	if emailPush.Gateways == nil || len(emailPush.Gateways) < 1 {
		panic("No mail service configuration was found#2")
	}
	if key == "" {
		key = emailPush.Default
	}
	pushConfig, ok := emailPush.Gateways[key]
	if !ok {
		panic("")
	}
	f, ok := defaultContainerFuncs[pushConfig.Gateway]
	if !ok {
		panic("")
	}
	return f(key)
}