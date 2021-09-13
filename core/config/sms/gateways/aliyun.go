package gateways

import (
	"encoding/json"
	"errors"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"log"
)

type aliyun struct {
	Client *sdk.Client
	Config
}

func (that aliyun) Send(data Data, config TemplateConfig) (result string, err error) {
	contentParams, err := json.Marshal(data.Content)
	if err != nil {

	}
	request := requests.NewCommonRequest()
	request.Method = "POST"
	request.Scheme = "https"
	request.Domain = "dysmsapi.aliyuncs.com"
	request.Version = "2018-05-01"
	request.ApiName = "SendSms"
	request.QueryParams["ContentCode"] = config.TemplateId
	request.QueryParams["PhoneNumbers"] = data.DialCode + data.Mobile
	request.QueryParams["ContentParam"] = string(contentParams)
	request.QueryParams["SignName"] = that.SignName
	response, err := that.Client.ProcessCommonRequest(request)
	if err != nil {
		return
	}
	if !response.IsSuccess() {
		err = errors.New("alibaba Cloud gateway failed to send SMS")
		return
	}
	result = response.String()
	return
}

// NewAliyunGateway 实例化阿里云网关
func NewAliyunGateway(config Config) *aliyun {
	client, err := sdk.NewClientWithAccessKey(config.RegionId, config.AccessKey, config.AccessSecret)
	if err != nil {
		log.Panicln("Instance client failed", err)
	}
	return &aliyun{
		Client: client,
		Config: config,
	}
}