package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/boshangad/go-api/app/controllers"
)

type mainController struct {
	routerStruct
}

func (c *mainController) Init(e *gin.Engine, g *gin.RouterGroup) *mainController {
	c.SetEngine(e).SetGroup(g)
	c.Account()
	c.Mp()
	c.Captcha()
	return c
}

func (c mainController) Account() mainController {
	g := controllers.AccountController{}
	groups := c.Group.Group("/account")
	{ 
		groups.POST("/login", func(c *gin.Context) {
			g.Init(c)
			g.Login()
		})
		groups.POST("/register", func(c *gin.Context) {
			g.Init(c)
			g.Register()
		})
	}
	return c
}

func (c mainController) Mp() {
	g := controllers.MpController{}
	groups := c.Group.Group("/mp")
	{ 
		groups.POST("/login", func(c *gin.Context) {
			g.Init(c)
			g.Login()
		})
		groups.POST("/user-profile", func(c *gin.Context) {
			g.Init(c)
			g.SetUserProfile()
		})
		groups.GET("/info", func(c *gin.Context) {
			g.Init(c)
			g.Info()
		})
		groups.POST("/qrcode", func(c *gin.Context) {
			g.Init(c)
			g.Qrcode()
		})
	}
	return
}

func (c mainController) Captcha() {
	g := controllers.CaptchaController{}
	groups := c.Group.Group("/captcha")
	{
		groups.GET("/image", func(c *gin.Context) {
			g.Init(c)
			g.Image()
		})
		groups.GET("/audit", func(c *gin.Context) {
			g.Init(c)
			g.Audit()
		})
	}
	return
}