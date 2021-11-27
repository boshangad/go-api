package cache

// 获取缓存接口
type GetCacheInterface interface {
	// 获取缓存数据
	Get(key string) interface{}
	// 批量获取缓存数据
	// MultiGet(data []string) map[string]interface{}
}

// 缓存设置接口
type SetCacheInterface interface {
	Set(key string, value interface{}, ttl int64) error
	// MultiSet(data map[string]interface{}, ttl int64) []string
}

// 删除缓存及其相关操作
type DeleteCacheInterface interface {
	Exists(key string) bool
	Delete(key string) error
	Flush() error
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
