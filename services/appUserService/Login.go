package appUserService

import (
	"context"
	"fmt"

	"github.com/boshangad/v1/app/global"
	"github.com/boshangad/v1/app/helpers"
	"github.com/boshangad/v1/ent"
	euser "github.com/boshangad/v1/ent/user"
	"go.uber.org/zap"
)

// 登录类型
type LoginType struct {
	// 登录参数
	Type string `json:"type,omitempty"`
}

// 登录验证码
type LoginCaptcha struct {
	// 验证码
	Captcha string `json:"captcha,omitempty"`
}

// 用户名登录
type LoginUsername struct {
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
}

// 手机号登录
type LoginMobile struct {
	DialCode string `json:"dial_code,omitempty"`
	Mobile   string `json:"mobile,omitempty"`
	SmsCode  string `json:"sms_code,omitempty"`
}

// 邮箱登录
type LoginEmail struct {
	Email string `json:"email,omitempty"`
	Code  string `json:"code,omitempty"`
}

// 用户名登录
func (that LoginUsername) Login() (user *ent.User, err error) {
	var (
		ctx context.Context = context.Background()
	)
	query := global.Db.User.Query()
	// 登录查询
	user, err = query.
		Where(euser.UsernameEQ(that.Username), euser.DeleteTimeEQ(0)).
		Order(ent.Desc(euser.FieldID)).
		First(ctx)
	if err != nil {
		if !ent.IsNotFound(err) {
			global.Log.Warn("query user failed", zap.String("username", that.Username), zap.Error(err))
		}
		return nil, fmt.Errorf("username or password incorrect#1")
	}
	// 检测密码是否有效
	if !helpers.PasswordVerify(that.Password, user.Password) {
		return nil, fmt.Errorf("username or password incorrect#2")
	}
	// 检测状态
	if user.Status != 10 {
		return nil, fmt.Errorf("the account has not been activated, please activate first")
	}
	return
}
