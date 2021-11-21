package public

import "github.com/gin-gonic/gin"

type SmsRouter struct {
}

// 初始化短信控制器
func (SmsRouter) Init(Group *gin.RouterGroup) {
	apiRouter := Group.Group("/sms")
	{
		apiRouter.POST("/send", func(c *gin.Context) {

		})
	}
}
