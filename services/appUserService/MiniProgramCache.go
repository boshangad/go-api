package appUserService

import (
	"time"

	"github.com/boshangad/v1/app/cache"
	"github.com/boshangad/v1/global"
)

// 默认的小程序缓存
var defaultMiniProgramCache = NewMiniProgramCache(global.Memoey)

// 小程序缓存
type MiniProgramCache struct {
	cache cache.CacheInterface
}

// 获取
func (that MiniProgramCache) Get(key string) interface{} {
	return that.cache.Get(key)
}

// 设置
func (that *MiniProgramCache) Set(key string, d interface{}, ttl time.Duration) (err error) {
	that.cache.Set(key, d, ttl)
	return
}

// 是否存在
func (that *MiniProgramCache) IsExist(key string) (ok bool) {
	return that.cache.Exists(key)
}

// 删除
func (that *MiniProgramCache) Delete(key string) error {
	that.cache.Delete(key)
	return nil
}

// 实例化小程序缓存
func NewMiniProgramCache(cache cache.CacheInterface) *MiniProgramCache {
	return &MiniProgramCache{
		cache: cache,
	}
}
