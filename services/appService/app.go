package appService

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"github.com/boshangad/v1/ent"
	eapp "github.com/boshangad/v1/ent/app"
	"github.com/boshangad/v1/global"

	"github.com/google/uuid"
	"go.uber.org/zap"
)

// 应用
type App struct {
	// 应用UUID
	UUID uuid.UUID `json:"uuid,omitempty"`
}

// 获取应用App
func (that *App) GetApp() (app *ent.App, err error) {
	var (
		ctx context.Context = context.Background()
	)
	app, err = global.Db.App.Query().
		Modify(func(s *sql.Selector) {
			s.Where(sql.EQ(s.C(eapp.FieldUUID), sql.Raw(string(that.UUID[0:]))))
		}).
		First(ctx)
	if err != nil {
		// 如果是连接数据库错误,则记录相关报错信息
		if !ent.IsNotFound(err) {
			global.Log.Warn("query appUser failed", zap.String("appAlias", that.UUID.String()), zap.Error(err))
		}
		return nil, err
	}
	return app, nil
}
