package smsLogService

import (
	"github.com/boshangad/v1/services/smsLogService/contact"
	"github.com/boshangad/v1/services/smsLogService/interfaces"
)

type Messager struct {
}

func (that Messager) Send(to interfaces.PhoneNumber, data interfaces.Message, gateways map[string]interfaces.Gateway) (results *contact.Results, err error) {
	var (
		isSuccessFul = false
	)
	for key, gateway := range gateways {
		result, err := gateway.Send(to, data)
		if err != nil {
			results.Set(key, result)
			continue
		}
		result.Success = true
		results.Set(key, result)
		isSuccessFul = true
		break
	}
	if !isSuccessFul {
		return nil, results
	}
	return results, nil
}

func NewMessager() *Messager {
	return &Messager{}
}
