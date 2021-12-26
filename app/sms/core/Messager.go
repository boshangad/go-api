package core

import (
	"fmt"

	"github.com/boshangad/v1/app/sms/config"
	"github.com/boshangad/v1/app/sms/interfaces"
)

type Messager struct {
	config *config.Config
}

// 发出信息
func (that Messager) Send(to interfaces.PhoneNumber, data interfaces.Message, gateways map[string]interfaces.Gateway) (results Results, err error) {
	var (
		isSuccessFul = false
	)
	if len(gateways) < 1 {
		return results, fmt.Errorf("failed to send SMS, please check whether the gateway configuration exists")
	}
	for key, gateway := range gateways {
		result := gateway.Send(to, data, that.config.Gateways[key])
		results.Set(key, result)
		if !result.IsSuccess() {
			continue
		}
		isSuccessFul = true
		break
	}
	if !isSuccessFul {
		return results, fmt.Errorf("failed to send SMS, please check the specific data")
	}
	return results, nil
}

// 发送人
func NewMessager(config *config.Config) *Messager {
	return &Messager{
		config: config,
	}
}
