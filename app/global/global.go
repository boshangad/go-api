package global

import (
	"github.com/boshangad/v1/app/cache"
	"github.com/boshangad/v1/app/config"

	"github.com/boshangad/v1/app/log"
	"go.uber.org/zap"
)

// 初始化
func init() {
	Db = initDb()
}

var (
	// 配置数据
	Config *config.Config = config.DefaultConfig().Load()
	// 日志配置
	Log *zap.Logger = log.NewLogger(Config.Log)
	// 数据库
	Db *db
	// 缓存服务
	Cache cache.CacheInterface
)
