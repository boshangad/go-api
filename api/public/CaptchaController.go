package public

import (
	"github.com/boshangad/v1/app/controller"
	"github.com/mojocn/base64Captcha"
)

type CaptchaController struct {
}

var captchaStore = base64Captcha.NewMemoryStore(10, 10000000000)

// 验证码详情
func (CaptchaController) View(c *controller.Context) {
	var (
	// name          = strings.ToLower(strings.TrimSpace(c.Param("name")))
	// width, _      = strconv.ParseInt(c.Query("width"), 0, 64)
	// height, _     = strconv.ParseInt(c.Query("height"), 0, 64)
	// t             = c.DefaultQuery("type", name)
	// scope         = c.DefaultQuery("scope", "login")
	// captchaStruct *base64Captcha.Captcha
	// captcha       config.Captcha
	)
	// // 获取验证码配置
	// captcha = global.Config.Captcha
	// if captcha.Type == "" && t != "" {
	// 	captcha.Type = t
	// }
	// if width > 0 {
	// 	captcha.Width = helpers.Int(width)
	// }
	// if height > 0 {
	// 	captcha.Height = helpers.Int(height)
	// }
	// driver := captcha.Driver()
	// captchaStruct = base64Captcha.NewCaptcha(driver, captchaStore)
	// id, b64s, err := captchaStruct.Generate()
	// if err != nil {
	// 	global.Log.Warn("captcha generation failed", zap.Error(err), zap.String("name", name))
	// 	c.JsonOut(global.ErrNotice, "code generation failed, please try again later.", nil)
	// 	return
	// }
	// b64s = b64s[22:]
	// // 成图片文件并把文件写入到buffer
	// contentBytes, _ := base64.StdEncoding.DecodeString(b64s)
	// c.SetCookie(
	// 	"captcha:"+scope,
	// 	id,
	// 	int(captcha.ExpireTime),
	// 	"/",
	// 	strings.Split(c.Request.Host, ":")[0],
	// 	true,
	// 	true,
	// )
	// c.Data(http.StatusOK, captcha.DriverMimeType(), contentBytes)
}
