package public

import "github.com/gin-gonic/gin"

type CaptchaRouter struct {
}

func (CaptchaRouter) Init(Group *gin.RouterGroup) {
	apiRouter := Group.Group("/captcha")
	{
		apiRouter.GET("/view")
		apiRouter.POST("/image")
	}
}
