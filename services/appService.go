package services

import (
	"context"
	"github.com/boshangad/go-api/ent"
	entApp "github.com/boshangad/go-api/ent/app"
	"github.com/boshangad/go-api/global/db"
	"log"
)

// GetAppModelByAlias 获取应用实体通过应用别名
func GetAppModelByAlias(appAlias string) *ent.App {
	ctx := context.Background()
	appModel, err := db.DefaultClient().App.
		Query().
		Where(entApp.AliasEQ(appAlias)).
		First(ctx)
	if err != nil {
		log.Println("app not found with:" + appAlias)
		return nil
	}
	return appModel
}