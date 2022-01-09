package routers

import (
	"github.com/boshangad/v1/app/controller"
	"github.com/boshangad/v1/routers/public"
	"github.com/gin-gonic/gin"
)

type GroupApi struct {
	public.RouterApi
}

// 注册路由
func (that *GroupApi) Init(group *gin.RouterGroup) {
	// 注册相关路由
	group.Use(controller.BindingFunc(nil))
	that.Captcha.Init(group)
	that.Account.Init(group)
	that.Mp.Init(group)
	that.Sms.Init(group)
	that.Email.Init(group)
	that.Resource.Init(group)
}

var V1 = new(GroupApi)
