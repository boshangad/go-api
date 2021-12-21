package global

import (
	"github.com/boshangad/v1/app/cache"
	"github.com/boshangad/v1/app/config"
	"github.com/boshangad/v1/app/db"
	"github.com/boshangad/v1/app/sms"

	"github.com/boshangad/v1/app/log"
	"go.uber.org/zap"
	"golang.org/x/sync/singleflight"
)

var (
	Viper *config.Viper = config.NewViper("config")
	// Config 配置数据
	Config *config.Config = config.NewConfig(Viper)
	// Log 日志配置
	Log *zap.Logger = log.NewLogger(Config.Log)
	// Db 数据库
	Db *db.Db = db.NewDb(db.OpenDbByConfig(Config.Db), Log)
	// Cache 缓存服务
	Cache cache.CacheInterface = cache.NewMemory(Config.Cache)
	// 内存缓存
	Memoey cache.CacheInterface = cache.NewMemory(Config.Cache)
	// ConcurrencyControl 防止缓存击穿，并发控制
	ConcurrencyControl = &singleflight.Group{}
	// Sms 短信推送网关
	Sms = sms.NewSmsGateway(Config.Sms)
)
