package emailService

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"github.com/boshangad/go-api/ent"
	"github.com/boshangad/go-api/ent/emaillog"
	"github.com/boshangad/go-api/global"
	"github.com/boshangad/go-api/global/db"
	"github.com/boshangad/go-api/services/emailService/gateways"
	"github.com/mitchellh/mapstructure"
	"github.com/mojocn/base64Captcha"
	"html/template"
)

type SendData struct {
	Email   string `json:"email,omitempty" form:"email"`
	Scope   string `json:"scope,omitempty" form:"scope"`
	Captcha string `json:"captcha,omitempty" form:"captcha"`
}

type Data struct {
	AppId uint64
	Scope string
	Ip string
	TemplateId string
	Subject string
	Data map[string]interface{}
	Content string
}

var GatewayClients = make(map[string]gateways.EmailSend)

// Send 发送消息
func Send(gateway, email string, data Data) (err error) {
	if data.AppId < 1 || data.Scope == "" {
		return errors.New("missing parameters #appId and #scope")
	}
	if _, ok := ScopeTypeList[data.Scope]; !ok {
		return errors.New("the scope value is not registered")
	}
	// 防止并发多次执行
	if _, ok := GatewayClients[gateway]; !ok {
		_, err, _ = global.G_Concurrency_Control.Do("emailGateway:" + gateway, func() (interface{}, error) {
			v, ok1 := global.G_CONFIG.Email.Gateways[gateway]
			if !ok1 {
				return nil, errors.New("configuration `" + gateway + "` not found")
			}
			if _, ok2 := v["gateway"]; !ok2 {
				return nil, errors.New("the configuration `" + gateway + "` did not find the `gateway` field")
			}
			// 检查网关
			switch v["gateway"].(string) {
			case "local":
				d := gateways.LocalEmail{}
				err = mapstructure.Decode(v, d)
				if err != nil {
					return nil, err
				}
				GatewayClients[gateway] = d
			case "aliyun":
				d := gateways.AliyunEmail{}
				err = mapstructure.Decode(v, &d)
				if err != nil {
					return nil, err
				}
				d.NewClient()
				GatewayClients[gateway] = d
			default:
				return nil, errors.New("configure `" + gateway + "` can not find the owning gateway")
			}
			return gateway, nil
		})
		if err != nil {
			return err
		}
	}
	// 开始执行发送邮件
	var (
		pushClient = GatewayClients[gateway]
		ctx = context.Background()
		dbClient = db.DefaultClient()
		typeId = ScopeTypeList[data.Scope]
		subject = data.Subject
		content = data.Content
		dataBytes []byte
		subjectBytes bytes.Buffer
		contentBytes bytes.Buffer
	)
	dataBytes, err = json.Marshal(data)
	if err != nil {
		return err
	}
	// 获取要发送的信息,先验证是否是采用模板发送
	content = data.Content
	subject = data.Subject
	if data.TemplateId != "" {
		// 获取相关的数据
	}
	err = template.Must(template.New("emailSubjectSend:" + data.Scope).Parse(subject)).
		Execute(&subjectBytes, data)
	if err != nil {
		global.G_LOG.Error(err.Error())
		return err
	}
	subject = subjectBytes.String()
	err = template.Must(template.New("emailContentSend:" + data.Scope).Parse(content)).
		Execute(&contentBytes, data)
	if err != nil {
		global.G_LOG.Error(err.Error())
		return err
	}
	content = contentBytes.String()

	// 检测是否存在验证码
	emailLogModel, err := dbClient.EmailLog.Query().Where(emaillog.And(
		emaillog.AppIDEQ(data.AppId),
		emaillog.EmailEQ(email),
		emaillog.ScopeEQ(data.Scope)),
		emaillog.StatusIn(StatusToSend),
	).First(ctx)
	if err != nil {
		if !ent.IsNotFound(err) {
			global.G_LOG.Error(err.Error())
			return
		}
		// 表示未找到相关数据
		emailLogModel, err = dbClient.EmailLog.Create().
			SetAppID(data.AppId).
			SetEmail(email).
			SetScope(data.Scope).
			SetTypeID(uint64(typeId)).
			SetGateway(gateway).
			SetIP(data.Ip).
			SetTitle(subject).
			SetContent(content).
			SetData(string(dataBytes)).
			SetCheckCount(0).
			SetStatus(StatusToSend).
			Save(ctx)
		if err != nil {
			return
		}
	}

	// 发送信息
	returnMsg, err := pushClient.Send(email, subject, content)
	if err != nil {
		global.G_LOG.Warn("push email fail:" + err.Error())
		return
	}
	_, err = emailLogModel.Update().
		SetTitle(subject).
		SetContent(content).
		SetData(string(dataBytes)).
		SetStatus(StatusSendSuccess).
		SetReturnMsg(returnMsg).
		Save(ctx)
	if err != nil {
		global.G_LOG.Error("update emailLog fail:" + err.Error())
		return
	}
	return
}

// SendCode 发送短信验证码
func SendCode(gateway, email, ip, scope string, appId uint64) error {
	var code = base64Captcha.RandText(6, base64Captcha.TxtNumbers)
	return Send(gateway, email, Data{
		AppId: appId,
		Scope: scope,
		Ip: ip,
		Subject: "",
		Content: "",
		Data: map[string]interface{}{"code": code},
	})
}