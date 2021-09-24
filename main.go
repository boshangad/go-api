package main

import (
	"github.com/boshangad/go-api/core"
	"github.com/boshangad/go-api/global"
)

func main() {
	// 初始化Viper
	global.G_UserVp = core.Viper()
	// 初始化zap日志库
	global.G_LOG = core.Zap()
	// 初始化数据库连接
	global.G_DB = core.InitDb()
	// 执行服务
	core.Run()
}