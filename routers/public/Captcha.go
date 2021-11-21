package public

import (
	"github.com/boshangad/v1/api/public"
	"github.com/boshangad/v1/app/controller"
	"github.com/gin-gonic/gin"
)

type CaptchaRouter struct {
}

// 初始化验证码控制器
func (CaptchaRouter) Init(Group *gin.RouterGroup) {
	c := public.CaptchaController{}
	apiRouter := Group.Group("/captcha")
	{
		apiRouter.GET("", controller.BindingFunc(c.View))
		apiRouter.GET("/:name", controller.BindingFunc(c.View))
	}
}
