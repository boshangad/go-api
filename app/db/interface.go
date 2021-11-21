package db

import "github.com/boshangad/v1/ent"

type DbClientInterface interface {
	// 打开连接池
	Open() (*ent.Client, error)
	Close() error
}

// 数据库应该要实现的方法
type DbInterace interface {
	DbClientInterface
}
