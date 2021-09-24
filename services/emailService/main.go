package emailService

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/utils"
	"github.com/boshangad/go-api/cccc/config"
	emailConfig "github.com/boshangad/go-api/cccc/config/email"
	"github.com/boshangad/go-api/cccc/config/email/gateways"
	"github.com/boshangad/go-api/cccc/db"
	"github.com/boshangad/go-api/ent"
	"github.com/boshangad/go-api/ent/emaillog"
	"log"
)

// DefaultPushClient 默认推送客户端
func DefaultPushClient() gateways.ConfigInterface {
	emailPush := config.Get().EmailPush
	if emailPush == nil {
		log.Println("the default push mail gateway is not configured")
		return nil
	}
	return PushClient(emailPush.Default)
}

// PushClient 推送客户端
func PushClient(key string) gateways.ConfigInterface {
	emailPush := config.Get().EmailPush
	if emailPush != nil {
		if data, ok := emailPush.Clients[key]; ok {
			return data
		}
	}
	log.Panicln("email push client was not found", "gateway is", key)
	return nil
}

// SendCode 发送验证码
func SendCode(appId uint64, scope string, email string, ip string) (err error) {
	typeId := emailConfig.GetTypeIdWithScope(scope)
	if typeId < 1 {
		err = errors.New("the range of use is not accurate, please select the correct range of use")
		return
	}
	var (
		ctx = context.Background()
		data = map[string]interface{}{}
		dbClient = db.DefaultClient()
	)
	// 检测是否存在验证码
	emailLogModel, err := dbClient.EmailLog.Query().Where(emaillog.And(
		emaillog.AppIDEQ(appId),
		emaillog.EmailEQ(email),
		emaillog.ScopeEQ(scope)),
		emaillog.StatusIn(emailConfig.StatusDraft, emailConfig.StatusPublished),
	).First(ctx)
	if err != nil && !ent.IsNotFound(err) {
		log.Println("email send fail", err)
		return
	}
	pushClient := DefaultPushClient()
	if emailLogModel == nil {
		data = map[string]interface{}{
			"code": utils.RandStringBytes(6),
		}
		dataStr, _ := json.Marshal(data)
		/*fset, err := os.Open("./template_code.html")
		if err != nil {
			return err
		}
		content, err := ioutil.ReadAll(fset)
		if err != nil {
			return err
		}*/
		emailLogModel, err = dbClient.EmailLog.Create().
			SetAppID(appId).
			SetEmail(email).
			SetScope(scope).
			SetTypeID(typeId).
			SetGateway(pushClient.Name()).
			SetIP(ip).
			SetTitle("验证码登录").
			SetContent(string(getCodeTemplate())).
			SetData(string(dataStr)).
			SetCheckCount(0).
			SetStatus(emailConfig.StatusDraft).
			Save(ctx)
		if err != nil {
			return
		}
	} else {
		_ = json.Unmarshal([]byte(emailLogModel.Data), &data)
	}
	returnMsg, err := pushClient.Send(gateways.Data{
		Email: email,
		Title: emailLogModel.Title,
		Content: emailLogModel.Content,
		Data: data,
	})
	if err != nil {
		log.Println("push email fail", err)
		return
	}
	_, _ = emailLogModel.Update().
		SetStatus(emailConfig.StatusPublished).
		SetReturnMsg(returnMsg).
		Save(ctx)
	return
}