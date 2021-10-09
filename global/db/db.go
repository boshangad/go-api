package db

import (
	"context"
	"github.com/boshangad/go-api/ent"
	"github.com/boshangad/go-api/global"
	"github.com/pkg/errors"
	"strings"
)

func DefaultClient() *ent.Client {
	key := strings.TrimSpace(global.G_CONFIG.DB.Default)
	if key == "" {
		global.G_LOG.Panic("")
	}
	return Client(key)
}

func Client(key string) *ent.Client {
	db, ok := global.G_DB[key]
	if !ok {
		global.G_LOG.Fatal("ddddd")
	}
	return db
}

func WithTx(ctx context.Context, client *ent.Client, fn func(tx *ent.Tx) error) error {
	if ctx == nil {
		ctx = context.Background()
	}
	if client == nil {
		client = DefaultClient()
	}
	tx, err := client.Tx(ctx)
	if err != nil {
		return err
	}
	defer func() {
		if v := recover(); v != nil {
			err = tx.Rollback()
			global.G_LOG.Error(err.Error())
			panic(v)
		}
	}()
	if err := fn(tx); err != nil {
		if rerr := tx.Rollback(); rerr != nil {
			err = errors.Wrapf(err, "rolling back transaction: %v", rerr)
		}
		return err
	}
	if err := tx.Commit(); err != nil {
		return errors.Wrapf(err, "committing transaction: %v", err)
	}
	return nil
}