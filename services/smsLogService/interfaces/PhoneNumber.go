package interfaces

type PhoneNumber interface {
	// 返回IDD代码 如:86
	GetIDD() string
	// 返回手机号 如:18888888888
	GetNumber() string
	// 返回手机号 如+8618888888888
	GetUniversalNumber() string
}
