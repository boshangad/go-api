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
	// Add a global hook that runs on all types and all operations.
	client.Use(func(next ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			start := time.Now()
			defer func() {
				log.Printf("Op=%s\tType=%s\tTime=%s\tConcreteType=%T\n", m.Op(), m.Type(), time.Since(start), m)
			}()
			return next.Mutate(ctx, m)
		})
	})
	// 返回实例
	return &Db{
		Client: client,
	}
}

// 实例化库通过配置文件
func OpenDbByConfig(data map[string]interface{}) *ent.Client {
	t, ok := data["driver"].(string)
	if !ok {
		panic("The instantiated DB class must have the `driver` field")
	}
	t = strings.ToLower(strings.TrimSpace(t))
	switch t {
	case dialect.MySQL:
		fallthrough
	default:
		client, err := NewMysql(data).Open()
		if err != nil {
			panic(err)
		}
		return client
	}
}
