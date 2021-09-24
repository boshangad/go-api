package core

import (
	"fmt"
	"github.com/boshangad/go-api/global"
	"github.com/boshangad/go-api/initialize"
	"go.uber.org/zap"
	"time"
)

// Run 启动应用
func Run() {
	addr := fmt.Sprintf(":%d", global.G_CONFIG.System.Addr)

	Router := initialize.Routers()

	s := initServer(addr, Router)
	// 保证文本顺序输出
	// In order to ensure that the text order output can be deleted
	time.Sleep(10 * time.Microsecond)
	global.G_LOG.Info("server run success on ", zap.String("address", addr))
	fmt.Printf(`
	欢迎使用
	当前版本:V1.0.0 beta
`)
	global.G_LOG.Error(s.ListenAndServe().Error())
}