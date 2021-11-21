package controller

import "github.com/gin-gonic/gin"

// 实例化的
type HandlerFunc func(*Context)

// 绑定方法函数
func BindingFunc(fun HandlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		if fun != nil {
			fun(&Context{
				Context: c,
			})
		}
	}
}
