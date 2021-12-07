package smsLogService

type PhoneNumber struct {
	Type   string
	IDD    string
	Mobile string
}

func (that PhoneNumber) GetMessageType() string {
	return that.Type
}

func (that PhoneNumber) GetIDD() string {
	return that.IDD
}

func (that PhoneNumber) GetNumber() string {
	return that.Mobile
}

func (that PhoneNumber) GetUniversalNumber() string {
	return that.IDD + that.Mobile
}

// 实例化手机号
func NewPhoneNumber(mobile interface{}) *PhoneNumber {
	return &PhoneNumber{
		Type:   "text",
		IDD:    "86",
		Mobile: mobile.(string),
	}
}
