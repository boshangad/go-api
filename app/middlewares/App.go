package middlewares

import (
	"github.com/boshangad/v1/app/controller"
	"github.com/boshangad/v1/app/global"
	"github.com/boshangad/v1/ent"
	"github.com/boshangad/v1/services/appService"
	"github.com/google/uuid"
)

// 应用中间件
func AppMiddleware(c *controller.Context) {
	var (
		err       error
		app       *ent.App = c.GetApp()
		appAlias  string
		appStruct = appService.App{}
	)
	if app != nil {
		return
	}
	appAlias = c.GetString("appAlias")
	if appAlias == "" {
		appAlias = c.GetHeader("appAlias")
		if appAlias == "" {
			_ = c.ShouldBind(&appStruct)
		}
	}
	uuid, err := uuid.Parse(appAlias)
	if err != nil {
		c.JsonOut(global.ErrNotice, "application access ID is failed", nil)
		return
	}
	appStruct.UUID = uuid
	app, err = appStruct.GetApp()
	if err != nil {
		c.JsonOut(global.ErrNotice, "invalid application access ID: "+appAlias, nil)
		return
	}
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
