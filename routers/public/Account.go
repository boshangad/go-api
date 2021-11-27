package public

import (
	"github.com/boshangad/v1/api/public"
	"github.com/boshangad/v1/app/controller"
	"github.com/gin-gonic/gin"
)

type AccountRouter struct {
}

// 初始化账号控制器
func (AccountRouter) Init(Group *gin.RouterGroup) {
	p := public.AccountController{}
	apiRouter := Group.Group("/account")
	{
		apiRouter.POST("/login", controller.BindingFunc(p.Login))
		apiRouter.POST("/register", controller.BindingFunc(nil))
	}
}
