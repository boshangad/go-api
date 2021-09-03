package appUserTokenService

import (
	"encoding/hex"
	"fmt"
	"github.com/dchest/captcha"
	"github.com/o1egl/paseto"
	"github.com/tend/wechatServer/core/config"
	"golang.org/x/crypto/ed25519"
	"log"
	"time"
)

// EnTokenByPaseto 加密
func EnTokenByPaseto(data string) string {
	var (
		token string
		err error
	)
	conf := config.Get().Paseto
	now := time.Now()
	jsonToken := paseto.JSONToken{
		IssuedAt: now,
	}
	// Add custom claim to the token
	jsonToken.Set("data", data)
	footer := captcha.RandomDigits(8)
	if conf.Version == "v1" {
		v1 := paseto.NewV1()
		if conf.Used == "local" {
			token, err = v1.Encrypt([]byte(conf.PrivateKey), jsonToken, footer)
		} else {
			b, _ := hex.DecodeString(conf.PrivateKey)
			privateKey := ed25519.PrivateKey(b)
			token, err = v1.Sign(privateKey, jsonToken, footer)
		}
	} else {
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
		log.Println(fmt.Sprintf("paseto encrypt%s-%s err:$%s", conf.Version, conf.Used, err))
		return ""
	}
	return token
}

// DeTokenByPaseto 解密
func DeTokenByPaseto(token string) paseto.JSONToken {
	var (
		err error
		newJsonToken paseto.JSONToken
		newFooter string
	)
	conf := config.Get().Paseto
	if conf.Version == "v1" {
		v1 := paseto.NewV1()
		if conf.Used == "local" {
			err = v1.Decrypt(token, []byte(conf.PrivateKey), &newJsonToken, &newFooter)
		} else {
			b, _ := hex.DecodeString(conf.PrivateKey)
			publicKey := ed25519.PublicKey(b)
			err = v1.Verify(token, publicKey, &newJsonToken, &newFooter)
		}
	} else {
		v2 := paseto.NewV2()
		if conf.Used == "local" {
			err = v2.Decrypt(token, []byte(conf.PrivateKey), &newJsonToken, &newFooter)
		} else {
			b, _ := hex.DecodeString(conf.PublicKey)
			publicKey := ed25519.PublicKey(b)
			err = v2.Verify(token, publicKey, &newJsonToken, &newFooter)
		}
	}
	if err != nil {
		log.Panic(fmt.Sprintf("pasetoencrypt%s-%s err:$%s", conf.Version, conf.Used, err))
		return newJsonToken
	}
	return newJsonToken
}