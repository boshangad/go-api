package v1

import (
	v1 "github.com/boshangad/go-api/api/v1"
	"github.com/gin-gonic/gin"
))

type MpRouter struct {
}

func (MpRouter) MpRouter(Router gin.RouterGroup) {
	var mpController = v1.ApiGroupApp.Public.MpController
	{
		Router.GET("", mpController.Login)
	}
}
