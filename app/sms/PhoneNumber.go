package sms

// 联系人
type PhoneNumber struct {
	Type   string
	IDD    string
	Mobile string
}

// 消息类型
func (that PhoneNumber) GetMessageType() string {
	return that.Type
}

// IDD 编码
func (that PhoneNumber) GetIDD() string {
	return that.IDD
}

// 手机号号码
func (that PhoneNumber) GetNumber() string {
	return that.Mobile
}

// 完整的手机号号码
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
