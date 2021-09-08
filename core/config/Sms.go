package config

type sms struct {
	// 默认使用的配置项
	Default string `json:"default"`
	// 验证码最大验证次数
	CodeMaxCheckNumber int64 `json:"code_max_check_number,omitempty"`
	// 配置项
	Gateways map[string]smsGateway `json:"gateways,omitempty"`
}

type smsGateway struct {
	Gateway string `json:"gateway"`
	RegionId string `json:"region_id,omitempty"`
	AccessKeyId string `json:"access_key_id,omitempty"`
	AccessKeySecret string `json:"access_key_secret,omitempty"`
}