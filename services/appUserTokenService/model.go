package appUserTokenService

import (
	"context"
	"encoding/hex"
	"fmt"
	"github.com/boshangad/go-api/ent"
	"github.com/boshangad/go-api/global"
	"github.com/boshangad/go-api/global/db"
	"github.com/dchest/captcha"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/o1egl/paseto"
	"golang.org/x/crypto/ed25519"
	"strings"
	"time"
)

type Model struct {
	AppUser *ent.AppUser
	AppUserToken *ent.AppUserToken
}

func (that *Model) LoginByAccessToken(accessToken string) (data string, err error) {
	switch strings.TrimSpace(strings.ToLower(global.G_CONFIG.System.UseTokenType)) {
	case "jwt":
		jwt := NewJWT()
		claim, err := jwt.ParseToken(accessToken)
		if err != nil {
			return "", err
		}
		return claim.UUID, nil
	default:
		var (
			newJsonToken paseto.JSONToken
			newFooter string
		)
		conf := global.G_CONFIG.Paseto
		switch strings.ToLower(strings.TrimSpace(conf.Version)) {
		case "v1":
			v1 := paseto.NewV1()
			if conf.Used == "local" {
				err = v1.Decrypt(accessToken, []byte(conf.PrivateKey), &newJsonToken, &newFooter)
			} else {
				b, _ := hex.DecodeString(conf.PrivateKey)
				publicKey := ed25519.PublicKey(b)
				err = v1.Verify(accessToken, publicKey, &newJsonToken, &newFooter)
			}
		default:
			v2 := paseto.NewV2()
			if conf.Used == "local" {
				err = v2.Decrypt(accessToken, []byte(conf.PrivateKey), &newJsonToken, &newFooter)
			} else {
				b, _ := hex.DecodeString(conf.PublicKey)
				publicKey := ed25519.PublicKey(b)
				err = v2.Verify(accessToken, publicKey, &newJsonToken, &newFooter)
			}
		}
		if err != nil {
			global.G_LOG.Info(fmt.Sprintf("pasetoencrypt%s-%s err:$%s", conf.Version, conf.Used, err))
			return "", nil
		}
		return newJsonToken.Get("uuid"), nil
	}
}

func (that *Model) SetAppUser(appUser *ent.AppUser) *Model {
	that.AppUser = appUser
	return that
}

func (that *Model) SetAppUserToken(appUserToken *ent.AppUserToken) *Model {
	that.AppUserToken = appUserToken
	return that
}

func (that *Model) CreateModel(c *gin.Context) (appUserTokenModel *ent.AppUserToken) {
	var (
		ctx = context.Background()
		err error = nil
	)
	appUserTokenModel, err = db.DefaultClient().AppUserToken.Create().
		SetAppID(that.AppUser.AppID).
		SetAppUserID(that.AppUser.ID).
		SetUserID(that.AppUser.UserID).
		SetClientVersion("").
		SetUUID(uuid.New()).
		SetIP(c.ClientIP()).
		SetExpireTime(uint64(time.Now().Add(time.Hour).Unix())).
		Save(ctx)
	if err != nil {
		global.G_LOG.Error("appUserToken插入失败：" + err.Error())
		return nil
	}
	that.AppUserToken = appUserTokenModel
	return
}

func (that Model) BuildAccessTokenByJwt() string {
	return ""
}

func (that Model) BuildAccessTokenByPaseto() string {
	var (
		token string
		err error
	)
	conf := global.G_CONFIG.Paseto
	now := time.Now()
	jsonToken := paseto.JSONToken{
		IssuedAt: now,
	}
	// Add custom claim to the token
	jsonToken.Set("uuid", that.AppUserToken.UUID.String())
	footer := captcha.RandomDigits(8)

	switch strings.TrimSpace(strings.ToLower(conf.Version)) {
	case "v1":
		v1 := paseto.NewV1()
		if conf.Used == "local" {
			token, err = v1.Encrypt([]byte(conf.PrivateKey), jsonToken, footer)
		} else {
			b, _ := hex.DecodeString(conf.PrivateKey)
			privateKey := ed25519.PrivateKey(b)
			token, err = v1.Sign(privateKey, jsonToken, footer)
		}
	default:
		v2 := paseto.NewV2()
		if conf.Used == "local" {
			token, err = v2.Encrypt([]byte(conf.PrivateKey), jsonToken, footer)
		} else {
			b, _ := hex.DecodeString(conf.PrivateKey)
			privateKey := ed25519.PrivateKey(b)
			token, err = v2.Sign(privateKey, jsonToken, footer)
		}
	}
	if err != nil {
		global.G_LOG.Error(fmt.Sprintf("paseto encrypt%s-%s err:$%s", conf.Version, conf.Used, err))
		return ""
	}
	return token
}

func (that Model) BuildAccessToken() string {
	switch strings.TrimSpace(strings.ToLower(global.G_CONFIG.System.UseTokenType)) {
	case "jwt":
		return that.BuildAccessTokenByJwt()
	default:
		return that.BuildAccessTokenByPaseto()
	}
}

func NewModel() *Model {
	return &Model{}
}