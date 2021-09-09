package core

import (
	"context"
	"fmt"
	"github.com/boshangad/go-api/core/config"
	_ "github.com/boshangad/go-api/core/db"
	"github.com/boshangad/go-api/core/gin"
	_ "github.com/boshangad/go-api/core/router"
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"
	"os/signal"
	"strconv"
	"time"
)

// Run 启动应用
func Run() {
	c := config.Get()
	addr := ""
	if c.App.Host != "" {
		addr = c.App.Host
	}
	if c.App.Port != 0 {
		addr += ":" + strconv.FormatInt(c.App.Port, 10)
	}
	srv := &http.Server{
		Addr:    addr,
		Handler: gin.Get(),
	}
	go func() {
		// 服务连接
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()
	go func() {
		// 开启pprof，监听请求
		ip := "0.0.0.0:6060"
		if err := http.ListenAndServe(ip, nil); err != nil {
			fmt.Printf("start pprof failed on %s\n", ip)
		}
	}()
	// 等待中断信号以优雅地关闭服务器（设置 5 秒的超时时间）
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutdown Server ...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	log.Println("Server exiting")
}