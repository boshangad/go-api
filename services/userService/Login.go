package userService

import (
	"context"
	"fmt"
	"time"

	"github.com/boshangad/v1/app/global"
	"github.com/boshangad/v1/app/helpers"
	"github.com/boshangad/v1/ent"
	"github.com/boshangad/v1/ent/appuser"
	"github.com/boshangad/v1/ent/appuserloginlog"
	"github.com/boshangad/v1/ent/appusertoken"
	euser "github.com/boshangad/v1/ent/user"
	"github.com/boshangad/v1/services/appUserTokenService"
	"go.uber.org/zap"
)

type Login struct {
	// 登录参数
	Type string `json:"type,omitempty" binding:"required,lowercase,oneof=password code all"`
	// 用户名
	Username string `json:"username,omitempty" binding:"required_if=Type username,omitempty,alphanum,min=5,max=32"`
	// 账户密码
	Password string `json:"password,omitempty" binding:"required_with=Username,omitempty,alphanum,min=8,max=32"`
	// 手机号
	DialCode string `json:"dial_code,omitempty" binding:"required_with=Mobile,omitempty,alphanum,min=2"`
	Mobile   string `json:"mobile,omitempty" binding:"required_if=Type mobile,omitempty"`
	// 邮箱
	Email string `json:"email,omitempty" binding:"required_if=Type email,omitempty,email"`
	Code  string `json:"code,omitempty" binding:"omitempty,min=4,max=8"`
	// 验证码
	Captcha string `json:"captcha,omitempty" binding:"omitempty,alphanum,min=4,max=8"`
}

// 登录
// @param appUserLoginLog *ent.AppUserLoginLog  用户登录日志
func (that Login) Login(appUserLoginLog *ent.AppUserLoginLog) (accessToken *appUserTokenService.AccessToken, err error) {
	var (
		// 进程上下文
		ctx context.Context = context.Background()
		// 更新代理日志
		appUserLoginLogUpdate = appUserLoginLog.Update()
		// 查询对象
		query = global.Db.User.Query().
			Where(euser.DeleteTimeEQ(0)).
			Order(ent.Desc(euser.FieldID))
	)
	// 验证码检测
	if that.Captcha == "" {
		// 检测是否要求有验证码
	} else {
		// 提取验证码逻辑器
	}
	// 用户名不为空
	if that.Mobile != "" {
		appUserLoginLogUpdate.SetLoginTypeID(appuserloginlog.LoginTypeMobile)
		query.Where(euser.MobileEQ(that.Mobile), euser.DialCodeEQ(that.DialCode))
	} else if that.Email != "" {
		appUserLoginLogUpdate.SetLoginTypeID(appuserloginlog.LoginTypeEmail)
		query.Where(euser.MobileEQ(that.Email))
	} else if that.Username != "" {
		appUserLoginLogUpdate.SetLoginTypeID(appuserloginlog.LoginTypeUsername)
		query.Where(euser.UsernameEQ(that.Username))
	} else {
		_, _ = appUserLoginLogUpdate.SetStatus(appuserloginlog.StatusFailed).Save(ctx)
		return nil, fmt.Errorf("")
	}
	user, err := query.First(ctx)
	if err != nil {
		if !ent.IsNotFound(err) {
			global.Log.Warn("query user failed", zap.Error(err))
			appUserLoginLogUpdate.SetStatus(appuserloginlog.StatusServerFailed)
		} else {
			appUserLoginLogUpdate.SetStatus(appuserloginlog.StatusFailed)
		}
		_, _ = appUserLoginLogUpdate.Save(ctx)
		return nil, fmt.Errorf("username or password incorrect#1")
	}
	appUserLoginLogUpdate.SetUser(user)
	// 检测是否要求用户验证码验证
	if that.Captcha == "" {
	}
	// 检测用户的登录索引
	switch that.Type {
	case "password":
		// 检测密码是否有效
		if !helpers.PasswordVerify(that.Password, user.Password) {
			_, _ = appUserLoginLogUpdate.SetStatus(appuserloginlog.StatusFailed).Save(ctx)
			return nil, fmt.Errorf("username or password incorrect#2")
		}
	case "code":
		// 验证码验证,手机号登录要求手机号验证码，邮箱登录要求邮箱验证码
		if !helpers.PasswordVerify(that.Password, user.Password) {
			_, _ = appUserLoginLogUpdate.SetStatus(appuserloginlog.StatusFailed).Save(ctx)
			return nil, fmt.Errorf("username or password incorrect#2")
		}
	case "all":
		fallthrough
	default:
		// 全部都需要
		if !helpers.PasswordVerify(that.Password, user.Password) {
			_, _ = appUserLoginLogUpdate.SetStatus(appuserloginlog.StatusFailed).Save(ctx)
			return nil, fmt.Errorf("username or password incorrect#2")
		}
		if !helpers.PasswordVerify(that.Password, user.Password) {
			_, _ = appUserLoginLogUpdate.SetStatus(appuserloginlog.StatusFailed).Save(ctx)
			return nil, fmt.Errorf("username or password incorrect#2")
		}
	}
	// 检测状态,判断是否需要激活
	if user.Status != euser.StatusActived {
		_, _ = appUserLoginLogUpdate.SetStatus(appuserloginlog.StatusFailed).Save(ctx)
		return nil, fmt.Errorf("the account has not been activated, please activate first")
	}
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
