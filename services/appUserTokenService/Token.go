package appUserTokenService

import (
	"context"
	"crypto/md5"
	"errors"
	"github.com/boshangad/go-api/ent"
	"github.com/boshangad/go-api/ent/appusertoken"
	"github.com/boshangad/go-api/global/db"
	uuid2 "github.com/google/uuid"
	"github.com/o1egl/paseto"
	"log"
	"time"
)

type Token struct {
	AppUserToken *ent.AppUserToken
	App *ent.App
	AppUser *ent.AppUser
	User *ent.User
}

type tokenServiceInterface interface {
	EncryptString(string) string
	DecryToken(string) string
	GetModelByToken(string) *ent.AppUserToken
}

type tokenService struct {
	EncryptionKey []byte
}

func (t tokenService) EncryptString(key string) string {
	encrypt, err := paseto.NewV1().Encrypt(t.EncryptionKey, key, md5.Sum(t.EncryptionKey))
	if err != nil {
		log.Panic("加密令牌:", key, "，err:", err)
		return ""
	}
	return encrypt
}

func (t tokenService) DecryToken(token string) string {
	var key string
	err := paseto.NewV1().Decrypt(token, t.EncryptionKey, &key, nil)
	if err != nil {
		log.Panic("解码令牌:", token, "，err:", err)
		return ""
	}
	return key
}

func (t tokenService) GetModelByToken(token string) *ent.AppUserToken {
	uuid := t.DecryToken(token)
	if uuid == "" {
		return nil
	}
	ctx := context.Background()
	tokenModel, err := db.DefaultClient().AppUserToken.Query().
		WithApp().WithAppUser().WithUser().
		Where(appusertoken.UUIDEQ(uuid2.Must(uuid2.Parse(uuid)))).
		Order(ent.Desc(appusertoken.FieldID)).
		First(ctx)
	if err != nil {
		log.Panic("通过token获取数据实体：", err)
		return nil
	}
	return tokenModel
}

func NewTokenService(ekey []byte) *tokenService {
	return &tokenService{
		EncryptionKey: ekey,
	}
}

// CheckTokenValid 检查用户令牌是否是有效的
func CheckTokenValid(accessToken string) (*Token, error) {
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
	return &Token{
		AppUserToken: tokenModel,
		App: appModel,
		AppUser: userModel,
	}, nil
}