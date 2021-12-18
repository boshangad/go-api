package sms

import (
	"github.com/boshangad/v1/app/sms/config"
	"github.com/boshangad/v1/app/sms/interfaces"
)

type Messager struct {
	config *config.Config
}

// 发出信息
func (that Messager) Send(to interfaces.PhoneNumber, data interfaces.Message, gateways map[string]interfaces.Gateway) (results interfaces.Result, err error) {
	var (
		isSuccessFul = false
	)
	for key, gateway := range gateways {
		_, err := gateway.Send(to, data, that.config.Gateways[key])
		if err != nil {
			// results.Set(key, result)
			continue
		}
		// results.Set(key, result)
		isSuccessFul = true
		break
	}
	if !isSuccessFul {
		return nil, results
	}
	return results, nil
}

// 发送人
func NewMessager(config *config.Config) *Messager {
	return &Messager{
		config: config,
	}
}
