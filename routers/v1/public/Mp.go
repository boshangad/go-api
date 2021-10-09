package public

import "github.com/gin-gonic/gin"

type MpRouter struct {
}

func (MpRouter) Init(Group *gin.RouterGroup) {
	apiRouter := Group.Group("/mp")
	{
		apiRouter.POST("/login")
		apiRouter.POST("/bind-mobile")
		apiRouter.POST("/bind-email")
		apiRouter.GET("/user-info")
	}
}
