package public

import (
	"github.com/boshangad/go-api/api/v1/public"
	"github.com/gin-gonic/gin"
)

type CaptchaRouter struct {
}

func (CaptchaRouter) Init(Group *gin.RouterGroup) {
	controller := public.CaptchaController{}
	apiRouter := Group.Group("/captcha")
	{
		apiRouter.GET("/image",controller.Image)
	}
}
