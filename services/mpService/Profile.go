package mpService

import (
	"context"
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"github.com/boshangad/go-api/ent"
	"github.com/boshangad/go-api/global"
	"strconv"
	"time"
)

type Profile struct {
	AppUser *ent.AppUser
	RawData string
	Signature string
	EncryptedData string
	Iv string
}

func (that *Profile) SetAppUser(appUser *ent.AppUser) *Profile {
	that.AppUser = appUser
	return that
}

func (that Profile) Check() (pass bool) {
	t := sha1.New()
	t.Write([]byte(that.RawData + that.AppUser.SessionKey))
	checkSign := t.Sum(nil)
	if hex.EncodeToString(checkSign) != that.Signature {
		return false
	}
	return true
}

func (that *Profile) Save() error {
	var ctx = context.Background()
	mp := DefaultWechat.MiniProgram(that.AppUser.Edges.App)
	data, err := mp.GetEncryptor().Decrypt(that.AppUser.SessionKey, that.EncryptedData, that.Iv)
	if err != nil {
		global.G_LOG.Error("weChat user information decryption failed:" + err.Error())
		return nil
	}
	// 协程下载文件
	go func(appUser *ent.AppUser) {
		avatarUrl := "avatar/" + strconv.FormatUint(appUser.ID, 10) + ".jpg"
		_, _ = appUser.Update().
			SetAvatar(avatarUrl).
			Save(ctx)
	}(that.AppUser)
	str,_ := json.Marshal(data.Watermark)
	appUser, err := that.AppUser.Update().
		SetAvatarURL(data.AvatarURL).
		SetGender(uint(data.Gender)).
		SetLanguage(data.Language).
		SetCounty(data.Country).
		SetCountryCode(data.CountryCode).
		SetProvince(data.Province).
		SetCity(data.City).
		SetOpenID(data.OpenID).
		SetUnionid(data.UnionID).
		SetPhoneNumber(data.PhoneNumber).
		SetPurePhoneNumber(data.PurePhoneNumber).
		SetNickname(data.NickName).
		SetIsLoadUserProfile(true).
		SetLoadUserProfileTime(uint64(time.Now().Unix())).
		SetWatermark(string(str)).
		Save(ctx)
	if err != nil {
		global.G_LOG.Error("failed to update table appUser:" + err.Error())
		return nil
	}
	that.AppUser = appUser
	return nil
}