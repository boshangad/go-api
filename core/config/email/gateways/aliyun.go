package gateways

import (
	"errors"
	"fmt"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/mitchellh/mapstructure"
	"log"
	"strconv"
	"strings"
)

type aliyunParams struct {
	RegionId     string `json:"region_id,omitempty" mapstructure:"region_id"`
	AccessKey    string `json:"access_key,omitempty" mapstructure:"access_key"`
	AccessSecret string `json:"access_secret,omitempty" mapstructure:"access_secret"`
	// 发信人
	FromName string `json:"from_name,omitempty" mapstructure:"from_name"`
	// 发信人昵称
	FromAddress string `json:"from_address,omitempty" mapstructure:"from_address"`
	// 阿里云邮件的参数
	TagName string `json:"tag_name,omitempty" mapstructure:"tag_name"`
	// 使用管理控制台中配置的回信地址（状态必须是验证通过）。
	ReplyToAddress bool `json:"reply_to_address,omitempty" mapstructure:"reply_to_address"`
	// 地址类型 0随机类型，1发信地址
	AddressType int `json:"address_type,omitempty" mapstructure:"address_type"`
	// 回信地址
	ReplyAddress string `json:"reply_address,omitempty" mapstructure:"reply_address"`
	// 回信地址别称
	ReplyAddressAlias string `json:"reply_address_alias,omitempty" mapstructure:"reply_address_alias"`
	// 打卡数据追踪功能
	ClickTrace int64  `json:"click_trace,omitempty" mapstructure:"click_trace"`
}

type AliyunConfig struct {
	client *sdk.Client
	aliyunParams
}

func (AliyunConfig) Name() string {
	return "aliyun"
}

// Send 发送短信
func (that *AliyunConfig) Send(data Data) (returnMsg string, err error) {
	request := requests.NewCommonRequest()
	request.Method = "POST"
	request.Scheme = "https" // https | http
	request.Domain = "dm.aliyuncs.com"
	request.Version = "2015-11-23"
	request.ApiName = "SingleSendMail"
	request.QueryParams["ToAddress"] = data.Email
	request.QueryParams["Subject"] = data.Title
	content := data.Content
	if data.Data != nil {
		for key, value := range data.Data {
			switch value.(type) {
			case string:
				content = strings.ReplaceAll(content, "{" + key + "}", value.(string))
			default:
				content = strings.ReplaceAll(content, "{" + key + "}", fmt.Sprintf("%s", value))
			}
		}
	}
	request.QueryParams["HtmlBody"] = content
	request.QueryParams["TextBody"] = ""
	if data.FromName != "" {
		request.QueryParams["FromAlias"] = data.FromName
	} else {
		request.QueryParams["FromAlias"] = that.FromName
	}
	if data.FromAddress != "" {
		request.QueryParams["AccountName"] = data.FromAddress
	} else {
		request.QueryParams["AccountName"] = that.FromAddress
	}
	request.QueryParams["AddressType"] = strconv.Itoa(that.AddressType)
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
	if !response.IsSuccess() {
		err = errors.New(response.String())
		return
	}
	returnMsg = response.GetHttpContentString()
	return
}

// NewAliyunGateway 实例化邮件推送服务网关
func NewAliyunGateway(params map[string]interface{}) *AliyunConfig {
	var config aliyunParams
	err := mapstructure.Decode(params, &config)
	if err != nil {
		log.Panicln(err)
	}
	client, err := sdk.NewClientWithAccessKey(config.RegionId, config.AccessKey, config.AccessSecret)
	if err != nil {
		log.Panicln(err)
	}
	aliyunConfig := AliyunConfig{
		aliyunParams: config,
		client: client,
	}
	return &aliyunConfig
}