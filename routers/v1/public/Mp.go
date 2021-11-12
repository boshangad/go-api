package public

import "github.com/gin-gonic/gin"

type MpRouter struct {
}

func (MpRouter) Init(Group *gin.RouterGroup) {
	apiRouter := Group.Group("/mp")
	{
		apiRouter.POST("/login", func(c *gin.Context) {

		})
		apiRouter.POST("/bind-mobile", func(c *gin.Context) {

		})
		apiRouter.POST("/bind-email", func(c *gin.Context) {

		})
		apiRouter.GET("/user-info", func(c *gin.Context) {

		})
	}
}
