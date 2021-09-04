package services

import (
	"context"
	"crypto/md5"
	"github.com/boshangad/go-api/core/db"
	"github.com/boshangad/go-api/ent"
	"github.com/boshangad/go-api/ent/appusertoken"
	uuid2 "github.com/google/uuid"
	"github.com/o1egl/paseto"
	"log"
)

type tokenServiceInterface interface {
	EncryptString(string) string
	DecryToken(string) string
	GetModelByToken(string) *ent.AppUserToken
}

type tokenService struct {
	tokenServiceInterface
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