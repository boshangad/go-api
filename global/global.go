package global

import (
	"github.com/boshangad/go-api/config"
	"github.com/boshangad/go-api/ent"
	"github.com/casbin/casbin/v2"
	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

var (
	// G_UserVp 用户配置
	G_UserVp *viper.Viper
	// G_DB 数据库客户端
	G_DB map[string]*ent.Client
	// G_REDIS redis服务器连接器
	G_REDIS map[string]*redis.Client
	// G_AsAccess 免登录访问配置器
	G_AsAccess *casbin.CachedEnforcer
	// G_LOG 日志管理器
	G_LOG *zap.Logger
	// G_CONFIG 配置服务器
	G_CONFIG config.Server
	// G_LOG GVA_LOG    *oplogging.Logger
)
