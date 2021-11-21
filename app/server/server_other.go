//go:build !windows
// +build !windows

package server

import (
	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
	"time"
)

// 初始化服务
func InitServer(addr string, e *gin.Engine) Server {
	s := endless.NewServer(addr, e)
	s.ReadHeaderTimeout = 10 * time.Millisecond
	s.WriteTimeout = 10 * time.Second
	s.MaxHeaderBytes = 1 << 20
	return s
}
