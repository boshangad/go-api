package controllers

import (
	"github.com/boshangad/go-api/core/global"
	"github.com/boshangad/go-api/core/mvvc"
	"github.com/dchest/captcha"
	"strconv"
)

type CaptchaController struct {
	mvvc.Controller
}

func (that CaptchaController) Image() {
	width := that.GetParamWithString("width")
	height := that.GetParamWithString("height")
	w, _ := strconv.Atoi(width)
	h, _ := strconv.Atoi(height)
	if w == 0 {
		w = 160
	}
	if h == 0 {
		h = 45
	}
	text := captcha.New()
	err := captcha.WriteImage(that.Context.Writer, text, w, h)
	if err != nil {
		that.JsonOut(global.ErrNotice, "参数错误", nil)
		return
	}
}

func (that CaptchaController) Audit() {
	lang := that.GetParamWithString("lang")
	text := captcha.New()
	err := captcha.WriteAudio(that.Context.Writer, text, lang)
	if err != nil {
		that.JsonOut(global.ErrNotice, "参数错误", nil)
		return
	}
}