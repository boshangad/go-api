package public

import "github.com/gin-gonic/gin"

type EmailRouter struct {
}

func (EmailRouter) Init(Group *gin.RouterGroup) {
	apiRouter := Group.Group("/email")
	{
		apiRouter.POST("/send")
	}
}
