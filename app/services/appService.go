package services

import (
	"context"
	"github.com/tend/wechatServer/core/db"
	"github.com/tend/wechatServer/ent"
	entApp "github.com/tend/wechatServer/ent/app"
)

// GetAppModelByAlias 获取应用实体通过应用别名
func GetAppModelByAlias(appAlias string) *ent.App {
	ctx := context.Background()
	appModel, err := db.DefaultClient().App.
		Query().
		Where(entApp.AliasEQ(appAlias)).
		First(ctx)
	if err != nil {
		return nil
	}
	return appModel
}