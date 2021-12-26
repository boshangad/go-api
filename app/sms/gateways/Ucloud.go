package gateways

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"sort"

	"github.com/boshangad/v1/app/helpers"
	"github.com/boshangad/v1/app/sms/config"
	"github.com/boshangad/v1/app/sms/interfaces"
	"github.com/boshangad/v1/app/sms/results"
)

// ucloud请求网关
type Ucloud struct {
	// 请求路径
	requestHost string
	// 请求方法
	requestAction string
}

func (Ucloud) Name() string {
	return "ucloud"
}

// Send 发送短信
func (that Ucloud) Send(to interfaces.PhoneNumber, message interfaces.Message, config interfaces.Config) interfaces.Result {
	var (
		params = that.buildParams(to, message, config)
		result = results.Ucloud{
			RetCode: -1,
			Action:  that.requestAction,
		}
	)
	url, _ := url.Parse(that.requestHost)
	url.RawQuery = params.Encode()
	res, err := http.Get(url.String())
	if err != nil {
		result.Message = err.Error()
		return result
	}
	defer res.Body.Close()
	response, err := ioutil.ReadAll(res.Body)
	if err != nil {
		result.Message = err.Error()
		return result
	}
	// 打印
	err = json.Unmarshal(response, &result)
	if err != nil {
		result.Message = err.Error()
		return result
	}
	return result
}

// 生产请求参数
func (that Ucloud) buildParams(to interfaces.PhoneNumber, message interfaces.Message, config interfaces.Config) url.Values {
	var (
		data   = message.GetData(&that)
		params = url.Values{}
	)
	params.Set("Action", that.requestAction)
	params.Set("SigContent", config.GetString("signName"))
	params.Set("TemplateId", message.GetTemplate(&that))
	params.Set("PublicKey", config.GetString("publicKey"))

	// 用户录入的参数
	for k := range data {
		params.Set("TemplateParams."+k, data.Get(k))
	}
	// 用户通知的手机号
	params.Set("PhoneNumbers.0", to.GetNumber())
	// 项目ID。不填写为默认项目，子帐号必须填写。
	if config.GetString("projectId") != "" {
		params.Set("ProjectId", config.GetString("projectId"))
	}
	// 短信扩展码，格式为阿拉伯数字串，默认不开通，如需开通请联系 UCloud技术支持
	if config.GetString("extendCode") != "" {
		params.Set("ExtendCode", config.GetString("extendCode"))
	}
	// 自定义的业务标识ID，字符串（ 长度不能超过32 位），不支持 单引号、表情包符号等特殊字符
	if config.GetString("userId") != "" {
		params.Set("UserId", config.GetString("userId"))
	}
	params.Set("Signature", that.signtureParams(params, config.GetString("privateKey")))
	return params
}

// 计算签名
func (config Ucloud) signtureParams(data url.Values, privateKey string) string {
	var (
		keys        = make([]string, len(data))
		k           string
		i                  = 0
		signtureStr string = ""
	)
	for k = range data {
		keys[i] = k
		i++
	}
	sort.Strings(keys)
	for _, k = range keys {
		signtureStr += k + data.Get(k)
	}
	signtureStr += privateKey
	return helpers.Sha1(signtureStr)
}

// 实例化短信
func NewUcloudGateway(gateway config.Gateway) *Ucloud {
	return &Ucloud{
		requestHost:   "https://api.ucloud.cn",
		requestAction: "SendUSMSMessage",
	}
}
