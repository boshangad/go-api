package app

import (
	"context"
	"net/http"
	"os/signal"
	"strings"
	"syscall"

	"github.com/boshangad/v1/app/config"
	"github.com/boshangad/v1/app/log"
	"github.com/boshangad/v1/app/server"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func Run(conf *config.Config, routerFuns ...func(*gin.Engine)) {
	// 初始化端口
	var (
		listen  = conf.App.Listen
		listens = make([]string, 2)
	)
	if listen != "" && listen[:1] != ":" {
		listens = strings.SplitN(listen, ":", 2)
		if len(listens) == 1 {
			listen = ":" + listen
		}
	}
	log.Logger.Info("Server start." + listen)
	// Create context that listens for the interrupt signal from the OS.
	reloadCtx, serverStop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer serverStop()
	// 优雅的重启服务
	serverCore := server.NewServer(listen, func(e *gin.Engine) {
		if len(routerFuns) > 0 {
			for _, f := range routerFuns {
				f(e)
			}
		}
	})
	go func() {
		// 服务连接
		if err := server.Run(serverCore); err != nil && err != http.ErrServerClosed {
			log.Logger.Fatal("server listen error: %s\n", zap.Error(err))
		}
	}()
	// Listen for the interrupt signal.
	<-reloadCtx.Done()
	// Restore default behavior on the interrupt signal and notify user of shutdown.
	serverStop()
	if err := server.Close(serverCore); err != nil {
		log.Logger.Error("server close error", zap.Error(err))
		return
	}
	log.Logger.Info("Server exiting")
}
