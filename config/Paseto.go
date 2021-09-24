package config

type Paseto struct {
	// 版本
	Version string `json:"version,omitempty"`
	// 使用范围
	Used string `json:"used,omitempty"`
	// 加密私钥
	PrivateKey string `json:"private_key,omitempty"`
	// 加密公钥
	PublicKey string `json:"public_key,omitempty"`
}