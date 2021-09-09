package connection

import (
	"context"
	"entgo.io/ent/dialect"
	"github.com/boshangad/go-api/ent"
	_ "github.com/boshangad/go-api/ent/runtime"
	"github.com/pkg/errors"
	"log"
)

type ConnectionInterface interface {
	Open() *ent.Client
}

type Params struct {
	Client *ent.Client `json:"-"`
	Driver string `json:"driver"`
	Database string `json:"database,omitempty"`
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
	Protocol string `json:"protocol,omitempty"`
	Host string `json:"host,omitempty"`
	Port int `json:"port,omitempty"`
	Dbname string `json:"dbname,omitempty"`
	Extras map[string]interface{} `json:"extras,omitempty"`
}

// Connect 连接器
func Connect(data Params) ConnectionInterface {
	switch data.Driver {
	case dialect.SQLite:
	case dialect.Postgres:
	case dialect.Gremlin:
	case dialect.MySQL:
		return NewMysqlConnect(data)
	}
	log.Panicln("Unsupported data connection method " + data.Driver)
	return nil
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


