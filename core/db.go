package core

import (
	"github.com/boshangad/go-api/config"
	"github.com/boshangad/go-api/ent"
	"github.com/boshangad/go-api/global"
)

// InitDb 初始化数据库客户端
func InitDb() (db map[string]*ent.Client) {
	db = map[string]*ent.Client{}
	connections := global.G_CONFIG.DB.Connections
	if connections == nil {
		global.G_CONFIG.DB.Connections = map[string]config.DbParams{}
		connections = global.G_CONFIG.DB.Connections
	}
	for key, dbParams := range connections {
		client, err := ent.Open(dbParams.Driver, dbParams.Dsn())
		if err != nil {
			return nil
		}
		db[key] = client
	}
	return
}