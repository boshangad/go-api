package controllers

import (
	"crypto/sha1"
	"encoding/hex"
	"github.com/tend/wechatServer/app/services"
	"github.com/tend/wechatServer/core/global"
	"strconv"
)

type MpController struct {
	Controller
}

// Login 微信code登录
// @route login [POST]
func (mp MpController) Login()  {
	code := mp.getParamWithString("code")
	if code == "" {
		mp.jsonOut(global.ErrMissLoginParams, "miss param #code", nil)
		return
	}
	mpService := services.NewsMpService()
	token, err := mpService.SetApp(mp.App).LoginByCode(code)
	if err != nil {
		mp.jsonOut(global.ErrNotice, err.Error(), nil)
		return
	}
	mp.jsonOut(global.ErrSuccess, "操作成功", services.StructLoginSuccess{
		AccessToken: token,
		ExpireTime: int64(mp.AppUserToken.ExpireTime),
		IsLoadUserProfile: mp.AppUser.IsLoadUserProfile,
		IsBindUser: false,
	})
}

// SetUserProfile 设置用户信息
// @route user-profile [POST]
func (mp MpController) SetUserProfile() {
	rawData := mp.Context.PostForm("rawData")
	signature := mp.Context.PostForm("signature")
	encryptedData := mp.Context.PostForm("encryptedData")
	iv := mp.Context.PostForm("iv")
	t := sha1.New()
	t.Write([]byte(rawData + mp.AppUser.SessionKey))
	checkSign := t.Sum(nil)
	if hex.EncodeToString(checkSign) != signature {
		mp.jsonOut(global.ErrNotice, "操作失败，签名验证不通过", nil)
		return
	}
	mpService := services.NewsMpService()
	mpService.AppUser = mp.AppUser
	mpService.App = mp.App
	data := mpService.SetUserProfile(encryptedData, iv)
	mp.jsonOut(global.ErrSuccess, "操作成功", data)
}

// Info 获取用户信息
// @route info [GET]
func (mp MpController) Info() {
	mp.jsonOut(global.ErrSuccess, "操作成功", returnAppUser{
		Id: mp.AppUser.ID,
		IsLoadUserProfile: mp.AppUser.IsLoadUserProfile,
		Nickname: mp.AppUser.Nickname,
		Avatar: mp.AppUser.Avatar,
		Gender: mp.AppUser.Gender,
		County: mp.AppUser.County,
		CountryCode: mp.AppUser.CountryCode,
		Province: mp.AppUser.Province,
		City: mp.AppUser.City,
		Language: mp.AppUser.Language,
	})
}

// Qrcode 获取微信小程序码-码数量较少的业务场景
// @route qrcode [POST]
func (mp MpController) Qrcode()  {
	path := mp.getParamWithString("path")
	width := mp.getParamWithString("width")
	var widthInt int
	var err error
	if width != "" {
		widthInt, err = strconv.Atoi(width)
		if err != nil {
			mp.jsonOut(global.ErrNotice, "", nil)
			return
		}
	}
	services.NewsMpService().SetApp(mp.App).Qrcode(path, widthInt)

}