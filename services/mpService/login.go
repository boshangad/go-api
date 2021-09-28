package mpService

import (
	"context"
	"errors"
	"github.com/boshangad/go-api/ent"
	"github.com/boshangad/go-api/ent/appuser"
	"github.com/boshangad/go-api/global"
	"github.com/boshangad/go-api/global/db"
	"time"
)

type Login struct {
	Code string
}

func (that *Login) Login(appModel *ent.App) (*ent.AppUser, error) {
	mp := DefaultWechat.MiniProgram(appModel)
	codeSession, err := mp.GetAuth().Code2Session(that.Code)
	if err != nil {
		global.G_LOG.Info("登录code无效:" + err.Error())
		return nil, errors.New("login error, code invalid")
	}
	ctx := context.Background()
	au := db.DefaultClient().AppUser
	appUserModel, err := au.Query().Where(appuser.And(
		appuser.AppIDEQ(appModel.ID),
		appuser.OpenIDEQ(codeSession.OpenID),
	)).First(ctx)
	if err != nil {
		// 表示查询失败
		if !ent.IsNotFound(err) {
			global.G_LOG.Info("数据查询错误：" + err.Error())
			return nil, err
		}
		appUserModel, err = au.Create().
			SetAppID(appModel.ID).
			SetOpenID(codeSession.OpenID).
			Save(ctx)
		if err != nil {
			global.G_LOG.Info("新增appUser表失败：" + err.Error())
			return nil, err
		}
	}
	appUserModel, err = appUserModel.Update().
		SetSessionKey(codeSession.SessionKey).
		SetUnionid(codeSession.UnionID).
		SetLastLoginTime(uint64(time.Now().Unix())).
		Save(ctx)
	if err != nil {
		global.G_LOG.Info("更新appUser表失败：" + err.Error())
		return nil, err
	}
	return appUserModel, nil
}