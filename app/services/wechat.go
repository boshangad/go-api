package services

import (
	"github.com/boshangad/go-api/app/services/cache"
	"github.com/boshangad/go-api/core/db"
	"github.com/boshangad/go-api/ent"
	"github.com/silenceper/wechat/v2"
	"github.com/silenceper/wechat/v2/miniprogram"
	mpConfig "github.com/silenceper/wechat/v2/miniprogram/config"
	"github.com/silenceper/wechat/v2/openplatform"
	opConfig "github.com/silenceper/wechat/v2/openplatform/config"
)

type wc struct {
	wc *wechat.Wechat
	wcItems map[string]interface{}
}
var weData *wc

// 初始化wechat
func init1() {
	w := wechat.NewWechat()
	memory := cache.NewAppTokenCache(
		3,
		db.DefaultClient(),
		nil,
		nil,
		)
	w.SetCache(memory)
	weData = &wc{
		wc: w,
	}
}

// GetMiniProgram 获取微信小程序
func GetMiniProgram(app *ent.App) *miniprogram.MiniProgram {
	re, ok := weData.wcItems[app.AppID]
	if ok {
		return re.(*miniprogram.MiniProgram)
	}
	mp := weData.wc.GetMiniProgram(&mpConfig.Config{
		AppID: app.AppID,
		AppSecret: app.AppSecret,
	})
	return mp
}

// GetOpenPlatform 获取微信开放平台
func GetOpenPlatform(app ent.App) *openplatform.OpenPlatform {
	re, ok := weData.wcItems[app.AppID]
	if ok {
		return re.(*openplatform.OpenPlatform)
	}
	of := weData.wc.GetOpenPlatform(&opConfig.Config{
		AppID: app.AppID,
		AppSecret: app.AppSecret,
	})
	return of
}