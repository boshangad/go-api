package interfaces

import "net/url"

// 消息实体
type Message interface {
	// 消息类型 text和voice
	GetType() string
	// 正文文案
	GetContent(gateway Gateway) string
	// 数据参数
	GetData(gateway Gateway) url.Values
	// 正文模板
	GetTemplate(gateway Gateway) string
	// 默认使用的网关
	GetGateways() map[string]Gateway
}
