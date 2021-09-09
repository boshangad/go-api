package db

import (
	"context"
	"github.com/boshangad/go-api/core/config"
	"github.com/boshangad/go-api/ent"
	"github.com/pkg/errors"
	"log"
)

// Client db客户端
func Client(key string) *ent.Client {
	if params, ok := config.Get().Db.Connections[key]; ok {
		if params.Client == nil {
			log.Panicln("db [" + key + "] connection fail, please try again")
		}
		return params.Client
	}
	log.Panicln("the connected database "+key+" was not found")
	return nil
}

func DefaultClient() *ent.Client {
	return Client(config.Get().Db.Default)
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


