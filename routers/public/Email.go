package public

import "github.com/gin-gonic/gin"

type EmailRouter struct {
}

// 初始化邮件控制器
func (EmailRouter) Init(Group *gin.RouterGroup) {
	apiRouter := Group.Group("/email")
	{
		apiRouter.POST("/send", func(c *gin.Context) {

		})
	}
}
