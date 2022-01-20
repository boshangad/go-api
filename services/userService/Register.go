package userService

import (
	"context"
	"fmt"
	"time"

	"github.com/boshangad/v1/app/controller"
	"github.com/boshangad/v1/app/helpers"
	"github.com/boshangad/v1/ent"
	"github.com/boshangad/v1/ent/appuser"
	"github.com/boshangad/v1/ent/appuserloginlog"
	"github.com/boshangad/v1/ent/appusertoken"
	euser "github.com/boshangad/v1/ent/user"
	"github.com/boshangad/v1/global"
	"github.com/boshangad/v1/services/appUserLoginLogService"
	"github.com/boshangad/v1/services/appUserTokenService"
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
	// 基础信息
	Birthday string      `json:"birthday,omitempty"`
	Age      helpers.Int `json:"age,omitempty"`
	Nickname string      `json:"nickname,omitempty"`
	Sex      string      `json:"sex,omitempty"`
	// 验证码
	Captcha string `json:"captcha,omitempty" binding:"omitempty,alphanum,min=4,max=8"`
}

// 注册
func (that Register) Register(c *controller.Context) (accessToken *appUserTokenService.AccessToken, err error) {
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
	// 检测占用情况通过
	userCreate := global.Db.User.Create().
		SetUsername(that.Username).
		SetMobile(that.Mobile).
		SetEmail(that.Email)
	if that.Password != "" {
		passwordHash, err := helpers.PasswordHash(that.Password, int(global.Config.App.PasswdLevel))
		if err != nil {
			return nil, fmt.Errorf("password is failed")
		}
		userCreate.SetPassword(passwordHash)
	}
	user, err := userCreate.Save(ctx)
	if err != nil {
		global.Log.Error("user create failed", zap.Error(err))
		return nil, err
	}
	// 注册的同时登录
	// 记录用户登录日志
	appUserLoginLog, err := appUserLoginLogService.CheckAndCreateLoginLogByConfirm(c)
	if err != nil {
		return nil, err
	}
	appUserLoginLogUpdate := appUserLoginLog.Update()
	// 创建登录令牌
	accessToken, err = appUserTokenService.NewAccessToken()
	if err != nil {
		_, _ = appUserLoginLogUpdate.SetStatus(appuserloginlog.StatusServerFailed).Save(ctx)
		return nil, err
	}
	// 创建token记录
	appUserLoginLogUpdate.SetStatus(appuserloginlog.StatusSuccess)
	err = global.Db.WithTx(ctx, func(db *ent.Client, tx *ent.Tx) error {
		// 查找是否存在用户
		appUser, err := global.Db.AppUser.Query().
			Where(appuser.UserIDEQ(user.ID), appuser.AppIDEQ(appUserLoginLog.AppID)).
			First(ctx)
		if err != nil {
			// 未找到记录
			if ent.IsNotFound(err) {
				appUser, err = global.Db.AppUser.Create().
					SetAppID(appUserLoginLog.AppID).
					SetUser(user).
					SetAvatar(user.Avatar).
					SetGender(user.Sex).
					SetNickname(user.Nickname).
					SetLastLoginTime(time.Now().Unix()).
					SetLoadUserProfileTime(time.Now().Unix()).
					Save(ctx)
				// 创建appUser用户失败
				if err != nil {
					if !ent.IsNotFound(err) {
						global.Log.Warn("create appUser failed to login", zap.Uint64("userId", user.ID), zap.Error(err))
					}
				}
			} else {
				global.Log.Warn("query appUser failed to login", zap.Uint64("userId", user.ID), zap.Error(err))
			}
		}
		// 创建登录令牌
		appUserLoginLogUpdate.SetAppUserID(appUser.ID)
		appUserToken, err := db.AppUserToken.Create().
			SetAppID(appUser.AppID).
			SetAppUserID(appUser.ID).
			SetUserID(appUser.UserID).
			SetUUID(&accessToken.Uuid).
			SetExpireTime(accessToken.ExpireTime).
			SetUserAgent(appUserLoginLog.UserAgent).
			SetClientVersion(appUserLoginLog.ClientVersion).
			SetIP(appUserLoginLog.IP).
			Save(ctx)
		if err != nil {
			return err
		}
		// 清除其它的登录令牌
		_, err = db.AppUserToken.Delete().
			Where(appusertoken.AppUserIDEQ(appUser.ID), appusertoken.IDNEQ(appUserToken.ID)).
			Exec(ctx)
		if err != nil {
			return err
		}
		_, _ = appUserLoginLogUpdate.Save(ctx)
		return nil
	})
	if err != nil {
		_, _ = appUserLoginLogUpdate.SetStatus(appuserloginlog.StatusServerFailed).Save(ctx)
		return nil, err
	}
	return accessToken, nil
}
