package smsLogService

// 验证码
type Message struct {
	DialCode string `json:"dialCode,omitempty" binding:"omitempty"`
	Mobile   string `json:"mobile,omitempty" binding:"required"`
	UseType  string `json:"useType,omitempty" binding:"required"`
	Captcha  string `json:"captcha,omitempty"`
}

func NewMessage(data map[string]interface{}) *Message {
	return &Message{}
}
