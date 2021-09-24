package config

type Jwt struct {
	// jwt签名
	SigningKey string
	// 过期时间
	ExpiresTime int64
	// 缓冲时间
	BufferTime int64
}