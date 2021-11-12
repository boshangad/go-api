package public

import "github.com/gin-gonic/gin"

type AccountRouter struct {
}

func (AccountRouter) Init(Group *gin.RouterGroup) {
	apiRouter := Group.Group("/account")
	{
		apiRouter.POST("/login", func(c *gin.Context) {

		})
		apiRouter.POST("/register", func(c *gin.Context) {

		})
	}
}
