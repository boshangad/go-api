package cache

import (
	"time"

	"github.com/patrickmn/go-cache"
)

// 缓存数据结构
type Memory struct {
	*cache.Cache
	// 缓存文件
	Filename string `json:"filename,omitempty" yaml:"filename"`
	// 默认失效时间
	DefaultExpiration time.Duration `json:"defaultExpiration,omitempty" yaml:"defaultExpiration"`
	// 清除任务执行时间
	CleanupInterval time.Duration `json:"cleanupInterval,omitempty" yaml:"cleanupInterval"`
}

// 获取缓存
func (that *Memory) Get(key string) interface{} {
	v, ok := that.Cache.Get(key)
	if ok {
		return v
	}
	return nil
}

// 缓存是否存在
func (that *Memory) Exists(key string) (ok bool) {
	_, ok = that.Cache.Get(key)
	return
}

// 实例化文件缓存
func NewMemory(config map[string]interface{}) *Memory {
	return &Memory{
		Cache: cache.New(120*time.Second, 10*time.Minute),
	}
}
