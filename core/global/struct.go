package global

import "github.com/boshangad/go-api/ent"

// JsonResponse 输出json格式对象
type JsonResponse struct {
	Error int64 `json:"error"`
	Msg string `json:"msg"`
	Data interface{} `json:"data,omitempty"`
}

// AuthData 鉴权数据结构
type AuthData struct {
	App *ent.App `json:"app,omitempty"`
	AppUser *ent.AppUser `json:"app_user,omitempty"`
	AppUserToken *ent.AppUserToken `json:"app_user_token,omitempty"`
}