package public

import (
	"encoding/base64"
	"math/rand"
	"net/http"
	"strings"
	"time"

	"github.com/boshangad/go-api/core/mvvc"
	"github.com/boshangad/go-api/global"
	"github.com/boshangad/go-api/services/captchaService"
	"github.com/boshangad/go-api/utils"
	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
)

type CaptchaController struct {
	mvvc.Controller
}

// 验证码详情
func (that CaptchaController) View(c *gin.Context) {
	t := strings.ToLower(strings.TrimSpace(c.Param("type")))
	w := utils.GetInt(c.Param("w"))
	h := utils.GetInt(c.Param("h"))
	that.Init(c)
	var data = captchaService.WHSize{}
	err := that.ShouldBind(&data)
	if err != nil {
		that.JsonOut(global.ErrNotice, err.Error(), nil)
		return
	}
	if t == "" {
		t = "image"
	}
	// 获取验证码配置
	captcha, ok := global.G_CONFIG.Captcha[t]
	if !ok {
		that.JsonOut(global.ErrNotice, "related configuration parameters", nil)
		return
	}
	if w != 0 {
		data.Width = w
	}
	if h != 0 {
		data.Height = h
	}
	if data.Width == 0 {
		data.Width = 160
	}
	if data.Height == 0 {
		data.Height = 45
	}
	var (
		driver         base64Captcha.Driver
		color          = base64Captcha.RandColor()
		wordLength     = captcha.Length
		mimeTypeString = base64Captcha.MimeTypeImage
	)
	// 随机长度
	if wordLength <= 0 {
		rand.Seed(time.Now().UnixNano())
		wordLength = rand.Intn(captcha.MaxLength-captcha.MinLength) + captcha.MinLength
	}
	switch strings.ToLower(strings.TrimSpace(captcha.Type)) {
	case "audio":
		mimeTypeString = base64Captcha.MimeTypeAudio
		driver = base64Captcha.NewDriverAudio(
			wordLength,       // 字符长度
			captcha.Language, // 使用语言
		)
	case "chinese":
		driver = base64Captcha.NewDriverChinese(
			data.Height,                        // 高度
			data.Width,                         // 宽度
			captcha.NoiseCount,                 // 噪点数
			base64Captcha.OptionShowSlimeLine,  // 显示线条数
			wordLength,                         // 字符长度
			base64Captcha.TxtSimpleCharaters,   // 参与随机文案
			&color,                             // 背景颜色
			base64Captcha.DefaultEmbeddedFonts, // 字体存储
			captcha.Fonts,                      // 使用字体
		)
	case "math":
		driver = base64Captcha.NewDriverMath(
			data.Height,                        // 高度
			data.Width,                         // 宽度
			captcha.NoiseCount,                 // 噪点数
			wordLength,                         // 验证字符长度
			&color,                             // 背景色
			base64Captcha.DefaultEmbeddedFonts, // 字体存储
			captcha.Fonts,                      // 使用字体
		)
	case "digit":
		driver = base64Captcha.NewDriverDigit(
			data.Height,                       // 高度
			data.Width,                        // 宽度
			wordLength,                        // 字符长度
			base64Captcha.OptionShowSlimeLine, // 倾斜角度
			captcha.NoiseCount,                // 噪点数
		)
	case "language":
		driver = base64Captcha.NewDriverLanguage(
			data.Height,                        // 高度
			data.Width,                         // 宽度
			captcha.NoiseCount,                 // 噪点数
			base64Captcha.OptionShowSlimeLine,  // 显示线条数
			wordLength,                         // 字符长度
			&color,                             // 背景色
			base64Captcha.DefaultEmbeddedFonts, // 字体存储
			nil,                                // 字体
			captcha.Language,                   // 语言
		)
	default:
		driver = base64Captcha.NewDriverString(
			data.Height,                        // 高度
			data.Width,                         // 宽度
			captcha.NoiseCount,                 // 噪点数
			base64Captcha.OptionShowSlimeLine,  // 线条数量
			wordLength,                         // 字符长度
			base64Captcha.TxtSimpleCharaters,   // 字符来源
			&color,                             // 背景颜色
			base64Captcha.DefaultEmbeddedFonts, // 字体存储
			captcha.Fonts,                      // 使用字体
		)
	}
	var captchaStruct = base64Captcha.NewCaptcha(driver, base64Captcha.DefaultMemStore)
	id, b64s, err := captchaStruct.Generate()
	if err != nil {
		that.JsonOut(global.ErrNotice, "verification code generation failed", nil)
		return
	}
	b64s = b64s[22:]
	// 成图片文件并把文件写入到buffer
	ddd, _ := base64.StdEncoding.DecodeString(b64s)
	that.Context.SetCookie("captchaKey", id, 0, "/captcha", that.Context.Request.Host, true, true)
	that.Context.Data(http.StatusOK, mimeTypeString, ddd)
}
