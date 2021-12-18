package gateways

import (
	"encoding/json"
	"net/url"
	"time"

	"github.com/boshangad/v1/app/helpers"
	"github.com/boshangad/v1/app/sms/interfaces"
)

type Aliyun struct {
	// 访问地址
	EndpointUrl string
	// 区域ID
	RegionId string
	// 参数格式
	Format string
	// 加密方法
	SignatureMethod string
	// 加密版本
	SignatureVersion string
	// 请求方法
	Action string
	// 版本
	Version string
}

func (*Aliyun) Name() string {
	return "aliyun"
}

func (that *Aliyun) Send(to interfaces.PhoneNumber, message interfaces.Message, config interfaces.Config) (interfaces.Result, error) {
	var (
		data        = message.GetData(that)
		urlParams   = url.Values{}
		phoneNumber = to.GetNumber()
	)
	if to.GetIDD() != "" {
		phoneNumber = to.GetUniversalNumber()
	}
	templateParam, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	urlParams.Set("RegionId", that.RegionId)
	urlParams.Set("AccessKeyId", config.GetString("accessKeyId"))
	urlParams.Set("Format", that.Format)
	urlParams.Set("SignatureMethod", that.SignatureMethod)
	urlParams.Set("SignatureVersion", that.SignatureVersion)
	urlParams.Set("SignatureNonce", helpers.RandStringBytesMaskImpr(8))
	urlParams.Set("Timestamp", time.Now().UTC().String())
	urlParams.Set("Action", that.Action)
	urlParams.Set("Version", that.Version)
	urlParams.Set("PhoneNumbers", phoneNumber)
	urlParams.Set("SignName", config.GetString("signName"))
	urlParams.Set("TemplateCode", message.GetTemplate(that))
	urlParams.Set("TemplateParam", string(templateParam))

	return nil, nil
}

// 实例化短信网关
func NewAliyunGateway(gateway interfaces.Gateway) *Aliyun {
	return &Aliyun{
		EndpointUrl:      "https://dysmsapi.aliyuncs.com",
		RegionId:         "cn-hangzhou",
		Action:           "SendSms",
		Format:           "JSON",
		Version:          "2017-05-25",
		SignatureMethod:  "HMAC-SHA1",
		SignatureVersion: "1.0",
	}
}
