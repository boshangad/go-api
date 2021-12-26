package gateways

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"encoding/json"
	"encoding/xml"
	"io/ioutil"
	"net/http"
	"net/url"
	"sort"
	"strings"
	"time"

	"github.com/boshangad/v1/app/helpers"
	"github.com/boshangad/v1/app/sms/config"
	"github.com/boshangad/v1/app/sms/interfaces"
	"github.com/boshangad/v1/app/sms/results"
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

func (Aliyun) Name() string {
	return "aliyun"
}

func (that *Aliyun) Send(to interfaces.PhoneNumber, message interfaces.Message, config interfaces.Config) interfaces.Result {
	var (
		data        = message.GetData(that)
		urlParams   = url.Values{}
		phoneNumber = to.GetNumber()
		result      = results.Aliyun{
			CodeOrigin: "Failed",
		}
		templateData = make(map[string]string)
	)
	if to.GetIDD() != "" {
		phoneNumber = to.GetUniversalNumber()
	}
	for k := range data {
		templateData[k] = data.Get(k)
	}
	templateParam, err := json.Marshal(templateData)
	if err != nil {
		result.Message = err.Error()
		return result
	}
	urlParams.Set("RegionId", that.RegionId)
	urlParams.Set("AccessKeyId", config.GetString("accessKeyId"))
	urlParams.Set("Format", that.Format)
	urlParams.Set("SignatureMethod", that.SignatureMethod)
	urlParams.Set("SignatureVersion", that.SignatureVersion)
	urlParams.Set("SignatureNonce", helpers.RandStringBytesMaskImpr(8))
	urlParams.Set("Timestamp", time.Now().UTC().Format("2006-01-02T15:04:05Z"))
	urlParams.Set("Action", that.Action)
	urlParams.Set("Version", that.Version)
	urlParams.Set("PhoneNumbers", phoneNumber)
	urlParams.Set("SignName", config.GetString("signName"))
	urlParams.Set("TemplateCode", message.GetTemplate(that))
	urlParams.Set("TemplateParam", string(templateParam))
	signtureStr := that.generateSign(urlParams, config.GetString("accessKeySecret"), http.MethodGet)
	urlParams.Set("Signature", signtureStr)
	// 开始发生请求
	res, err := http.Get(that.EndpointUrl + "?" + urlParams.Encode())
	// http.Post("","", io.MultiReader(urlParams.Encode()))
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
	// 判断是否为json
	if strings.ToUpper(that.Format) == "JSON" {
		if err := json.Unmarshal(response, &result); err != nil {
			result.Message = "json failed:" + err.Error() + ". json value:" + string(response)
			return result
		}
	} else {
		if err := xml.Unmarshal(response, &result); err != nil {
			result.Message = "xml failed:" + err.Error() + ". xml value:" + string(response)
			return result
		}
	}

	return result
}

func (that Aliyun) generateSign(data url.Values, privateKey, requestMethod string) string {
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
		signtureStr += k + "=" + url.QueryEscape(data.Get(k)) + "&"
	}
	signtureStr = strings.TrimRight(signtureStr, "&")
	// urlParams RFC3986
	signtureStr = requestMethod + "&%2F&" + url.QueryEscape(signtureStr)
	// hmac-sha1
	key := []byte(privateKey + "&")
	mac := hmac.New(sha1.New, key)
	mac.Write([]byte(signtureStr))
	return base64.StdEncoding.EncodeToString(mac.Sum(nil))
}

// 实例化短信网关
func NewAliyunGateway(gateway config.Gateway) *Aliyun {
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
