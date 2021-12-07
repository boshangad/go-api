package smsLogService

import (
	"strings"

	"github.com/boshangad/v1/app/config"
	"github.com/boshangad/v1/services/smsLogService/contact"
	sgateways "github.com/boshangad/v1/services/smsLogService/gateways"
	"github.com/boshangad/v1/services/smsLogService/interfaces"
)

// 短信推送
type SmsGateway struct {
	// 已经被初始化过的网关
	gateways map[string]interfaces.Gateway
	// 推送人
	messager *Messager
	// 配置项
	Config *config.Sms
}

// 网关操作人
func (that *SmsGateway) GetMessager() *Messager {
	if that.messager == nil {
		that.messager = NewMessager()
	}
	return that.messager
}

// 短信网关
func (that *SmsGateway) Gateway(name string) interfaces.Gateway {
	if _, ok := that.gateways[name]; !ok {
		if gc, ok := that.Config.Gateways[name]; !ok {
			if cType, ok1 := gc["gateway"]; ok1 {
				if cType2, ok2 := cType.(string); ok2 {
					switch strings.ToLower(strings.TrimSpace(cType2)) {
					case "aliyun":
						that.gateways[name] = sgateways.NewAliyunGateway(gc)
					case "aliyunrest":
					case "avatardata":
					case "baidu":
					case "chuanglan":
					case "huawei":
					case "huaxin":
					case "huyi":
					case "juhe":
					case "kingtto":
					case "luosimao":
					case "moduyun":
					case "qcloud":
					case "qiniu":
					case "rongcloud":
					case "rongheyun":
					case "sendcloud":
					case "smsbao":
					case "submail":
					case "tianyiwuxian":
					case "tiniyo":
					case "twilio":
					case "ucloud":
					case "ue35":
					case "yunpian":
					case "yuntongyun":
					case "yunxin":
					case "yunzhixun":
					case "zzyun":
					default:
					}
				}
			}
		}
	}
	return that.gateways[name]
}

// 网关
func (that *SmsGateway) GetGateways(gs []string) map[string]interfaces.Gateway {
	var gateways = make(map[string]interfaces.Gateway)
	if len(gs) > 0 {
		for _, key := range gs {
			gateways[key] = that.Gateway(key)
		}
	}
	return gateways
}

// 发出短信
func (that SmsGateway) Send(to interface{}, data map[string]interface{}, gateways []string) (results *contact.Results, err error) {
	var (
		phoneNumber = NewPhoneNumber(to)
		message     = NewMessage(data)
	)
	if gateways == nil || len(gateways) < 1 {
		gateways = that.Config.Defaults
	}
	// if gateways == nil || len(gateways) < 1 {
	// 	// 获取全部网关
	// }
	return that.messager.Send(phoneNumber, message, that.GetGateways(gateways))
}

// 实例化
func NewSmsGateway(config *config.Sms) *SmsGateway {
	return &SmsGateway{
		Config: config,
	}
}
