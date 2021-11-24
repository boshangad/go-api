package config

import (
	"image/color"
	"strings"

	"github.com/boshangad/v1/app/helpers"
	"github.com/mojocn/base64Captcha"
)

type Captcha struct {
	mimeType string `json:"-" yaml:"-"`
	// 失效时间，单位秒，-1表示不限制
	ExpireTime int `json:"expireTime,omitempty" yaml:"expireTime"`
	// cookie配置
	CookieName     string `json:"cookieName,omitempty" yaml:"cookieName"`
	CookiePath     string `json:"cookiePath,omitempty" yaml:"cookiePath"`
	CookieSecure   bool   `json:"cookieSecure,omitempty" yaml:"cookieSecure"`
	CookieHttpOnly bool   `json:"cookieHttpOnly,omitempty" yaml:"cookieHttpOnly"`
	// 验证码类型
	// image 图片、audio 音频、chinese 中文、math 数学、digit 点位、language 语言
	Type string `json:"type,omitempty" yaml:"type"`
	// 验证码宽度
	Width int `json:"width,omitempty" yaml:"width"`
	// 验证码最低宽度
	MinWidth int `json:"minWidth,omitempty" yaml:"minWidth"`
	// 验证码最大宽度
	MaxWidth int `json:"maxWidth,omitempty" yaml:"maxWidth"`
	// 验证码高度
	Height int `json:"height,omitempty" yaml:"height"`
	// 验证码最底高度
	MinHeight int `json:"minHeight,omitempty" yaml:"minHeight"`
	// 验证码最大高度
	MaxHeight int `json:"maxHeight,omitempty" yaml:"maxHeight"`
	// 验证码最大长度
	Length int `json:"length,omitempty" yaml:"length"`
	// 验证码噪点数
	NoiseCount int `json:"nosiseCount,omitempty" yaml:"nosiseCount"`
	// 验证码最小噪点数
	MinNoiseCount int `json:"minNosiseCount,omitempty" yaml:"minNosiseCount"`
	// 验证码最大噪点数
	MaxNoiseCount int `json:"maxNosiseCount,omitempty" yaml:"maxNosiseCount"`
	// 语音验证码语言、或语言验证码有效
	Language string `json:"language,omitempty" yaml:"language"`
	// 生成验证码的字体
	Fonts []string `json:"fonts,omitempty" yaml:"fonts"`
	// 背景颜色
	Color string `json:"color,omitempty" yaml:"color"`
}

// 实例生成验证码服务对象
func (that *Captcha) Driver() (driver base64Captcha.Driver) {
	var (
		colorRgba  = base64Captcha.RandColor()
		R, G, B, A uint8
	)
	that.mimeType = base64Captcha.MimeTypeImage
	if that.Color != "" {
		strings.SplitN(that.Color, ",", 4)
		colorRgba = color.RGBA{
			R: R,
			G: G,
			B: B,
			A: A,
		}
	}
	// 宽度锁定
	if that.Width < 1 {
		if that.MaxWidth < 1 {
			that.Width = 120
		} else {
			that.Width = helpers.RandomRangeInt(1, that.MaxWidth)
		}
	}
	// 高度锁定
	if that.Height < 1 {
		if that.MaxHeight < 1 {
			that.Height = 60
		} else {
			that.Height = helpers.RandomRangeInt(1, that.MaxHeight)
		}
	}
	// 噪点数
	if that.NoiseCount < 1 {
		if that.MinNoiseCount > 0 && that.MaxNoiseCount > 0 {
			that.NoiseCount = helpers.RandomRangeInt(that.MaxNoiseCount, that.MaxNoiseCount)
		} else if that.MinNoiseCount > 0 {
			that.NoiseCount = that.MinNoiseCount
		} else if that.MaxNoiseCount > 0 {
			that.NoiseCount = helpers.RandomRangeInt(0, that.MaxNoiseCount)
		}
	}
	// 生成验证码
	switch strings.TrimSpace(that.Type) {
	case "audio":
		driver = base64Captcha.NewDriverAudio(
			that.Length,
			that.Language,
		)
		that.mimeType = base64Captcha.MimeTypeAudio
	case "chinese":
		driver = base64Captcha.NewDriverChinese(
			that.Height,
			that.Width,
			that.NoiseCount,
			base64Captcha.OptionShowSlimeLine,
			that.Length,
			base64Captcha.TxtSimpleCharaters,
			&colorRgba,
			base64Captcha.DefaultEmbeddedFonts,
			that.Fonts,
		)
	case "math":
		driver = base64Captcha.NewDriverMath(
			that.Height,
			that.Width,
			that.NoiseCount,
			that.Length,
			&colorRgba,
			base64Captcha.DefaultEmbeddedFonts,
			that.Fonts,
		)
	case "digit":
		driver = base64Captcha.NewDriverDigit(
			that.Height,
			that.Width,
			that.Length,
			base64Captcha.OptionShowSlimeLine,
			that.NoiseCount,
		)
	case "language":
		driver = base64Captcha.NewDriverLanguage(
			that.Height,
			that.Width,
			that.NoiseCount,
			base64Captcha.OptionShowSlimeLine,
			that.Length,
			&colorRgba,
			base64Captcha.DefaultEmbeddedFonts,
			nil,
			that.Language,
		)
	default:
		driver = base64Captcha.NewDriverString(
			that.Height,
			that.Width,
			that.NoiseCount,
			base64Captcha.OptionShowSlimeLine,
			that.Length,
			base64Captcha.TxtSimpleCharaters,
			&colorRgba,
			base64Captcha.DefaultEmbeddedFonts,
			that.Fonts,
		)
	}
	return
}

// 生成的验证码服务的mime类型
func (that Captcha) DriverMimeType() string {
	return that.mimeType
}
