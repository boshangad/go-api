package smsService

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/boshangad/go-api/cccc/db"
	"github.com/boshangad/go-api/ent"
	"github.com/boshangad/go-api/ent/smslog"
	"log"
)

type codeStruct struct {
	Code string `json:"code,omitempty"`
}

func CheckCodeIsValid(dialCode, mobile, code string, typeId uint64) error {
	ctx := context.Background()
	smsLogModel, err := db.DefaultClient().SmsLog.Query().
		Where(smslog.And(smslog.MobileEQ(mobile), smslog.DialCodeEQ(dialCode), smslog.TypeIDEQ(typeId))).
		Order(ent.Desc(smslog.FieldID)).
		First(ctx)
	if err != nil {
		log.Println("查询数据失败，请重试", err)
		return errors.New("code invalid, please try again")
	}
	if smsLogModel.CheckCount > 3 {
		smsLogModel.Update()
		return errors.New("code invalid. please re-acquire")
	}
	err = db.DefaultClient().SmsLog.Update().AddCheckCount(1).Exec(ctx)
	if err != nil {
		log.Println("自增失败，请稍后重试", err)
		return err
	}
	data := codeStruct{}
	err = json.Unmarshal([]byte(smsLogModel.Content), &data)
	if err != nil {
		log.Println("短信内容转换json失败", smsLogModel.Content, err)
		return errors.New("code invalid, please try again")
	}
	if data.Code == "" {
		return errors.New("the code is abnormal, please get the verification code again")
	}
	if data.Code != code {
		return errors.New("code invalid, please try again")
	}
	// 已被消费，设置为已使用
	smsLogModel, err = smsLogModel.Update().SetStatus(1).Save(ctx)
	if err != nil {
		log.Println("设置为已更新失败", err)
		return errors.New("code invalid, please try again")
	}

	return nil
}