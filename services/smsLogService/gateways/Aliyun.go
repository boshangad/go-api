package gateways

import (
	"net/http"

	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	dysmsapi "github.com/alibabacloud-go/dysmsapi-20170525/v2/client"
	"github.com/alibabacloud-go/tea/tea"
	"github.com/boshangad/v1/services/smsLogService/contact"
	"github.com/boshangad/v1/services/smsLogService/interfaces"
)

type aliyunGateway struct {
	*openapi.Config
}

func (aliyunGateway) Name() string {
	return "aliyun"
}

func (that *aliyunGateway) Send(to interfaces.PhoneNumber, message interfaces.Message) (result *contact.Result, err error) {
	// 访问的域名
	that.Endpoint = tea.String("dysmsapi.aliyuncs.com")
	client, _ := dysmsapi.NewClient(that.Config)
	request := &dysmsapi.AddShortUrlRequest{}
	response, err := client.AddShortUrl(request)
	if err != nil {
		return nil, result
	}
	result.Status = http.StatusOK
	result.Body = response.String()
	result.Err = *response.Body.Code
	result.Message = *response.Body.Message
	return result, nil
}

func NewAliyunGateway(gateway map[string]interface{}) *aliyunGateway {
	return &aliyunGateway{}
}
