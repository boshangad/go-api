package appUserTokenService

import (
	"context"
	"crypto/ed25519"
	"errors"
	"time"

	"github.com/boshangad/v1/app/global"
	"github.com/boshangad/v1/ent"
	"github.com/boshangad/v1/ent/appusertoken"
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
func (that AccessToken) Login() (appUser *ent.AppUser, err error) {
	var (
		publicKey, _, _ = ed25519.GenerateKey(nil)
		pk              = pvx.NewAsymmetricPublicKey(publicKey, pvx.Version4)
		pv4             = pvx.NewPV4Public()
		pvxToken        *pvx.Token
		cliams          pvx.RegisteredClaims = pvx.RegisteredClaims{}
		tokenUuid       uuid.UUID
		appUserToken    *ent.AppUserToken
		ctx             = context.Background()
		nowTime         = time.Now()
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
	appUser, err = appUserToken.Edges.AppUserOrErr()
	if err != nil {
		return nil, errors.New("token is abnormal, please obtain the token again")
	}
	return
}

// 实例化登录AccessToken
func NewAccessToken() (*AccessToken, error) {
	var (
		nowTime     = time.Now()
		expireTime  = nowTime.Add(2 * time.Hour)
		accessToken = AccessToken{
			Uuid:  uuid.New(),
			ExpireTime: expireTime.Unix(),
		}
		_, privateKey, err = ed25519.GenerateKey(nil)
		sk                 = pvx.NewAsymmetricSecretKey(privateKey, pvx.Version4)
		pv4                = pvx.NewPV4Public()

		cliams pvx.RegisteredClaims = pvx.RegisteredClaims{
			NotBefore:  &nowTime,
			Expiration: &expireTime,
			TokenID:    accessToken.Uuid.String(),
			// KeyID:      "",
		}
	)
	token, err := pv4.Sign(sk, &cliams)
	if err != nil {
		return nil, err
	}
	accessToken.AccessToken = token[10:]
	return &accessToken, nil
}
