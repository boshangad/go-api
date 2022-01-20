package middlewares

import (
	"context"
	"time"

	"github.com/boshangad/v1/app/controller"
	"github.com/boshangad/v1/app/helpers"
	"github.com/boshangad/v1/ent"
	"github.com/boshangad/v1/global"
	"github.com/google/uuid"
)

// 获取应用缓存键名
func getAppCacheKey(uuid uuid.UUID) string {
	return "@entApp:" + uuid.String()
}

// 缓存应用
func cacheApp() {
	var (
		ctx  = context.Background()
		apps = global.Db.App.Query().AllX(ctx)
	)
	for _, app := range apps {
		global.Memoey.Set(getAppCacheKey(*app.UUID), app, time.Duration(helpers.RandomRange64(1, 30)))
	}
}

// 应用中间件
func AppMiddleware(c *controller.Context) {
	var (
		err      error
		app      *ent.App = c.GetApp()
		appAlias string
	)
	if app != nil {
		return
	}
	appAlias = c.GetString("appAlias")
	if appAlias == "" {
		appAlias = c.GetHeader("appAlias")
		if appAlias == "" {
			appAlias = c.DefaultQuery("appAlias", "")
		}
	}
	uuid, err := uuid.Parse(appAlias)
	if err != nil {
		c.JsonOut(global.ErrNotice, "application access ID is failed", nil)
		return
	}
	if !global.Memoey.Exists(getAppCacheKey(uuid)) {
		// 测验缓存
		_, _, _ = global.ConcurrencyControl.Do("@entCacheInit", func() (v interface{}, err error) {
			cacheApp()
			return
		})
	}
	if !global.Memoey.Exists(getAppCacheKey(uuid)) {
		c.JsonOut(global.ErrNotice, "application access ID is failed", nil)
		return
	}
	app = global.Memoey.Get(getAppCacheKey(uuid)).(*ent.App)
	// 应用已被删除
	if app.DeleteTime > 0 {
		c.JsonOut(global.ErrNotice, "invalid application access ID: "+appAlias, nil)
		return
	}
	// 应用状态异常
	if app.Status != 1 {
		c.JsonOut(global.ErrNotice, "application is not activated, please contact management activation", nil)
		return
	}
	c.SetApp(app)
}
