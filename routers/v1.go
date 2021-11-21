package routers

import (
	"github.com/boshangad/v1/routers/public"
	"github.com/gin-gonic/gin"
)

type GroupApi struct {
	public.RouterApi
}

// 注册路由
func (that *GroupApi) Init(Group *gin.RouterGroup) {
	that.Captcha.Init(Group)
	that.Account.Init(Group)
	that.Mp.Init(Group)
	that.Sms.Init(Group)
	that.Email.Init(Group)
}

var V1 = new(GroupApi)

