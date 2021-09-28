package appUserTokenService

import (
	"errors"
	"github.com/boshangad/go-api/global"
	"github.com/dgrijalva/jwt-go"
)

var (
	TokenExpired     = errors.New("token is expired")
	TokenNotValidYet = errors.New("token not active yet")
	TokenMalformed   = errors.New("that's not even a token")
	TokenInvalid     = errors.New("couldn't handle this token: ")
)

type LoginClaim struct {
	UUID string
	jwt.StandardClaims
}

type JWT struct {
	SigningKey []byte
}

// CreateToken 创建一个token
func (j *JWT) CreateToken(claims jwt.Claims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.SigningKey)
}

// CreateTokenByOldToken 旧token 换新token 使用归并回源避免并发问题
func (j *JWT) CreateTokenByOldToken(oldToken string, claims jwt.Claims) (string, error) {
	v, err, _ := global.G_Concurrency_Control.Do("JWT:" + oldToken, func() (interface{}, error) {
		return j.CreateToken(claims)
	})
	return v.(string), err
}

// ParseToken 解析token
func (j *JWT) ParseToken(tokenString string) (*LoginClaim, error) {
	token, err := jwt.ParseWithClaims(tokenString, &LoginClaim{}, func(token *jwt.Token) (i interface{}, e error) {
		return j.SigningKey, nil
	})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, TokenMalformed
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				// Token is expired
				return nil, TokenExpired
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, TokenNotValidYet
			} else {
				return nil, TokenInvalid
			}
		}
	}
	if token != nil {
		if claims, ok := token.Claims.(*LoginClaim); ok && token.Valid {
			return claims, nil
		}
		return nil, TokenInvalid

	}
	return nil, TokenInvalid
}

func NewJWT() *JWT {
	return &JWT{
		[]byte(global.G_CONFIG.Jwt.SigningKey),
	}
}
