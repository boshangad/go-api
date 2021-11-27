package public

import (
	"github.com/boshangad/v1/api/public"
	"github.com/boshangad/v1/app/controller"
	"github.com/gin-gonic/gin"
)

type MpRouter struct {
}

// 初始化登录控制器
func (MpRouter) Init(Group *gin.RouterGroup) {
	c := public.MpController{}
	apiRouter := Group.Group("/mp")
	{
		apiRouter.POST("/login", controller.BindingFunc(c.Login))
		apiRouter.POST("/bind-mobile", controller.BindingFunc(c.Login))
		apiRouter.POST("/bind-email", controller.BindingFunc(c.Login))
		apiRouter.GET("/user-info", controller.BindingFunc(c.Login))
	}
}
