package controller

import (
	"github.com/gin-gonic/gin"
)

// 实例化的
type HandlerFunc func(*Context)

// 绑定方法函数
func BindingFunc(fun HandlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 如果传入的func 不是nil值，则调用
		if fun != nil {
			fun(&Context{
				Context: c,
			})
		}
	}
}
