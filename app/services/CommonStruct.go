package services

// StructLoginSuccess 定义返回登录成功的数据结构
type StructLoginSuccess struct {
	// 用户登录token
	AccessToken string `json:"access_token,omitempty"`
	// N秒后失效
	ExpireTime int64 `json:"expire_time,omitempty"`
	// 失效的unix时间
	ExpiredTime uint64 `json:"expired_time,omitempty"`
	// 是否已加载用户信息
	IsLoadUserProfile bool `json:"is_load_user_profile,omitempty"`
	// 是否绑定了用户
	IsBindUser bool `json:"is_bind_user,omitempty"`
	// 用户别名
	UserAlias string `json:"user_alias,omitempty"`
}
