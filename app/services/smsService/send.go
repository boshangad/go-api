package smsService

import (
	"fmt"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
)

var aliyunClient *sdk.Client
func init()  {
	client, err := sdk.NewClientWithAccessKey("REGION_ID", "ACCESS_KEY_ID", "ACCESS_KEY_SECRET")
	if err != nil {
		panic(err)
	}
	aliyunClient = client
}

func Send(dialCode, mobile string, typeId uint64) error {
	return SendByAliyun(dialCode, mobile, typeId)
}

func SendByAliyun(dialCode, mobile string, typeId uint64) error {
	request := requests.NewCommonRequest()
	request.Method = "POST"
	request.Scheme = "https" // https | http
	request.Domain = "dysmsapi.aliyuncs.com"
	request.Version = "2018-05-01"
	request.ApiName = "SendSms"
	request.QueryParams["SignName"] = "签名"
	request.QueryParams["PhoneNumbers"] = dialCode+mobile
	request.QueryParams["ContentCode"] = "模板ID"
	request.QueryParams["ContentParam"] = "值"
	request.QueryParams["ExternalId"] = "流水扩展字段"

	response, err := aliyunClient.ProcessCommonRequest(request)
	if err != nil {
		panic(err)
	}
	fmt.Print(response.GetHttpContentString())
	return nil
}