package userService

import (
	"context"
	"fmt"

	"github.com/boshangad/v1/app/global"
	"github.com/boshangad/v1/app/helpers"
	"github.com/boshangad/v1/ent"
	euser "github.com/boshangad/v1/ent/user"
	"go.uber.org/zap"
)

type Register struct {
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
	// 区号
	DialCode string `json:"dial_code,omitempty" binding:""`
	Mobile   string `json:"mobile,omitempty" binding:""`
	SmsCode  string `json:"sms_code,omitempty" binding:"required_with=Mobile,omitempty,min=4,max=8"`
	// 邮箱
	Email     string `json:"email,omitempty" binding:"required_if=Type email,omitempty,email"`
	EmailCode string `json:"code,omitempty" binding:"required_with=Email,omitempty,min=4,max=8"`
	// 验证码
	Captcha string `json:"captcha,omitempty" binding:"omitempty,alphanum,min=4,max=8"`
}

// 注册
func (that Register) Register() (user *ent.User, err error) {
	var (
		query                 = global.Db.User.Query()
		ctx   context.Context = context.Background()
	)
	// 判断用户名是否被占用
	if that.Username != "" {
		_, err = query.Clone().
			Where(euser.UsernameEQ(that.Username)).
			First(ctx)
		if err == nil {
			return nil, fmt.Errorf("user %s exists", that.Username)
		} else if !ent.IsNotFound(err) {
			global.Log.Warn("query user failed", zap.Error(err))
			return nil, fmt.Errorf("db error")
		}
	}
	// 判断手机号是否被占用
	if that.Mobile != "" {
		_, err = query.Clone().
			Where(euser.DialCodeEQ(that.DialCode), euser.MobileEQ(that.Mobile)).
			First(ctx)
		if err == nil {
			return nil, fmt.Errorf("user %s exists", that.Username)
		} else if !ent.IsNotFound(err) {
			global.Log.Warn("query user failed", zap.Error(err))
			return nil, fmt.Errorf("db error")
		}
	}
	// 判断邮箱是否被占用
	if that.Email != "" {
		_, err = query.Clone().
			Where(euser.EmailEQ(that.Mobile)).
			First(ctx)
		if err == nil {
			return nil, fmt.Errorf("user %s exists", that.Username)
		} else if !ent.IsNotFound(err) {
			global.Log.Warn("query user failed", zap.Error(err))
			return nil, fmt.Errorf("db error")
		}
	}
	passwordHash, err := helpers.PasswordHash(that.Password, 0)
	if err != nil {
		return nil, fmt.Errorf("password is failed")
	}
	// 检测占用情况通过
	user, err = global.Db.User.Create().
		SetUsername(that.Username).
		SetMobile(that.Mobile).
		SetEmail(that.Email).
		SetPassword(passwordHash).
		Save(ctx)
	return
}
