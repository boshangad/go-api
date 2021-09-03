package controllers

import (
	"github.com/dchest/captcha"
	"github.com/boshangad/go-api/core/global"
	"strconv"
)

type CaptchaController struct {
	Controller
}

func (gh CaptchaController) Image() {
	width := gh.getParamWithString("width")
	height := gh.getParamWithString("height")
	w, _ := strconv.Atoi(width)
	h, _ := strconv.Atoi(height)
	if w == 0 {
		w = 160
	}
	if h == 0 {
		h = 45
	}
	text := captcha.New()
	err := captcha.WriteImage(gh.Context.Writer, text, w, h)
	if err != nil {
		gh.jsonOut(global.ErrNotice, "参数错误", nil)
		return
	}
}

func (gh CaptchaController) Audit() {
	lang := gh.getParamWithString("lang")
	text := captcha.New()
	err := captcha.WriteAudio(gh.Context.Writer, text, lang)
	if err != nil {
		gh.jsonOut(global.ErrNotice, "参数错误", nil)
		return
	}
}