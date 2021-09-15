package emailService

import (
	"github.com/boshangad/go-api/core/config"
	"github.com/boshangad/go-api/core/config/email/gateways"
	"log"
)

// DefaultPushClient 默认推送客户端
func DefaultPushClient() gateways.ConfigInterface {
	emailPush := config.Get().EmailPush
	if emailPush == nil {
		log.Println("the default push mail gateway is not configured")
		return nil
	}
	return PushClient(emailPush.Default)
}

// PushClient 推送客户端
func PushClient(key string) gateways.ConfigInterface {
	emailPush := config.Get().EmailPush
	if emailPush != nil {
		if data, ok := emailPush.Clients[key]; ok {
			return data
		}
	}
	log.Panicln("email push client was not found", "gateway is", key)
	return nil
}