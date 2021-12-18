package config

import "strings"

// 短信网关配置项
type Gateway map[string]interface{}

// 短信配置
type Config struct {
	// 默认使用网关
	Defaults []string `json:"defaults,omitempty" yaml:"defaults"`
	// 网关配置项
	Gateways map[string]Gateway `json:"gateways,omitempty" yaml:"gateways"`
}

// 获取字符串
func (that Gateway) GetString(key string) (s string) {
	if v, ok := that[key]; ok {
		if s, ok1 := v.(string); ok1 {
			return strings.TrimSpace(s)
		}
	}
	return
}

// 获取int64参数值
func (that Gateway) GetInt64(key string) (i int64) {
	if v, ok := that[key]; ok {
		if i, ok1 := v.(int64); ok1 {
			return i
		}
	}
	return
}

// 获取uint64参数值
func (that Gateway) GetUint64(key string) (i uint64) {
	if v, ok := that[key]; ok {
		if i, ok1 := v.(uint64); ok1 {
			return i
		}
	}
	return
}
