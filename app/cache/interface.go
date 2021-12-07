package cache

import "time"

// 获取缓存接口
type GetCacheInterface interface {
	Get(key string) interface{}
}

// 缓存设置接口
type SetCacheInterface interface {
	Set(key string, value interface{}, ttl time.Duration)
}

// 删除缓存及其相关操作
type DeleteCacheInterface interface {
	Exists(key string) bool
	Delete(key string)
	Flush()
}

// 缓存接口
type CacheInterface interface {
	// 获取器
	GetCacheInterface
	// 设置器
	SetCacheInterface
	// 删除器
	DeleteCacheInterface
}
