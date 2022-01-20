package captchaService

import (
	"context"

	"github.com/boshangad/v1/global"
	"go.uber.org/zap"
)

type RedisStore struct {
}

func (that RedisStore) Get(id string, clear bool) (digits []byte) {
	var (
		ctx context.Context = context.Background()
		z                   = global.Redis.Get(ctx, id)
	)
	c, err := z.Bytes()
	if err != nil {
		global.Log.Error("redis get failed", zap.Error(err))
		return []byte{}
	}
	return c
}

func (that RedisStore) Set(id string, digits []byte) {
	var ctx context.Context = context.Background()
	global.Redis.Set(ctx, id, digits, 0)
}

func NewRedisStore(config map[string]interface{}) *RedisStore {
	return &RedisStore{}
}
