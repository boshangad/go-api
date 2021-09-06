package userService

import (
	"context"
	"errors"
	"github.com/boshangad/go-api/core/db"
	"github.com/boshangad/go-api/core/mvvc"
	"github.com/boshangad/go-api/ent"
	"github.com/boshangad/go-api/ent/user"
	"github.com/boshangad/go-api/utils"
	"log"
	"strings"
)

type structUserRegister struct {
	mvvc.BaseServiceStruct
	Username string `json:"username,omitempty" form:"username" filter:"trim"`
	Password string `json:"password,omitempty" form:"password" filter:"trim"`
	DialCode string `json:"dial_code,omitempty" form:"dial_code" filter:"trim"`
	Mobile   string `json:"mobile,omitempty" form:"mobile" filter:"trim"`
	Email    string `json:"email,omitempty" form:"email" filter:"trim"`
	Code     string `json:"code,omitempty" form:"code" filter:"trim"`
	Captcha  string `json:"captcha,omitempty" form:"captcha" filter:"trim"`
	Nickname string `json:"nickname,omitempty" form:"nickname" filter:"trim"`
	Name     string `json:"name,omitempty" form:"name" filter:"trim"`
}

func (that *structUserRegister) Filter() *structUserRegister {
	that.Username = strings.TrimSpace(that.Username)
	that.Password = strings.TrimSpace(that.Password)
	that.DialCode = strings.TrimSpace(that.DialCode)
	that.Mobile = strings.TrimSpace(that.Mobile)
	that.Code = strings.TrimSpace(that.Code)
	that.Captcha = strings.TrimSpace(that.Captcha)
	that.Nickname = strings.TrimSpace(that.Nickname)
	that.Name = strings.TrimSpace(that.Name)
	return that
}

// Register 注册用户
func (that *structUserRegister) Register(controller mvvc.Controller) (userModel *ent.User, err error) {
	var ctx context.Context
	if that.Username != "" {
		if that.Captcha == "" {
			err = errors.New("username and captcha can't is empty")
			return
		} else if that.Password == "" {
			err = errors.New("username and password can't is empty")
			return
		}
	}
	if that.Mobile != "" {
		if that.DialCode == "" {
			that.DialCode = "86"
		}
		if that.DialCode == "86" && !utils.ValidateMobile(that.Mobile) {
			err = errors.New("inaccurate phone number format")
			return
		}
	}
	if that.Email != "" && !utils.ValidateEmail(that.Email) {
		err = errors.New("inaccurate email format")
		return
	}
	// 检查密码长度
	// 检查code是否有效
	if that.Code == "" {
		if that.Mobile != "" {
			err = errors.New("mobile and code can't is empty")
			return
		} else if that.Email != "" {
			err = errors.New("email and code can't is empty")
			return
		}
	} else {
		if that.Mobile != "" {

		} else if that.Email != "" {

		}
	}
	client := db.DefaultClient()
	// 检查用户名是否匹配
	if that.Username != "" {
		userModel, err = client.User.Query().
			Where(user.And(user.UsernameEQ(that.Username))).
			First(ctx)
		if !ent.IsNotFound(err) {
			return
		}
	}
	// 检查手机号是否匹配
	if that.Mobile != "" {
		userModel, err = client.User.Query().
			Where(user.And(user.MobileEQ(that.Mobile), user.DialCodeEQ(that.DialCode))).
			First(ctx)
		if !ent.IsNotFound(err) {
			return
		}
	}
	// 检查邮箱是否存在
	if that.Email != "" {
		userModel, err = client.User.Query().
			Where(user.And(user.EmailEQ(that.Email))).
			First(ctx)
		if !ent.IsNotFound(err) {
			return
		}
	}
	// 用户名存在，是否将软删用户还原
	err = db.WithTx(ctx, client, func(tx *ent.Tx) error {
		createUser := client.User.Create()
		userModel, err = createUser.
			SetUsername(that.Username).
			SetDialCode(that.DialCode).
			SetUsername(that.Mobile).
			SetEmail(that.Email).
			SetName(that.Name).
			SetNickname(that.Nickname).
			SetUsername(that.Username).
			SetStatus(StatusEnable).
			Save(ctx)
		if err != nil {
			userModel = nil
			log.Println("register fail", err)
			err = errors.New("register fail, please try again")
			return err
		}
		_, err = client.AppUser.Create().
			SetAppID(controller.App.ID).
			SetUser(userModel).
			SetNickname(userModel.Nickname).
			Save(ctx)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		userModel = nil
		return
	}
	return
}

func NewUserRegister() structUserRegister {
	c := structUserRegister{}
	return c
}