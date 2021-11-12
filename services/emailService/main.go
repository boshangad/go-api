package emailService

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"html/template"
	"strings"

	"github.com/boshangad/go-api/ent"
	"github.com/boshangad/go-api/ent/emaillog"
	"github.com/boshangad/go-api/global"
	"github.com/boshangad/go-api/global/db"
	"github.com/boshangad/go-api/services/emailService/gateways"
	"github.com/mojocn/base64Captcha"
)

// 网关客户端
var GatewayClients = make(map[string]Email)

type Email interface {
	// 推送
	Send(email, subject, content string, cc, bcc []string) (err error)
	MultiSend(email []string, subject, content string, cc, bcc []string) (isSuccess bool, errors map[string]error)
}

type SendData struct {
	Email   string `json:"email,omitempty" form:"email"`
	Scope   string `json:"scope,omitempty" form:"scope"`
	Captcha string `json:"captcha,omitempty" form:"captcha"`
}

type Data struct {
	AppId      uint64
	Scope      string
	Ip         string
	TemplateId string
	Subject    string
	Data       map[string]interface{}
	Content    string
}

// 实例化相关邮件推送服务
func NewGateWay(name string) Email {
	var (
		emailGateways map[string]map[string]interface{}
		typeStr       string
	)
	emailGateways = global.G_CONFIG.Email.Gateways
	if emailGateway, ok := emailGateways[name]; ok {
		typeStr = strings.ToLower(strings.TrimSpace(emailGateway["type"].(string)))
		if typeStr != "" {
			switch typeStr {
			case "aliyun":
			case "tencent":
			case "local":
				fallthrough
			default:
				return gateways.NewLocalGateway(emailGateway)
			}
		}
		panic("Mail configuration type " + typeStr + " is not supported, currently only supports local, aliyun")
	}
	panic("Undefined mail configuration: " + name)
}

// 线程安全的实例网关
func NewGatewayWithConcurrencyControl(name string) Email {
	mailClient, err, _ := global.G_Concurrency_Control.Do("emailGateway:"+name, func() (interface{}, error) {
		var (
			err error = nil
			ok  bool  = false
		)
		defer func() {
			if e := recover(); e != nil {
				err, ok = e.(error)
				if !ok {
					err = errors.New("new gateway is failed")
				}
			}
		}()
		return NewGateWay(name), err
	})
	if err != nil {
		global.G_LOG.Error(err.Error())
		return nil
	}
	g, ok := mailClient.(Email)
	if !ok {
		global.G_LOG.Error("Unknown error, email gateway fail")
		return nil
	}
	return g
}

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
		GatewayClients[gateway] = NewGatewayWithConcurrencyControl(gateway)
	}
	// 开始执行发送邮件
	var (
		pushClient   = GatewayClients[gateway]
		ctx          = context.Background()
		dbClient     = db.DefaultClient()
		typeId       = ScopeTypeList[data.Scope]
		subject      = data.Subject
		content      = data.Content
		dataBytes    []byte
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
	err = template.Must(template.New("emailSubjectSend:"+data.Scope).Parse(subject)).
		Execute(&subjectBytes, data)
	if err != nil {
		global.G_LOG.Error(err.Error())
		return err
	}
	subject = subjectBytes.String()
	err = template.Must(template.New("emailContentSend:"+data.Scope).Parse(content)).
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
	err = pushClient.Send(email, subject, content, []string{}, []string{})
	if err != nil {
		global.G_LOG.Warn("push email fail:" + err.Error())
		return
	}
	_, err = emailLogModel.Update().
		SetTitle(subject).
		SetContent(content).
		SetData(string(dataBytes)).
		SetStatus(StatusSendSuccess).
		SetReturnMsg("").
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
		AppId:   appId,
		Scope:   scope,
		Ip:      ip,
		Subject: "",
		Content: "",
		Data:    map[string]interface{}{"code": code},
	})
}
