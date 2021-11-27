package public

import (
	"context"
	"net/http/httputil"

	"github.com/boshangad/v1/app/controller"
	"github.com/boshangad/v1/app/global"
	"github.com/boshangad/v1/ent"
	"github.com/boshangad/v1/ent/appusertoken"
	"github.com/boshangad/v1/services/appUserTokenService"
)

type AccountController struct {
}

// Login 账号密码登录
// @route login [POST]
func (that AccountController) Login(c *controller.Context) {
	var (
		ctx            context.Context = context.Background()
		appUser        *ent.AppUser
		err            error
		accessToken    *appUserTokenService.AccessToken
		httpRequest, _ = httputil.DumpRequest(c.Request, false)
	)
	appUser, err = global.Db.AppUser.Query().First(ctx)
	if err != nil {
		c.JsonOut(global.ErrNotice, err.Error(), nil)
		return
	}
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
			SetLoginTypeID(1).
			SetContent(string(httpRequest)).
			SetStatus(1).
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
