package email

import (
	"github.com/boshangad/go-api/core/config/email/gateways"
)

const (
	TypeSystem = 1
	TypeNotify = 2
	TypeLogin    = 3
	TypeRegister = 3
	TypeForget = 4
	TypeSafe   = 5
)

const (
	// StatusDraft 未发布
	StatusDraft = 0
	// StatusUsed 已核销
	StatusUsed = 1
	// StatusPublished 已发布
	StatusPublished = 2
	// StatusExpire 已失效
	StatusExpire = 3
)

type Push struct {
	// 默认使用的配置项
	Default string `json:"default"`
	// 验证码最大验证次数
	CodeMaxCheckNumber int64 `json:"code_max_check_number,omitempty"`
	// 网关
	Gateways map[string]map[string]interface{} `json:"gateways,omitempty"`
	// client
	Clients map[string]gateways.ConfigInterface `json:"-"`
}

// Init 初始化相关数据
func (that *Push) Init() {
	if that.Gateways == nil {
		that.Gateways = make(map[string]map[string]interface{})
	}
	that.Clients = map[string]gateways.ConfigInterface{}
	// 循环初始化邮箱推送客户端
	for key, gatewayConfig := range that.Gateways {
		// 默认配置网关
		if that.Default == "" {
			that.Default = key
		}
		c := gateways.NewGateWay(gatewayConfig)
		that.Clients[key] = c
	}
}

// GetTypeIdWithScope 获取类型通过来源
func GetTypeIdWithScope(scope string) uint64 {
	l := map[string]uint64{
		"system": TypeSystem,
		"notify": TypeNotify,
		"register": TypeRegister,
		"login": TypeLogin,
		"forget": TypeForget,
		"safe": TypeSafe,
	}
	if typeId, ok := l[scope]; ok {
		return typeId
	}
	return 0
}

// StatusList 状态列表
func StatusList() map[int64]string {
	return map[int64]string{
		StatusDraft: "To be sent",
		StatusUsed: "Used",
		StatusPublished: "Has been sent",
		StatusExpire: "Expired",
	}
}