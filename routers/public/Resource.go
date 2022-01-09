package public

import (
	"github.com/boshangad/v1/api/public"
	"github.com/boshangad/v1/app/controller"
	"github.com/gin-gonic/gin"
)

type ResourceRouter struct {
}

// 初始化账号控制器
func (ResourceRouter) Init(Group *gin.RouterGroup) {
	p := public.ResourceController{}
	apiRouter := Group.Group("/resource")
	{
		apiRouter.GET("/image", controller.BindingFunc(p.Image))
		apiRouter.GET("/download", controller.BindingFunc(p.Download))
		apiRouter.POST("/upload", controller.BindingFunc(p.Upload))
		apiRouter.GET("/us3-token", controller.BindingFunc(p.Us3Token))
	}
}
