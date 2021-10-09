package public

import (
	"encoding/base64"
	"github.com/boshangad/go-api/core/mvvc"
	"github.com/boshangad/go-api/global"
	"github.com/boshangad/go-api/services/captchaService"
	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
	"net/http"
)

type CaptchaController struct {
	mvvc.Controller
}

func (that CaptchaController) Image(c *gin.Context) {
	that.Init(c)
	var data = captchaService.WHSize{}
	err := that.ShouldBind(&data)
	if data.Width == 0 {
		data.Width = 160
	}
	if data.Height == 0 {
		data.Height = 45
	}
	color := base64Captcha.RandColor()
	driver := base64Captcha.NewDriverString(
		data.Height,
		data.Width,
		0,
		base64Captcha.OptionShowSlimeLine,
		6,
		base64Captcha.TxtSimpleCharaters,
		&color,
		base64Captcha.DefaultEmbeddedFonts,
		[]string{},
	)
	captcha := base64Captcha.NewCaptcha(driver, base64Captcha.DefaultMemStore)
	id, b64s, err := captcha.Generate()
	if err != nil {
		that.JsonOut(global.ErrNotice, "参数错误", nil)
		return
	}
	b64s = b64s[22:]
	// 成图片文件并把文件写入到buffer
	ddd, _:= base64.StdEncoding.DecodeString(b64s)
	that.Context.SetCookie("captchaKey", id, 0, "/captcha", that.Context.Request.Host, true, true)
	that.Context.Data(http.StatusOK, "image/png", ddd)
}