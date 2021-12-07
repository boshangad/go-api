package interfaces

import "github.com/boshangad/v1/services/smsLogService/contact"

type Gateway interface {
	Name() string
	Send(to PhoneNumber, message Message) (result *contact.Result, err error)
}
