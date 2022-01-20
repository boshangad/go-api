package global

import (
	"github.com/boshangad/v1/app/cache"
	"github.com/boshangad/v1/app/config"
	"github.com/boshangad/v1/app/db"
	"github.com/boshangad/v1/app/redis"

	"go.uber.org/zap"
	"golang.org/x/sync/singleflight"
)

var (
	Viper *config.Viper = config.NewViper("config")
	// Config 配置数据
	Config *config.Config = config.NewConfig(Viper)
	// Log 日志配置
	Log *zap.Logger = config.NewLog(Viper, Config)
	// Db 数据库
	Db *db.Db = db.NewDb(db.OpenDbByConfig(Config.Db), Log)
	// Cache 缓存服务
	Cache cache.CacheInterface = cache.NewMemory(Config.Cache)
	// 内存缓存
	Memoey cache.CacheInterface = cache.NewMemory(Config.Cache)
	// redis缓存
	Redis = redis.NewRedis(nil)
	// ConcurrencyControl 防止缓存击穿，并发控制
	ConcurrencyControl = &singleflight.Group{}
)
