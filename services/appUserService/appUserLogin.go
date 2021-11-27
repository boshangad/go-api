package appUserService

type AccessToken struct {
	// 登录用户的Token
	AccessToken string `json:"access_token,omitempty"`
}

func (a *AccessToken) String() string {
	return ""
}
