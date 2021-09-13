package email

import "github.com/boshangad/go-api/core/config/email/gateways"

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
	Gateways map[string]gateways.ConfigInterface `json:"gateways,omitempty"`
}
