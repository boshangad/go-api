package cache

type FileCache struct {
	// 缓存键前缀
	KeyPrefix string `mapstructure:"keyPrefix" json:"keyPrefix,omitempty" yaml:"keyPrefix"`
	// 缓存路径
	CachePath string `mapstructure:"cachePath" json:"cachePath,omitempty" yaml:"cachePath"`
	// 缓存文件前缀
	CacheFileSuffix string `mapstructure:"cacheFileSuffix" json:"cacheFileSuffix,omitempty" yaml:"cacheFileSuffix"`
	// 目录层级
	DirectoryLevel int `mapstructure:"directoryLevel" json:"directoryLevel,omitempty" yaml:"directoryLevel"`
	// 缓存回收时间限制
	GcProbability int64 `mapstructure:"gcProbability" json:"gcProbability,omitempty" yaml:"gcProbability"`
	// 文件权限模式
	FileMode int64 `mapstructure:"fileMode" json:"fileMode,omitempty" yaml:"fileMode"`
	// 目录权限
	DirMode uint64 `mapstructure:"dirMode" json:"dirMode,omitempty" yaml:"dirMode"`
	// 以下是私有属性
	cachePath string
}

// 缓存路径
func (that *FileCache) getCachePath() string {
	if that.cachePath == "" {
		that.cachePath = that.CachePath
	}
	return that.cachePath
}

// 缓存文件名称
func (that *FileCache) getCacheFileKey(key string) string {
	return that.KeyPrefix + buildingKey(key)
}

// 获取缓存数据
func (that *FileCache) Get(key string) interface{} {
	// os.DirFS()
	return nil
}
