package userService

import (
	"context"
	"errors"
	"github.com/boshangad/go-api/ent"
	"github.com/boshangad/go-api/ent/user"
	"github.com/boshangad/go-api/global/db"
	"github.com/boshangad/go-api/services/smsService"
	"github.com/boshangad/go-api/utils"
)

// LoginByUsername 用户登录使用用户名登录
func LoginByUsername(username, password string) (*ent.User, error) {
	if !CheckAllowUsernameLogin(username) {
		return nil, errors.New("login with username field is forbidden, please use other methods")
	}
	ctx := context.Background()
	userModel, err := db.DefaultClient().User.Query().
		Where(user.And(user.UsernameEQ(username), user.StatusEQ(StatusEnable))).
		First(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, errors.New("username or password is invalid")
		}
		return nil, err
	}
	if !utils.PasswordVerify(password, userModel.Password) {
		return nil, errors.New("username or password is invalid")
	}
	if userModel.Status != StatusEnable {
		return nil, errors.New("the user is not activated, please activate and try again")
	}
	return userModel, nil
}

// LoginByMobileWithPassword 用户登录使用手机号和密码登录
func LoginByMobileWithPassword(dialCode, mobile, password string) (*ent.User, error) {
	if !CheckAllowMobileLogin(dialCode, mobile) {
		return nil, errors.New("login with mobile field is forbidden, please use other methods")
	}
	ctx := context.Background()
	userModel, err := db.DefaultClient().User.Query().
		Where(user.And(user.Mobile(mobile), user.DialCodeEQ(dialCode), user.DeleteTimeEQ(0))).
		First(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, errors.New("username or password is invalid")
		}
		return nil, err
	}
	if !utils.PasswordVerify(password, userModel.Password) {
		return nil, errors.New("username or password is invalid")
	}
	if userModel.Status != StatusEnable {
		return nil, errors.New("the user is not activated, please activate and try again")
	}
	return userModel, nil
}

// LoginByMobileWithCode 用户登录使用手机号和短信验证码登录
func LoginByMobileWithCode(dialCode, mobile, code string) (*ent.User, error) {
	if !CheckAllowMobileLogin(dialCode, mobile) {
		return nil, errors.New("login with mobile field is forbidden, please use other methods")
	}
	ctx := context.Background()
	// 检查是否存在登录的验证码发送日志记录
	err := smsService.CheckCodeIsValid(dialCode, mobile, code, smsService.SmsTypeLogin)
	if err != nil {
		return nil, err
	}
	// 通过手机号识别用户
	userModel, err := db.DefaultClient().User.Query().
		Where(user.And(user.Mobile(mobile), user.DialCodeEQ(dialCode), user.DeleteTimeEQ(0))).
		First(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, errors.New("username or password is invalid")
		}
		return nil, err
	}
	// 检查用户是否激活
	if userModel.Status != StatusEnable {
		return nil, errors.New("the user is not activated, please activate and try again")
	}
	return userModel, nil
}

// LoginByMobileWithPasswordAndCode 用户登录使用手机号、短信验证码和密码登录
func LoginByMobileWithPasswordAndCode(dialCode, mobile, password, code string)  {

}

// LoginByEmailWithPassword 用户登录使用邮箱和密码登录
func LoginByEmailWithPassword(email, password string)  {

}

// LoginByEmailWithCode 用户登录使用邮箱和邮箱验证码登录
func LoginByEmailWithCode(email, code string)  {

}

// LoginByEmailWithPasswordAndCode 用户登录使用邮箱、密码和邮箱验证码登录
func LoginByEmailWithPasswordAndCode(email, password, code string)  {

}