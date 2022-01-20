package redis

import (
	rredis "github.com/go-redis/redis/v8"
)

func NewRedis(c map[string]interface{}) *rredis.Client {
	return rredis.NewClient(&rredis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
}
