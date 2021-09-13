package gateways

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"log"
	"strconv"
)

type aliyunConfig struct {
	client *sdk.Client
	IsAliyun bool `json:"is_aliyun,omitempty"`
	Gateway string `json:"gateway,omitempty"`
	RegionId string `json:"region_id,omitempty"`
	AccessKey string `json:"access_key,omitempty"`
	AccessSecret string `json:"access_secret,omitempty"`
	// 发信人
	FromName string `json:"from_name,omitempty"`
	// 发信人昵称
	FromAddress string `json:"from_address,omitempty"`
	// 阿里云邮件的参数
	TagName string `json:"tag_name,omitempty"`
	// 使用管理控制台中配置的回信地址（状态必须是验证通过）。
	ReplyToAddress bool `json:"reply_to_address,omitempty"`
	// 地址类型 0随机类型，1发信地址
	AddressType int `json:"address_type,omitempty"`
	// 回信地址
	ReplyAddress string `json:"reply_address,omitempty"`
	// 回信地址别称
	ReplyAddressAlias string `json:"reply_address_alias,omitempty"`
	// 打卡数据追踪功能
	ClickTrace int64 `json:"click_trace,omitempty"`
	ReturnMsg string `json:"return_msg,omitempty"`
}

// Send 发送短信
func (that *aliyunConfig) Send(data Data) (isSuccess bool, err error) {
	request := requests.NewCommonRequest()
	request.Method = "POST"
	request.Scheme = "https" // https | http
	request.Domain = "dm.aliyuncs.com"
	request.Version = "2015-11-23"
	request.ApiName = "SingleSendMail"
	request.QueryParams["ToAddress"] = data.Email
	request.QueryParams["Subject"] = data.Title
	request.QueryParams["HtmlBody"] = data.Content
	request.QueryParams["TextBody"] = ""
	if data.FromName != "" {
		request.QueryParams["FromAlias"] = data.FromName
	} else {
		request.QueryParams["FromAlias"] = that.FromName
	}
	if data.FromAddress != "" {
		request.QueryParams["FromAlias"] = data.FromAddress
	} else {
		request.QueryParams["FromAlias"] = that.FromAddress
	}
	request.QueryParams["AccountName"] = data.FromAddress
	request.QueryParams["AddressType"] = string(rune(that.AddressType))
	request.QueryParams["ReplyToAddress"] = strconv.FormatBool(that.ReplyToAddress)
	request.QueryParams["TagName"] = that.TagName
	request.QueryParams["ReplyAddress"] = that.ReplyAddress
	request.QueryParams["ReplyAddressAlias"] = that.ReplyAddressAlias
	request.QueryParams["ClickTrace"] = strconv.FormatInt(that.ClickTrace, 10)
	response, err := that.client.ProcessCommonRequest(request)
	if err != nil {
		log.Println("aliyun email push fail", err)
		return
	}
	that.ReturnMsg = response.GetHttpContentString()
	isSuccess = true
	return
}

// NewAliyunGateway 实例化邮件推送服务网关
func NewAliyunGateway(config aliyunConfig) *aliyunConfig {
	client, err := sdk.NewClientWithAccessKey(config.RegionId, config.AccessKey, config.AccessSecret)
	if err != nil {
		panic(err)
	}
	config.Gateway = "aliyun"
	config.client = client
	return &config
}