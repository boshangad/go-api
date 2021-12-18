package db

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"entgo.io/ent/dialect"
	"github.com/boshangad/v1/ent"
	_ "github.com/boshangad/v1/ent/runtime"
	"go.uber.org/zap"
)

// db服务
type Db struct {
	// 继承ent的全部属性
	*ent.Client
	logger *zap.Logger
}

// 数据库执行事务
func (that *Db) WithTx(ctx context.Context, fn func(db *ent.Client, tx *ent.Tx) error) error {
	tx, err := that.Client.Tx(ctx)
	if err != nil {
		return err
	}
	// 异常捕获
	defer func() {
		if v := recover(); v != nil {
			rollbackErr := tx.Rollback()
			if rollbackErr != nil {
				err = fmt.Errorf("%s, rolling back transaction: %v", v, rollbackErr)
			} else {
				err = fmt.Errorf("%s", v)
			}
			return
		}
	}()
	// 执行事务方法
	if err := fn(that.Client, tx); err != nil {
		if rerr := tx.Rollback(); rerr != nil {
			err = fmt.Errorf("%s, rolling back transaction: %v", err, rerr)
		}
		return err
	}
	// 提交事务
	if err := tx.Commit(); err != nil {
		return err
	}
	return nil
}

// 实例化DB服务
func NewDb(client *ent.Client, logger *zap.Logger) *Db {
	client = client.Debug()
	// Add a global hook that runs on all types and all operations.
	client.Use(func(next ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			start := time.Now()
			defer func() {
				logger.Info(
					"db usage log",
					zap.String("op", m.Op().String()),
					zap.String("type", m.Type()),
					zap.String("time", time.Since(start).String()),
					zap.String("fields", strings.Join(m.Fields(), ",")),
				)
			}()
			return next.Mutate(ctx, m)
		})
	})
	// 返回实例
	return &Db{
		Client: client,
		logger: logger,
	}
}

// 实例化库通过配置文件
func OpenDbByConfig(data map[string]interface{}) *ent.Client {
	var db DbInterace
	t, ok := data["driver"].(string)
	if !ok {
		t = dialect.MySQL
	}
	t = strings.ToLower(strings.TrimSpace(t))
	switch t {
	case dialect.Postgres:
		db = NewPostgreSql(data)
	case dialect.MySQL:
		db = NewMysql(data)
	default:
		db = NewMysql(data)
	}
	if db == nil {
		log.Panicln("database configuration information conversion failed", data)
	}
	client, err := db.Open()
	if err != nil {
		log.Panicln("open db failed", err)
	}
	return client
}
