package appUserTokenService

import (
	"context"
	"crypto/ed25519"
	"encoding/hex"
	"errors"
	"io/ioutil"
	"time"

	"github.com/boshangad/v1/ent"
	"github.com/boshangad/v1/ent/appusertoken"
	"github.com/boshangad/v1/global"
	"github.com/google/uuid"
	"github.com/vk-rv/pvx"
)

// 令牌
type AccessToken struct {
	// 用户的Token
	Uuid uuid.UUID `json:"uuid,omitempty"`
	// 用户登录的AccessToken
	AccessToken string `json:"access_token,omitempty" validator:"required"`
	// token有效期
	ExpireTime int64 `json:"expire_time,omitempty"`
}

// 通过Token登录用户
func (that AccessToken) Login() (appUserToken *ent.AppUserToken, err error) {
	var (
		publicKey, _, _ = ed25519.GenerateKey(nil)
		pk              = pvx.NewAsymmetricPublicKey(publicKey, pvx.Version4)
		pv4             = pvx.NewPV4Public()
		pvxToken        *pvx.Token
		cliams          pvx.RegisteredClaims = pvx.RegisteredClaims{}
		tokenUuid       uuid.UUID
		// appUser         *ent.AppUser
		ctx     = context.Background()
		nowTime = time.Now()
	)
	pvxToken = pv4.Verify(that.AccessToken, pk)
	if pvxToken.Err() != nil {
		return nil, pvxToken.Err()
	}
	err = pvxToken.ScanClaims(&cliams)
	if err != nil {
		return nil, err
	}
	tokenUuid, err = uuid.Parse(cliams.TokenID)
	if err != nil {
		return nil, err
	}
	// 连接数据库检测
	appUserToken, err = global.Db.AppUserToken.Query().
		WithAppUser(func(auq *ent.AppUserQuery) {
			auq.WithApp().WithUser()
		}).
		Where(appusertoken.UUIDEQ(&tokenUuid)).
		First(ctx)
	if err != nil {
		return nil, err
	}
	// 已经失效的token
	if appUserToken.ExpireTime <= nowTime.Unix() {
		return nil, errors.New("token has expired")
	}
	return
}

// 实例化登录AccessToken
func NewAccessToken() (*AccessToken, error) {
	var (
		nowTime     = time.Now()
		expireTime  = nowTime.Add(2 * time.Hour)
		accessToken = AccessToken{
			Uuid:       uuid.New(),
			ExpireTime: expireTime.Unix(),
		}
		publicKey, privateKey, err                      = ed25519.GenerateKey(nil)
		sk                                              = pvx.NewAsymmetricSecretKey(privateKey, pvx.Version4)
		pv4                                             = pvx.NewPV4Public()
		cliams                     pvx.RegisteredClaims = pvx.RegisteredClaims{
			NotBefore: &nowTime,
			// Expiration: &expireTime,
			TokenID: hex.EncodeToString(accessToken.Uuid[0:]),
			// KeyID:      "",
		}
	)
	ioutil.WriteFile("publicKey.txt", []byte(hex.EncodeToString(publicKey)), 0664)
	ioutil.WriteFile("privateKey.txt", []byte(hex.EncodeToString(privateKey)), 0664)

	token, err := pv4.Sign(sk, &cliams)
	if err != nil {
		return nil, err
	}
	accessToken.AccessToken = token[10:]
	return &accessToken, nil
}
