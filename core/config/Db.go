package config

import (
	"entgo.io/ent/dialect"
	"github.com/boshangad/go-api/core/db/connection"
)

// Db连接配置
type dbConfig struct {
	Default string `json:"default,omitempty"`
	Connections map[string]connection.Params `json:"connections"`
}

// 初始化数据
func (that *dbConfig) initDefaultData() {
	if that.Default == "" {
		that.Default = "db"
	}
	if that.Connections == nil {
		that.Connections = map[string]connection.Params{
			that.Default: {
				Driver: dialect.MySQL,
				Host: "127.0.0.1",
				Username: "root",
				Password: "123456",
			},
		}
	}
	// 初始化连接器
	if len(that.Connections) > 0 {
		for _, params := range that.Connections {
			params.Client = connection.Connect(params).Open()
		}
	}
}

