package global

import (
	"errors"
	"strings"

	"entgo.io/ent/dialect"
	dbbase "github.com/boshangad/v1/app/db"
	"github.com/boshangad/v1/ent"
	"github.com/mitchellh/mapstructure"
	"go.uber.org/zap"
)

// 数据库结构
type db struct {
	DefaultDbname string
	// 客户端
	clients map[string]*ent.Client
}

// 获取客户端
func (that db) Get(name string) (client *ent.Client) {
	var ok bool
	if name == "" {
		if that.clients == nil {
			return
		}
		name = that.DefaultDbname
		if name == "" {
			return
		}
	}
	client, ok = that.clients[name]
	if !ok {
		panic("database [" + name + "] not found.")
	}
	return client
}

// 设置数据库客户端
func (that *db) Set(key string, client *ent.Client) {
	that.clients[key] = client
}

// 实例化库通过配置文件
func (that *db) NewDbByConfig(data map[string]interface{}) (*ent.Client, error) {
	t, ok := data["driver"].(string)
	if !ok {
		return nil, errors.New("The parameter must have a field `driver`")
	}
	t = strings.ToLower(strings.TrimSpace(t))
	switch t {
	case dialect.MySQL:
		fallthrough
	default:
		mysql := dbbase.Mysql{}
		err := mapstructure.Decode(data, &mysql)
		mysql.Driver = dialect.MySQL
		if err != nil {
			return nil, err
		}
		return mysql.Open()
	}
}

// 初始化数据库连接
func initDb() *db {
	var database = &db{DefaultDbname: "db", clients: make(map[string]*ent.Client)}
	if Config.Db != nil {
		for dbName, dbConfig := range Config.Db {
			dbName = strings.TrimSpace(dbName)
			dbClient, err := database.NewDbByConfig(dbConfig)
			if err != nil {
				Log.Warn("db config fail", zap.String("name", dbName), zap.Error(err))
				return database
			}
			database.Set(dbName, dbClient)
		}
	}
	return database
}
