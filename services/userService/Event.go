package userService

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/boshangad/go-api/services/appUserLoginService"
	"github.com/boshangad/go-api/GLOBAL/db"
	"github.com/boshangad/go-api/ent"
	"github.com/boshangad/go-api/ent/appuser"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"log"
	"time"
)

// EventByLoginWithUser 用户登录事件通过用户实例进行
func EventByLoginWithUser(userModel *ent.User, c *gin.Context) error {
	var (
		appModel *ent.App
		appUserModel *ent.AppUser
		appUserToken *ent.AppUserToken
		appUserLogin *ent.AppUserLogin
	)
	ctx := context.Background()
	if val, ok := c.Get("App"); ok && val != nil {
		appModel, _ = val.(*ent.App)
	}
	if appModel == nil {
		return errors.New("AppModel not found")
	}
	client := db.DefaultClient()
	err := db.WithTx(ctx, client, func(tx *ent.Tx) error {
		var err error
		appUserModel, _ = client.AppUser.Query().Where(appuser.And(
			appuser.UserIDEQ(userModel.ID),
			appuser.AppIDEQ(appModel.ID),
		)).First(ctx)
		if appUserModel == nil {
			appUserModel, err = client.AppUser.Create().
				SetAppID(appModel.ID).
				SetUserID(userModel.ID).
				Save(ctx)
			if err != nil {
				log.Println("插入appUser数据失败", err)
				return err
			}
		}
		appUserModel, err = appUserModel.Update().
			SetLastLoginTime(uint64(time.Now().Unix())).
			Save(ctx)
		if err != nil {
			log.Println("保存appUser最后登录时间失败", err)
			return err
		}
		var data = struct {
			Scheme string `json:"scheme"`
			Host string `json:"host"`
			Method string `json:"method"`
			Path string `json:"path"`
			Query string `json:"query"`
			Headers map[string][]string `json:"headers"`
		}{
			c.Request.URL.Scheme,
			c.Request.Host,
			c.Request.Method,
			c.Request.URL.Path,
			c.Request.URL.RawQuery,
			c.Request.Header,
		}
		contentData ,_ := json.Marshal(data)
		appUserLogin, err = client.AppUserLogin.Create().
			SetAppID(appModel.ID).
			SetAppUserID(appUserModel.ID).
			SetUserID(userModel.ID).
			SetLoginTypeID(appUserLoginService.LOGIN_TYPE_USERNAME).
			SetIP(c.ClientIP()).
			SetContent(string(contentData)).
			SetStatus(appUserLoginService.STATUS_SUCCESS).
			Save(ctx)
		if err != nil {
			log.Println("插入appUser登录日志失败", err)
			return err
		}
		appUserToken, err = client.AppUserToken.Create().
			SetAppID(appModel.ID).
			SetAppUserID(appUserModel.ID).
			SetUserID(userModel.ID).
			SetClientVersion(c.Request.UserAgent()).
			SetIP(c.ClientIP()).
			SetUUID(uuid.New()).
			SetExpireTime(uint64(time.Now().Unix() + 3600)).
			Save(ctx)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return err
	}
	c.Set("AppUser", appUserModel)
	c.Set("AppUserLogin", appUserLogin)
	c.Set("AppUserToken", appUserToken)
	c.Set("User", userModel)
	return nil
}