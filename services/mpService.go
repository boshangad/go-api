package services

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/boshangad/go-api/ent"
	"github.com/boshangad/go-api/ent/appuser"
	"github.com/boshangad/go-api/ent/appusertoken"
	"github.com/boshangad/go-api/global/db"
	"github.com/boshangad/go-api/utils"
	"github.com/google/uuid"
	"github.com/silenceper/wechat/v2/miniprogram/encryptor"
	"github.com/silenceper/wechat/v2/miniprogram/qrcode"
	"io/ioutil"
	"log"
	"time"
)

type mpService struct {
	App *ent.App `json:"-"`
	AppUser *ent.AppUser `json:"-"`
	AppUserToken *ent.AppUserToken `json:"-"`
}

// SetApp SetAppByAlias 设置应用通过别名
func (_this *mpService) SetApp(app *ent.App) *mpService {
	_this.App = app
	return _this
}

// LoginByCode 通过微信Code登录
func (_this *mpService) LoginByCode(code string) (string, error) {
	mp := GetMiniProgram(_this.App)
	codeSession, err := mp.GetAuth().Code2Session(code)
	if err != nil {
		log.Println("登录code无效", err)
		return "", errors.New("login error, code invalid")
	}
	ctx := context.Background()
	au := db.DefaultClient().AppUser
	appUserModel, err := au.Query().Where(appuser.And(
		appuser.AppIDEQ(_this.App.ID),
		appuser.OpenIDEQ(codeSession.OpenID),
	)).First(ctx)
	if err != nil {
		// 表示查询失败
		if !ent.IsNotFound(err) {
			log.Println("数据查询错误：", err)
			return "", err
		}
		appUserModel, err = au.Create().
			SetAppID(_this.App.ID).
			SetOpenID(codeSession.OpenID).
			Save(ctx)
		if err != nil {
			log.Println("新增appUser表失败：", err)
			return "", err
		}
	}
	appUserModel, err = appUserModel.Update().
		SetSessionKey(codeSession.SessionKey).
		SetUnionid(codeSession.UnionID).
		SetLastLoginTime(uint64(time.Now().Unix())).
		Save(ctx)
	if err != nil {
		log.Println("更新appUser表失败：", err)
		return "", err
	}
	_this.AppUser = appUserModel
	return _this.BuildTokenWithSave()
}

// BuildTokenWithSave 创建登录Token并且保存
func (_this *mpService) BuildTokenWithSave() (string, error) {
	if _this.AppUser == nil || _this.App == nil {
		log.Println("操作异常，未被定义的用户")
		return "", errors.New("操作异常，未被定义的用户")
	}
	ctx := context.Background()
	clientAUT := db.DefaultClient().AppUserToken
	_, err := clientAUT.Delete().
		Where(appusertoken.And(
			appusertoken.AppIDEQ(_this.App.ID),
			appusertoken.AppUserIDEQ(_this.AppUser.ID),
		)).
		Exec(ctx)
	if err != nil {
		log.Println(err)
		return "", err
	}
	tokenModel, err := clientAUT.Create().
		SetAppID(_this.App.ID).
		SetUUID(uuid.New()).
		SetExpireTime(uint64(time.Now().Unix() + 3600 * 2)).
		SetAppUserID(_this.AppUser.ID).
		Save(ctx)
	if err != nil {
		return "", err
	}
	_this.AppUserToken = tokenModel
	// 生成token
	encrypt := ""
	//ts := tokenService{}
	//encrypt := ts.EncryptString(tokenModel.UUID.String())
	return encrypt, nil
}

func (_this mpService) SetUserProfile(encryptedData string, iv string) *encryptor.PlainData {
	mp := GetMiniProgram(_this.App)
	data, err := mp.GetEncryptor().Decrypt(_this.AppUser.SessionKey, encryptedData, iv)
	if err != nil {
		log.Println("解析MP应用用户信息失败", err)
		return nil
	}
	// 协程下载文件
	avatarUrl := "/avatar/" + data.OpenID + "_132.png"
	avatarPath := "" + avatarUrl
	go utils.DownLoadFile(data.AvatarURL, avatarPath)

	str,_ := json.Marshal(data.Watermark)
	appUser, err := _this.AppUser.Update().
		SetAvatar(avatarUrl).
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
		Save(context.Background())
	if err != nil {
		log.Println("保存user信息失败", err)
		return nil
	}
	_this.AppUser = appUser
	return data
}

func (_this mpService) Qrcode(path string, width int) (string, error) {
	if width == 0 {
		width = 460
	} else if width > 1280 || width < 280 {
		return "", errors.New("小程序码宽度必须在280与1280之间")
	}
	wc := GetMiniProgram(_this.App)
	content, err := wc.GetQRCode().GetWXACode(qrcode.QRCoder{
		Path: path,
		Width: width,
	})
	if err != nil {
		log.Println("获取小程序码失败", err)
		return "", nil
	}
	fileName := "C:\\Users\\huanghu\\Desktop\\avatar\\test.txt"
	err = ioutil.WriteFile(fileName, content, 0644)
	return fileName, err
}

// NewsMpService NewsMp 实例化mpService
func NewsMpService() *mpService {
	return &mpService{}
}