package routers

import (
	"github.com/boshangad/go-api/app/controllers"
	"github.com/gin-gonic/gin"
)

type mainController struct {
	routerStruct
}

func (c *mainController) Init(e *gin.Engine, g *gin.RouterGroup) *mainController {
	c.SetEngine(e).SetGroup(g)
	c.Account()
	c.Mp()
	c.Captcha()
	c.Email()
	c.Sms()
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

func (c mainController) Sms() {
	g := controllers.SmsController{}
	groups := c.Group.Group("/sms")
	{
		groups.POST("/send", func(c *gin.Context) {
			g.Init(c)
			g.Send()
		})
	}
	return
}

func (c mainController) Email() {
	g := controllers.EmailController{}
	groups := c.Group.Group("/email")
	{
		groups.POST("/send", func(c *gin.Context) {
			g.Init(c)
			g.Send()
		})
	}
	return
}
