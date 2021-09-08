package config

type emailPush struct {
	// 默认使用的配置项
	Default string `json:"default"`
	// 验证码最大验证次数
	CodeMaxCheckNumber int64 `json:"code_max_check_number,omitempty"`
	// 网关
	Gateways map[string]emailPushGateway `json:"gateways,omitempty"`
}

type emailPushGateway struct {
	Gateway string `json:"gateway"`
	RegionId string `json:"region_id,omitempty"`
	AccessKeyId string `json:"access_key_id,omitempty"`
	AccessKeySecret string `json:"access_key_secret,omitempty"`
}