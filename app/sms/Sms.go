package sms

import (
	"fmt"
	"strings"
	"sync"

	// aconfig "github.com/boshangad/v1/app/config"
	"github.com/boshangad/v1/app/sms/config"
	"github.com/boshangad/v1/app/sms/interfaces"
)

// 短信推送
type SmsGateway struct {
	// 已经被初始化过的网关
	gateways map[string]interfaces.Gateway
	// 锁定
	mu *sync.RWMutex
	// 推送人
	messager *Messager
	// 配置项
	config *config.Config
}

// 网关操作人
func (that *SmsGateway) GetMessager() *Messager {
	if that.messager == nil {
		that.messager = NewMessager(that.config)
	}
	return that.messager
}

// 短信网关不是安全的进程
//
func (that *SmsGateway) Gateway(name string) interfaces.Gateway {
	if _, ok := that.gateways[name]; !ok {
		if gc, ok := that.config.Gateways[name]; !ok {
			switch strings.ToLower(gc.GetString("gateway")) {
			case "aliyun":
				// that.gateways[name] = sgateways.NewAliyunGateway(gc)
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
func (that SmsGateway) Send(to interface{}, data map[string]interface{}, gateways []string) (results interfaces.Result, err error) {
	var (
		phoneNumber = NewPhoneNumber(to)
		message     = NewMessage(data)
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
			return nil, fmt.Errorf("SMS gateway not found")
		}
	}
	return that.messager.Send(phoneNumber, message, that.GetGateways(gateways))
}

// 配置变更重启
func (that *SmsGateway) Reload(c *config.Config) {
	// 防止进程脏读
	that.mu.Lock()
	defer that.mu.Unlock()
	that.config = c
	that.gateways = make(map[string]interfaces.Gateway)
	that.messager = NewMessager(that.config)
}

// 实例化
func NewSmsGateway(c *config.Config) *SmsGateway {
	g := SmsGateway{
		config:   c,
		mu:       &sync.RWMutex{},
		gateways: make(map[string]interfaces.Gateway),
		messager: NewMessager(c),
	}
	// 注册回调
	return &g
}
