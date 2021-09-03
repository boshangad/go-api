package db

import (
	"context"
	"github.com/pkg/errors"
	"github.com/tend/wechatServer/core/config"
	"github.com/tend/wechatServer/ent"
	_ "github.com/tend/wechatServer/ent/runtime"
)

type dbContainer struct {
	key string
	isConnect bool
	Client *ent.Client
}

// 数据库连接容器
var dbContainers = make(map[string]dbContainer)

// Client 连接器客户端
func Client(key string) *ent.Client {
	connections := config.Get().Db.Connections
	if connections == nil {
		return nil
	}
	container, ok := dbContainers[key]
	// 如果map存在且已经连接过了
	if ok && container.isConnect {
		return container.Client
	}
	container = dbContainer{key: key}
	client := container.connectDb()
	dbContainers[key] = container
	return client
}

// DefaultClient 默认数据库连接器客户端
func DefaultClient() *ent.Client {
	key := config.Get().Db.Default
	client := Client(key)
	return client
}

// WithTx 使用事务
func WithTx(ctx context.Context, client *ent.Client, fn func(tx *ent.Tx) error) error {
	tx, err := client.Tx(ctx)
	if err != nil {
		return err
	}
	defer func() {
		if v := recover(); v != nil {
			err = tx.Rollback()
			if err != nil {
				return
			}
			panic(v)
		}
	}()
	if err = fn(tx); err != nil {
		if rollErr := tx.Rollback(); rollErr != nil {
			err = errors.Wrapf(err, "rolling back transaction: %v", rollErr)
		}
		return err
	}
	if err = tx.Commit(); err != nil {
		return errors.Wrapf(err, "committing transaction: %v", err)
	}
	return nil
}

// CloseClient 关闭连接
func CloseClient(key string) bool {
	container, ok := dbContainers[key]
	// 如果map存在且已经连接过了
	if ok && container.isConnect {
		err := container.Client.Close()
		if err != nil {
			return false
		}
	}
	delete(dbContainers, key)
	return true
}

// CloseAllClient 关闭全部连接
func CloseAllClient() bool {
	for key, _ := range dbContainers {
		CloseClient(key)
	}
	return true
}

// ReloadClient 重载连接
func ReloadClient(key string)  {
	isBool := CloseClient(key)
	if !isBool {
		return
	}
	Client(key)
	return
}

