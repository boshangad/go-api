package gateways

import (
	"errors"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/boshangad/go-api/global"
	"strconv"
)

type AliyunEmail struct {
	client *sdk.Client
	RegionId string
	AccessKeyId string
	AccessKeySecret string
	StsToken string
	BaseEmail
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

func (that *AliyunEmail) SetRegionId(regionId string) *AliyunEmail {
	that.RegionId = regionId
	return that
}

func (that *AliyunEmail) SetAccessKeyId(accessKeyId string) *AliyunEmail {
	that.AccessKeyId = accessKeyId
	return that
}

func (that *AliyunEmail) SetAccessKeySecret(accessKeySecret string) *AliyunEmail {
	that.AccessKeySecret = accessKeySecret
	return that
}

func (that *AliyunEmail) SetStsToken(stsToken string) *AliyunEmail {
	that.StsToken = stsToken
	return that
}

func (that *AliyunEmail) SetTagName(tagName string) *AliyunEmail {
	that.TagName = tagName
	return that
}

func (that *AliyunEmail) SetReplyToAddress(replyToAddress bool) *AliyunEmail {
	that.ReplyToAddress = replyToAddress
	return that
}

func (that *AliyunEmail) SetAddressType(addressType int) *AliyunEmail {
	that.AddressType = addressType
	return that
}

func (that *AliyunEmail) SetReplyAddress(replyAddress string) *AliyunEmail {
	that.ReplyAddress = replyAddress
	return that
}

func (that *AliyunEmail) SetReplyAddressAlias(replyAddressAlias string) *AliyunEmail {
	that.ReplyAddressAlias = replyAddressAlias
	return that
}

func (that *AliyunEmail) SetClickTrace(clickTrace int64) *AliyunEmail {
	that.ClickTrace = clickTrace
	return that
}

func (that *AliyunEmail) SetClient(client *sdk.Client) *AliyunEmail {
	that.client = client
	return that
}

func (that AliyunEmail) GetClient() *sdk.Client {
	return that.client
}

func (that AliyunEmail) Send(to, subject, body string) (string, error) {
	request := requests.NewCommonRequest()
	request.Method = "POST"
	request.Scheme = "https" // https | http
	request.Domain = "dm.aliyuncs.com"
	request.Version = "2015-11-23"
	request.ApiName = "SingleSendMail"
	request.QueryParams["FromAlias"] = that.FromAlias
	request.QueryParams["AccountName"] = that.FromAddress
	request.QueryParams["AddressType"] = string(rune(that.AddressType))
	request.QueryParams["ReplyToAddress"] = strconv.FormatBool(that.ReplyToAddress)
	request.QueryParams["TagName"] = that.TagName
	request.QueryParams["ToAddress"] = to
	request.QueryParams["Subject"] = subject
	request.QueryParams["HtmlBody"] = body
	request.QueryParams["TextBody"] = ""
	request.QueryParams["ReplyAddress"] = that.ReplyAddress
	request.QueryParams["ReplyAddressAlias"] = that.ReplyAddressAlias
	request.QueryParams["ClickTrace"] = strconv.FormatInt(that.ClickTrace, 10)
	response, err := that.client.ProcessCommonRequest(request)
	if err != nil {
		return "", err
	}
	if !response.IsSuccess() {
		return "", errors.New(response.String())
	}
	return response.String(), nil
}

func (that *AliyunEmail) NewClient() *AliyunEmail {
	var (
		c   *sdk.Client
		err error
	)
	if "" == that.StsToken {
		c, err = sdk.NewClientWithAccessKey(that.RegionId, that.AccessKeyId, that.AccessKeySecret)
	} else {
		c, err = sdk.NewClientWithStsToken(that.RegionId, that.AccessKeyId, that.AccessKeySecret, that.StsToken)
	}
	if err != nil {
		global.G_LOG.Panic(err.Error())
	}
	that.client = c
	return that
}

func NewAliyun(regionId, accessKeyId, accessKeySecret, stsToken string) *AliyunEmail {
	data := AliyunEmail{
		RegionId:        regionId,
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
		StsToken:        stsToken,
	}
	return data.NewClient()
}