package app

import (
	"context"
	"net/http"
	"os/signal"
	"syscall"

	"github.com/boshangad/v1/app/global"
	"github.com/boshangad/v1/app/server"
	"go.uber.org/zap"
)

func Run() {
	global.Log.Info("Server start."+global.Config.App.Listen)
	// Create context that listens for the interrupt signal from the OS.
	reloadCtx, serverStop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer serverStop()
	// 优雅的重启服务
	serverCore := server.NewServer(global.Config.App.Listen)
	go func() {
		// 服务连接
		if err := server.Run(serverCore); err != nil && err != http.ErrServerClosed {
			global.Log.Fatal("server listen error: %s\n", zap.Error(err))
		}
	}()
	// Listen for the interrupt signal.
	<-reloadCtx.Done()
	// Restore default behavior on the interrupt signal and notify user of shutdown.
	serverStop()
	if err := server.Close(serverCore); err != nil {
		global.Log.Error("server close error", zap.Error(err))
		return
	}
	global.Log.Info("Server exiting")
}
