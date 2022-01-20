package captchaService

import (
	"strings"

	"github.com/boshangad/v1/global"
	"github.com/dchest/captcha"
)

func init() {
	captcha.SetCustomStore(getCustomStore())
}

func getCustomStore() captcha.Store {
	switch strings.ToLower(global.Config.Captcha.Store) {
	case "redis":
		return NewRedisStore(global.Config.Cache)
	case "file":
	case "memory":
	default:
	}
	return nil
}

func NewCaptcha() *captcha.Image {
	return captcha.NewImage(captcha.NewLen(7), captcha.RandomDigits(9), 300, 200)
}

func NewAudio() *captcha.Audio {
	return captcha.NewAudio(captcha.NewLen(9), captcha.RandomDigits(9), "zh")
}
