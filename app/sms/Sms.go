package sms

import (
	"fmt"
	"net/url"
	"strings"
	"sync"

	"github.com/boshangad/v1/app/sms/config"
	"github.com/boshangad/v1/app/sms/core"
	"github.com/boshangad/v1/app/sms/gateways"
	"github.com/boshangad/v1/app/sms/interfaces"
)

// 短信推送
type SmsGateway struct {
	// 已经被初始化过的网关
	gateways map[string]interfaces.Gateway
	// 锁定
	mu *sync.RWMutex
	// 推送人
	messager *core.Messager
	// 配置项
	config *config.Config
}

// 网关操作人
func (that *SmsGateway) GetMessager() *core.Messager {
	if that.messager == nil {
		that.messager = core.NewMessager(that.config)
	}
	return that.messager
}

// 短信网关不是安全的进程
//
func (that *SmsGateway) Gateway(name string) interfaces.Gateway {
	if _, ok := that.gateways[name]; !ok {
		if gc, ok := that.config.Gateways[name]; ok {
			gatewayName := strings.TrimSpace(gc.GetString("gateway"))
			if gatewayName == "" {
				gatewayName = name
			}
			gatewayName = strings.ToLower(gatewayName)
			switch gatewayName {
			case "aliyun":
				that.gateways[name] = gateways.NewAliyunGateway(gc)
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
				that.gateways[name] = gateways.NewUcloudGateway(gc)
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
	return that.gateways[name]
}

// 网关 安全的进程调用
func (that *SmsGateway) GetGateways(gs []string) map[string]interfaces.Gateway {
	// 锁定，防止进程重复写
	that.mu.RLock()
	defer that.mu.RUnlock()
	var gateways = make(map[string]interfaces.Gateway)
	if len(gs) > 0 {
		for _, key := range gs {
			gateways[key] = that.Gateway(key)
		}
	}
	return gateways
}

// 发出短信
func (that SmsGateway) Send(to string, data url.Values, gateways []string) (results core.Results, err error) {
	var (
		phoneNumber = core.NewPhoneNumber(to)
		message     = core.NewMessage().SetData(data)
	)
	if gateways == nil || len(gateways) < 1 {
		gateways = that.config.Defaults
	}
	if gateways == nil || len(gateways) < 1 {
		// 获取全部网关
		gateways = []string{}
		for gateway, _ := range that.config.Gateways {
			gateways = append(gateways, gateway)
		}
		if len(gateways) < 1 {
			return results, fmt.Errorf("SMS gateway not found")
		}
	}
	smsGateways := that.GetGateways(gateways)
	if smsGateways == nil || len(smsGateways) < 1 {
		return results, fmt.Errorf("can't found sms gateway")
	}
	return that.messager.Send(phoneNumber, message, smsGateways)
}

// 配置变更重启
func (that *SmsGateway) Reload(c *config.Config) {
	// 防止进程脏读
	that.mu.Lock()
	defer that.mu.Unlock()
	that.config = c
	that.gateways = make(map[string]interfaces.Gateway)
	that.messager = core.NewMessager(that.config)
}

// 实例化
func NewSmsGateway(c *config.Config) *SmsGateway {
	g := SmsGateway{
		config:   c,
		mu:       &sync.RWMutex{},
		gateways: make(map[string]interfaces.Gateway),
		messager: core.NewMessager(c),
	}
	return &g
}

// 实例化消息体
func NewMessage() *core.Message {
	return core.NewMessage()
}

// 实例化单个手机号
func NewPhoneNumber(mobile string) core.PhoneNumber {
	return core.NewPhoneNumber(mobile)
}

// 实例化多个手机号
func NewPhoneNumbers(mobile ...string) core.PhoneNumber {
	return core.NewPhoneNumber(mobile[0])
}
