package controllers

import (
	"crypto/sha1"
	"encoding/hex"
	"github.com/boshangad/go-api/app/services"
	"github.com/boshangad/go-api/core/global"
	"github.com/boshangad/go-api/core/mvvc"
	"github.com/boshangad/go-api/ent"
	"strconv"
)

type MpController struct {
	mvvc.Controller
}

// Login 微信code登录
// @route login [POST]
func (that MpController) Login()  {
	code := that.GetParamWithString("code")
	if code == "" {
		that.JsonOut(global.ErrMissLoginParams, "miss param #code", nil)
		return
	}
	mpService := services.NewsMpService()
	token, err := mpService.SetApp(that.App).LoginByCode(code)
	if err != nil {
		that.JsonOut(global.ErrNotice, err.Error(), nil)
		return
	}
	that.JsonOut(global.ErrSuccess, "操作成功", services.StructLoginSuccess{
		AccessToken:       token,
		ExpireTime:        int64(that.AppUserToken.ExpireTime),
		IsLoadUserProfile: that.AppUser.IsLoadUserProfile,
		IsBindUser:        false,
	})
}

// SetUserProfile 设置用户信息
// @route user-profile [POST]
func (that MpController) SetUserProfile() {
	rawData := that.Context.PostForm("rawData")
	signature := that.Context.PostForm("signature")
	encryptedData := that.Context.PostForm("encryptedData")
	iv := that.Context.PostForm("iv")
	t := sha1.New()
	t.Write([]byte(rawData + that.AppUser.SessionKey))
	checkSign := t.Sum(nil)
	if hex.EncodeToString(checkSign) != signature {
		that.JsonOut(global.ErrNotice, "操作失败，签名验证不通过", nil)
		return
	}
	mpService := services.NewsMpService()
	mpService.AppUser = that.AppUser
	mpService.App = that.App
	data := mpService.SetUserProfile(encryptedData, iv)
	that.JsonOut(global.ErrSuccess, "操作成功", data)
}

// Info 获取用户信息
// @route info [GET]
func (that MpController) Info() {
	that.JsonOut(global.ErrSuccess, "操作成功", struct {
		*ent.AppUser
	}{AppUser: that.AppUser})
}

// Qrcode 获取微信小程序码-码数量较少的业务场景
// @route qrcode [POST]
func (that MpController) Qrcode()  {
	path := that.GetParamWithString("path")
	width := that.GetParamWithString("width")
	var widthInt int
	var err error
	if width != "" {
		widthInt, err = strconv.Atoi(width)
		if err != nil {
			that.JsonOut(global.ErrNotice, "", nil)
			return
		}
	}
	services.NewsMpService().SetApp(that.App).Qrcode(path, widthInt)

}