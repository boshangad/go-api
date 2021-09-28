package public

import "github.com/gin-gonic/gin"

type AccountRouter struct {
}

func (AccountRouter) Init(Group *gin.RouterGroup) {
	apiRouter := Group.Group("/account")
	{
		apiRouter.POST("/login")
		apiRouter.POST("/register")
	}
}
