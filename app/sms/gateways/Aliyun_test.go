package gateways

import (
	"net/url"
	"testing"

	"github.com/boshangad/v1/app/sms/config"
	"github.com/boshangad/v1/app/sms/core"
)

func TestAliyunSend(t *testing.T) {
	var (
		config = config.Gateway{
			// 应用键
			"accessKeyId": "",
			// 应用私钥
			"accessKeySecret": "",
			// 签名
			"signName": "",
		}
		data = url.Values{
			"code": []string{"123456"},
		}
		mobile  string
		aliyun  = NewAliyunGateway(config)
		message = core.NewMessage()
	)
	message.SetData(data).SetTemplate(aliyun, "")
	r := aliyun.Send(core.NewPhoneNumber(mobile), message, config)
	if !r.IsSuccess() {
		t.Error(r.Error())
	}
}
