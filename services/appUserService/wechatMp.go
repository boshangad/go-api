package appUserService

import (
	"context"

	"github.com/boshangad/v1/app/controller"
	"github.com/boshangad/v1/app/global"
	"github.com/boshangad/v1/ent"
	"github.com/boshangad/v1/ent/appuser"
	wechat "github.com/silenceper/wechat/v2"
	"github.com/silenceper/wechat/v2/miniprogram/config"
	"go.uber.org/zap"
)

type MiniProgram struct {
	// 登录Code
	Code string `json:"code,omitempty"`
	// 用户AccessToken
	AccessToken string `json:"access_token,omitempty"`
	// Token失效时间
	ExpireTime int64 `json:"expire_time,omitempty"`
}

func (that MiniProgram) Login(c *controller.Context) (err error) {
	var (
		wc      *wechat.Wechat
		ctx     context.Context = context.Background()
		appUser *ent.AppUser
		app     *ent.App = c.GetApp()
	)
	wc = wechat.NewWechat()
	m := wc.GetMiniProgram(&config.Config{
		AppID:     app.AppID,
		AppSecret: app.AppSecret,
	})
	r, err := m.GetAuth().Code2Session(that.Code)
	if err != nil {
		return
	}
	// 成功处理
	appUser, err = global.Db.AppUser.Query().
		Where(
			appuser.AppID(app.ID),
			appuser.OpenID(r.OpenID),
		).
		First(ctx)
	if err != nil {
		// 判断是否是不存在
		if !ent.IsNotFound(err) {
			global.Log.Warn("query appUser failed", zap.String("openId", r.OpenID), zap.Error(err))
			return
		}
		// 创建新的用户
		createAppUser := global.Db.AppUser.Create().
			SetAppID(app.ID).
			SetOpenID(r.OpenID).
			SetUnionid(r.UnionID).
			SetSessionKey(r.SessionKey)
		// 新用户需要检测 UnionID 是否一致
		if r.UnionID != "" {
			oldAppUser, err := global.Db.AppUser.Query().
				Select(appuser.FieldUserID).
				Where(
					appuser.Unionid(r.UnionID),
					appuser.UserIDNEQ(0),
				).
				First(ctx)
			if err == nil {
				// 查找到用户，进行匹配关联
				createAppUser.SetUserID(oldAppUser.UserID)
			} else if !ent.IsNotFound(err) {
				global.Log.Warn("query appUser failed", zap.String("unionid", r.UnionID), zap.Error(err))
				return err
			}
		}
		appUser, err = createAppUser.
			Save(ctx)
		if err != nil {
			global.Log.Warn("create appUser failed", zap.String("openId", r.OpenID), zap.Error(err))
			return
		}
	} else {
		appUser, err = appUser.Update().
			SetUnionid(r.UnionID).
			SetSessionKey(r.SessionKey).
			Save(ctx)
		// 判断用户是否存在
		if err != nil {
			global.Log.Info("update appUser failed", zap.String("openId", r.OpenID), zap.Error(err))
			return
		}
	}
	// 创建用户token 和 其它的数据项

	// 创建用户登录日志
	return err
}
