package v1

import (
	"github.com/boshangad/go-api/routers/v1/public"
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

var ApiGroup = new(GroupApi)
