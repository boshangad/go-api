package gateways

import (
	"net/url"
	"testing"

	"github.com/boshangad/v1/app/sms/config"
	"github.com/boshangad/v1/app/sms/core"
)

func TestUcloudSend(t *testing.T) {
	var (
		config = config.Gateway{
			// 应用键
			"privateKey": "",
			// 应用私钥
			"publicKey": "",
			// 签名
			"signName": "",
			// 项目成员
			"ProjectId": "",
		}
		data = url.Values{
			"0": []string{"123456"},
		}
		mobile  string = ""
		ucloud         = NewUcloudGateway(config)
		message        = core.NewMessage()
	)
	message.SetData(data).SetTemplate(ucloud, "")
	r := ucloud.Send(core.NewPhoneNumber(mobile), message, config)
	if !r.IsSuccess() {
		t.Error(r.Error())
	}
}
