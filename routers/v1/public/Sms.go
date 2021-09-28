package public

import "github.com/gin-gonic/gin"

type SmsRouter struct {
}

func (SmsRouter) Init(Group *gin.RouterGroup) {
	apiRouter := Group.Group("/sms")
	{
		apiRouter.POST("/send")
	}
}
