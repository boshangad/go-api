package appUserTokenService

import (
	"github.com/boshangad/go-api/ent"
)

// CreateTokenWithModel 通过用户登录model创建token
func CreateTokenWithModel(appUserTokenModel *ent.AppUserToken) string {
	return EnTokenByPaseto(appUserTokenModel.UUID)
}
