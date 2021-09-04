package app

// 定义返回登录成功的数据结构
type Login struct {
	AccessToken string `json:"access_token,omitempty"`
	ExpireTime int64 `json:"expire_time,omitempty"`
	IsLoadUserProfile bool `json:"is_load_user_profile,omitempty"`
	IsBindUser bool `json:"is_bind_user,omitempty"`
	UserAlias string `json:"user_alias,omitempty"`
}

type returnApp struct {
	Alias string `json:"alias,omitempty"`
	Title string `json:"title,omitempty"`
}

type returnAppUser struct {
	Id uint64 `json:"id,omitempty"`
	IsLoadUserProfile bool `json:"is_load_user_profile,omitempty"`
	Nickname string `json:"nickname,omitempty"`
	Avatar string `json:"avatar,omitempty"`
	AvatarURL string `json:"avatar_url,omitempty"`
	Gender int `json:"gender,omitempty"`
	County string `json:"county,omitempty"`
	CountryCode string `json:"country_code,omitempty"`
	Province string `json:"province,omitempty"`
	City string `json:"city,omitempty"`
	Language string `json:"language,omitempty"`
	PhoneNumber string `json:"phone_number,omitempty"`
	PurePhoneNumber string `json:"pure_phone_number,omitempty"`
	LoadUserProfileTime uint64 `json:"load_user_profile_time,omitempty"`
	LastLoginTime int64     `json:"last_login_time,omitempty"`
	App           returnApp `json:"app"`
}
