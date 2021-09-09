package config

import (
	"github.com/boshangad/go-api/core/db/connection"
)

// Db连接配置
type dbConfig struct {
	Default string `json:"default,omitempty"`
	Connections map[string]*connection.Params `json:"connections"`
}

// Init 初始化数据
func (that *dbConfig) Init() {
	if that.Default == "" {
		that.Default = "db"
	}
	if that.Connections == nil {
		that.Connections = make(map[string]*connection.Params)
	}
	that.ConnectionAllClient()
}

// ConnectionAllClient 关闭全部的db客户端
func (that *dbConfig) ConnectionAllClient() {
	if len(that.Connections) > 0 {
		for _, params := range that.Connections {
			client := connection.Connect(*params).Open()
			params.Client = client
		}
	}
}

// CloseAllClient 关闭全部的db客户端
func (that *dbConfig) CloseAllClient() {
	if len(that.Connections) > 0 {
		for _, params := range that.Connections {
			_ = params.Client.Close()
			params.Client = nil
		}
	}
}
