package emailService

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/boshangad/go-api/core/config"
	"github.com/boshangad/go-api/core/db"
	"github.com/boshangad/go-api/ent"
	"github.com/boshangad/go-api/ent/emaillog"
	"strconv"
)

var aliyunGateways map[string]*aliyunStruct

type aliyunStruct struct {
	EmailInterface
	client *sdk.Client
	// 阿里云邮件的参数
	TagName string `json:"tag_name,omitempty"`
	// 使用管理控制台中配置的回信地址（状态必须是验证通过）。
	ReplyToAddress bool `json:"reply_to_address"`
	// 地址类型 0随机类型，1发信地址
	AddressType int `json:"address_type"`
	// 回信地址
	ReplyAddress string `json:"reply_address"`
	// 回信地址别称
	ReplyAddressAlias string `json:"reply_address_alias"`
	// 打卡数据追踪功能
	ClickTrace int64 `json:"click_trace"`
}

// Send 发送短信
func (that aliyunStruct) Send(emailData EmailStruct) (isSuccess bool, err error) {
	client := db.DefaultClient()
	var (
		dataStr []byte
		ctx = context.Background()
		emailLogModel *ent.EmailLog
	)
	dataStr, err = json.Marshal(emailData.Data)
	if err != nil {
		return
	}
	emailLogModel, err = client.EmailLog.Create().
		SetAppID(emailData.AppId).
		SetEmail(emailData.Email).
		SetScope(emailData.Scope).
		SetTypeID(emailData.TypeId).
		SetGateway(emailData.GateWay).
		SetIP(emailData.Ip).
		SetFromName(emailData.FromName).
		SetFromAddress(emailData.FromAddress).
		SetTitle(emailData.Title).
		SetContent(emailData.Content).
		SetData(string(dataStr)).
		SetStatus(StatusDraft).
		Save(ctx)
	if err != nil {
		return
	}
	request := requests.NewCommonRequest()
	request.Method = "POST"
	request.Scheme = "https" // https | http
	request.Domain = "dm.aliyuncs.com"
	request.Version = "2015-11-23"
	request.ApiName = "SingleSendMail"
	request.QueryParams["AccountName"] = emailData.FromAddress
	request.QueryParams["AddressType"] = string(rune(that.AddressType))
	request.QueryParams["ReplyToAddress"] = strconv.FormatBool(that.ReplyToAddress)
	request.QueryParams["TagName"] = that.TagName
	request.QueryParams["ToAddress"] = emailData.Email
	request.QueryParams["Subject"] = emailData.Title
	request.QueryParams["HtmlBody"] = emailData.Content
	request.QueryParams["TextBody"] = ""
	request.QueryParams["FromAlias"] = emailData.FromName
	request.QueryParams["ReplyAddress"] = that.ReplyAddress
	request.QueryParams["ReplyAddressAlias"] = that.ReplyAddressAlias
	request.QueryParams["ClickTrace"] = strconv.FormatInt(that.ClickTrace, 10)
	response, err := that.client.ProcessCommonRequest(request)
	if err != nil {
		return
	}
	_, err = emailLogModel.Update().
		SetStatus(StatusPublished).
		SetReturnMsg(response.GetHttpContentString()).
		Save(ctx)
	if err != nil {
		return
	}
	isSuccess = true
	return
}

func (aliyunStruct) CheckCode(email, code string, typeId uint64, appId uint64) (err error) {
	var (
		client = db.DefaultClient()
		ctx = context.Background()
		query = client.EmailLog.Query()
		emailLogModel *ent.EmailLog
		dataMap map[string]string
	)
	if appId > 0 {
		query.Where(emaillog.AppIDEQ(appId))
	}
	emailLogModel, err = query.
		Where(emaillog.And(emaillog.EmailEQ(email), emaillog.TypeIDEQ(typeId))).
		Where(emaillog.And(emaillog.StatusEQ(StatusPublished))).
		Order(ent.Desc(emaillog.FieldID)).
		First(ctx)
	if err != nil {
		return
	}
	if emailLogModel.Data == "" {
		err = errors.New("data does not exist verification code#1")
		return
	}
	err = json.Unmarshal([]byte(emailLogModel.Data), &dataMap)
	if err != nil {
		return
	}
	if emailLogModel.CheckCount > 6 {
		return errors.New("the number of verifications has been exceeded, please obtain the verification code again")
	}
	err = emailLogModel.Update().AddCheckCount(1).Exec(ctx)
	if err != nil {
		return
	}
	checkCode, ok := dataMap["code"]
	if !ok {
		return errors.New("data does not exist verification code#2")
	}
	if checkCode != code {
		return errors.New("code does not match, please try again later")
	}
	// 更新码状态为已使用
	_, _ = emailLogModel.Update().SetStatus(StatusUsed).Save(ctx)
	return
}

func NewAliyun(key string) *aliyunStruct {
	aliyunGateway, ok := aliyunGateways[key]
	if ok {
		return aliyunGateway
	}
	pushConfig, ok := config.Get().EmailPush.Gateways[key]
	if !ok {
		panic("")
	}
	client, err := sdk.NewClientWithAccessKey(pushConfig.RegionId, pushConfig.AccessKeyId, pushConfig.AccessKeySecret)
	if err != nil {
		panic(err)
	}
	way := aliyunStruct{}
	way.client = client
	// 加入变量
	aliyunGateways[key] = &way
	return &way
}