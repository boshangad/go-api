package controller

import (
	"github.com/gin-gonic/gin"
)

// 实例化的
type HandlerFunc func(*Context)

// 绑定方法函数
func BindingFunc(fun HandlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 这里就开始查询用户登录等情况
		// 第一步先检测是否存在token等信息
		// appUserMiddleware(c)
		// 第二步，用户没有登录需检测用户所属应用
		// appMiddleware(c)
		// 如果传入的func 不是nil值，则调用
		if fun != nil {
			fun(&Context{
				Context: c,
			})
		}
	}
}

// func appUserMiddleware(c *gin.Context) {
// 	authorization := c.GetHeader("authorization")
// 	if authorization == "" {
// 		if c.Request.Method == http.MethodPost {
// 			authorization = c.DefaultPostForm("access_token", "")
// 		}
// 		if authorization == "" {
// 			authorization = c.DefaultQuery("access_token", "")
// 		}
// 	}
// 	if authorization == "" {
// 		return
// 	}
// 	var (
// 		ctx = context.Background()
// 	)
// 	// 验证登录用户
// 	global.Db.AppUserToken.Query().Where(appusertoken.UUIDEQ(authorization)).First(ctx)

// }

// func appMiddleware(c *gin.Context) {

// }
