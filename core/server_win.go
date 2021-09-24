//go:build windows
// +build windows

package core

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

// 初始化服务
func initServer(addr string, e *gin.Engine) server {
	return &http.Server{
		Addr: addr,
		Handler: e,
		ReadTimeout: 10 * time.Second,
		WriteTimeout: 10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
}