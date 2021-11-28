package public

import (
	"context"
	"net/http/httputil"
	"time"

	"github.com/boshangad/v1/app/controller"
	"github.com/boshangad/v1/app/global"
	"github.com/boshangad/v1/ent"
	"github.com/boshangad/v1/ent/appuser"
	"github.com/boshangad/v1/ent/appuserloginlog"
	"github.com/boshangad/v1/ent/appusertoken"
	"github.com/boshangad/v1/services/appUserService"
	"github.com/boshangad/v1/services/appUserTokenService"
	"go.uber.org/zap"
)

type AccountController struct {
}

// Login 账号密码登录
// @route login [POST]
func (that AccountController) Login(c *controller.Context) {
	var (
		appUserLogin struct {
			appUserService.LoginType
			appUserService.LoginCaptcha
			appUserService.LoginUsername
			appUserService.LoginMobile
			appUserService.LoginEmail
		}
		// 应用
		app = c.GetApp()
		// db上下文
		ctx context.Context = context.Background()
		// 登录用户
		user *ent.User
		// 登录的应用用户
		appUser *ent.AppUser
		// 错误信息
		err error
		// 登录成功的令牌
		accessToken    *appUserTokenService.AccessToken
		httpRequest, _ = httputil.DumpRequest(c.Request, false)
	)
	if err = c.ShouldBind(&appUserLogin); err != nil {
		c.JsonOut(global.ErrNotice, err.Error(), nil)
		return
	}
	switch appUserLogin.Type {
	case "username":
		_, err = appUserLogin.LoginUsername.Login()
		if err != nil {
			c.JsonOut(global.ErrNotice, err.Error(), nil)
			return
		}
	case "mobile":
	case "email":
	default:
		c.JsonOut(global.ErrNotice, "Invalid type, please choose an accurate login method", nil)
		return
	}
	// 查询用户
	global.Db.WithTx(ctx, func(db *ent.Client, tx *ent.Tx) (err error) {
		appUser, err = db.AppUser.Query().
			Where(appuser.UserIDEQ(user.ID), appuser.AppIDEQ(app.ID)).
			First(ctx)
		if err != nil {
			if !ent.IsNotFound(err) {
				global.Log.Warn("query appUser failed to login", zap.Uint64("userId", user.ID), zap.Error(err))
				return err
			}
			appUser, err = db.AppUser.Create().
				SetApp(app).
				SetUser(user).
				SetAvatar(user.Avatar).
				SetGender(user.Sex).
				SetNickname(user.Nickname).
				SetLastLoginTime(time.Now().Unix()).
				SetLoadUserProfileTime(time.Now().Unix()).
				Save(ctx)
			if err != nil {
				if !ent.IsNotFound(err) {
					global.Log.Warn("create appUser failed to login", zap.Uint64("userId", user.ID), zap.Error(err))
				}
				return err
			}
		}
		return
	})
	// 创建登录令牌
	accessToken, err = appUserTokenService.NewAccessToken()
	if err != nil {
		c.JsonOut(global.ErrNotice, err.Error(), nil)
		return
	}
	// 创建token记录
	err = global.Db.WithTx(ctx, func(db *ent.Client, tx *ent.Tx) error {
		// 创建登录令牌
		appUserToken, err := db.AppUserToken.Create().
			SetAppID(appUser.AppID).
			SetAppUserID(appUser.ID).
			SetUserID(appUser.UserID).
			SetUUID(&accessToken.Uuid).
			SetExpireTime(accessToken.ExpireTime).
			SetClientVersion(c.Request.UserAgent()).
			SetIP(c.ClientIP()).
			Save(ctx)
		if err != nil {
			return err
		}
		// 创建登录日志
		_, err = db.AppUserLoginLog.Create().
			SetAppID(appUser.AppID).
			SetAppUserID(appUser.ID).
			SetUserID(appUser.UserID).
			SetIP(c.ClientIP()).
			SetLoginTypeID(appuserloginlog.LoginTypeUnknow).
			SetContent(string(httpRequest)).
			SetStatus(appuserloginlog.StatusSuccess).
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
		return nil
	})
	if err != nil {
		c.JsonOut(global.ErrNotice, err.Error(), nil)
		return
	}
	c.JsonOut(global.ErrSuccess, "success", accessToken)
}

// Register 账号注册
// @route register [POST]
func (that AccountController) Register() {
	// var serviceWithUser = userService.UserRegisterParams{}
	// err := that.Context.ShouldBind(&serviceWithUser)
	// if err != nil {
	// 	that.JsonOutByError(global.ErrNotice, err, nil)
	// 	return
	// }
	// serviceWithUser.Filter()
	// if serviceWithUser.Username == "" && serviceWithUser.Mobile == "" && serviceWithUser.Email == "" {
	// 	that.JsonOut(global.ErrNotice, "用户名不能为空", nil)
	// 	return
	// }
	// userModel, err := serviceWithUser.Register(that.Controller)
	// if err != nil {
	// 	that.JsonOutByError(global.ErrNotice, err, nil)
	// 	return
	// }
	// that.JsonOut(global.ErrSuccess, "success", userModel)
}
