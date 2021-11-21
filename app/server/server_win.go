//go:build windows
// +build windows

package server

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// 初始化服务
func InitServer(addr string, e *gin.Engine) Server {
	return &http.Server{
		Addr:           addr,
		Handler:        e,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
}
