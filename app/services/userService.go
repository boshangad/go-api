package services

import (
	"context"
	"errors"
	"github.com/boshangad/go-api/core/global"
	"time"
)

// CheckTokenValid 检查用户令牌是否是有效的
func CheckTokenValid(accessToken string) (*global.AuthData, error) {
	ts := NewTokenService(nil)

	tokenModel := ts.GetModelByToken(accessToken)
	if tokenModel == nil {
		return nil, errors.New("token invalid")
	}
	// 检查有效期
	if tokenModel.ExpireTime < uint64(time.Now().Unix()) {
		return nil, errors.New("token is expired")
	}
	// 检查IP地址变更
	ctx := context.Background()
	// 拉取用户model
	userModel, err := tokenModel.QueryAppUser().First(ctx)
	if err != nil {
		return nil, err
	}
	appModel, err := tokenModel.QueryApp().First(ctx)
	return &global.AuthData{
		App: appModel,
		AppUser: userModel,
		AppUserToken: tokenModel,
	}, nil
}
