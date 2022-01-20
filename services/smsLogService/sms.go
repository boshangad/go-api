package smsLogService

import (
	"context"
	"errors"
	"net/url"
	"strconv"
	"time"

	"github.com/boshangad/v1/app/config"
	"github.com/boshangad/v1/app/controller"
	"github.com/boshangad/v1/app/sms"
	"github.com/boshangad/v1/ent"
	"github.com/boshangad/v1/ent/smslog"
	"github.com/boshangad/v1/global"
	"go.uber.org/zap"
)

var smsGateway *sms.SmsGateway = NewGateway(global.Config)

func NewGateway(c *config.Config) *sms.SmsGateway {
	smsGateway := sms.NewSmsGateway(c.Sms)
	c.AddObserver("smsGateway", func(c *config.Config) {
		smsGateway.Reload(c.Sms)
	})
	return smsGateway
}

// 验证码
type ValidateCode struct {
	// 手机号
	Mobile string `json:"mobile" binding:"required"`
	// 国际区号
	DialCode string `json:"dial_code,omitempty" binding:""`
	// 可用范围
	Scope string `json:"scope,omitempty" binding:"required"`
}

// 发出验证码
func (that ValidateCode) SendRandomCode(c *controller.Context) error {
	var (
		ctx          = context.Background()
		gatewayNames = []string{}
		data         = url.Values{}
	)
	if that.DialCode == "" {
		that.DialCode = "86"
	}
	smsLog, err := global.Db.SmsLog.Query().
		Where(
			smslog.MobileEQ(that.Mobile),
			smslog.ScopeEQ(that.Scope),
		).
		Order(ent.Desc(smslog.FieldID)).
		First(ctx)
	if err != nil {
		if !ent.IsNotFound(err) {
			return err
		}
		// 数据未被找到
	} else {
		// 数据被找到
	}
	// 60秒内发送过
	maxCreateTime := smsLog.CreateTime + 60
	nowTime := time.Now().Unix()
	if maxCreateTime < nowTime {
		return errors.New("the text message has been sent, please try again in " + strconv.Itoa(int(maxCreateTime-smsLog.CreateTime)) + " seconds")
	}
	// 使用次数小于3次
	if smsLog.CheckCount < (uint8(3) + 1) {

	}
	// 频次验证 【30分钟5次 3小时7次， 24小时10次】
	frequencyRange := []struct {
		IntervalTime int64
		Count        int
	}{
		{IntervalTime: 30 * 60, Count: 5},
		{IntervalTime: 3 * 3600, Count: 7},
		{IntervalTime: 12 * 3600, Count: 10},
	}
	for _, s := range frequencyRange {
		beforeTime := nowTime - s.IntervalTime
		n, err := global.Db.SmsLog.Query().Where(smslog.MobileEQ(that.Mobile), smslog.DialCodeEQ(that.DialCode)).Where(smslog.CreateTimeGTE(beforeTime)).Count(ctx)
		if err != nil {
			global.Log.Warn("sms limit query failed, request mobile is restricted.", zap.Error(err))
			return err
		}
		if n >= s.Count {
			return errors.New("the mobile phone number has triggered the system current limit")
		}
	}
	// 同IP发送频次限制
	frequencyRange = []struct {
		IntervalTime int64
		Count        int
	}{
		{IntervalTime: 600, Count: 20},
		{IntervalTime: 3600, Count: 50},
		{IntervalTime: 12 * 3600, Count: 100},
	}
	for _, s := range frequencyRange {
		beforeTime := nowTime - s.IntervalTime
		n, err := global.Db.SmsLog.Query().Where(smslog.IPEQ(c.ClientIP()), smslog.CreateTimeGTE(beforeTime)).Count(ctx)
		if err != nil {
			global.Log.Warn("sms limit query failed, request IP is restricted.", zap.Error(err))
			return err
		}
		if n >= s.Count {
			return errors.New("request limited, request frequency is too high")
		}
	}
	// 检测是否连续发送短信验证失败3次
	_, err = smsGateway.Send(
		that.Mobile,
		data,
		gatewayNames,
	)
	if err != nil {
		return err
	}
	// gateway := results.Get("")
	return nil
}
