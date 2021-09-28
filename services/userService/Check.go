package userService

import (
	"context"
	"github.com/boshangad/go-api/ent/user"
	"github.com/boshangad/go-api/global/db"
	"log"
)

func CheckIsExistByMobile(dialCode, mobile string) (exist bool) {
	ctx := context.Background()
	exist, err := db.DefaultClient().User.Query().
		Where(user.And(
			user.MobileEQ(mobile),
			user.DialCodeEQ(dialCode),
		)).
		Exist(ctx)
	if err != nil {
		log.Println("操作失败", err)
		return false
	}
	return
}

// CheckAllowUsernameLogin 检查是否允许用户使用用户名登录
func CheckAllowUsernameLogin(username string) bool {
	return true
}

// CheckAllowMobileLogin 检查用户是否允许手机号进行登录
func CheckAllowMobileLogin(dialCode, mobile string) bool {
	return true
}

// CheckAllowEmailLogin 检查用户是否允许使用邮箱登录
func CheckAllowEmailLogin() bool {
	return true
}

// CheckRequiredPassword 检查用户是否必须使用密码登录
func CheckRequiredPassword() bool {
	return true
}

// CheckRequiredCode 检查用户是否必须使用邮箱验证码或手机验证码登录
func CheckRequiredCode() bool {
	return true
}
