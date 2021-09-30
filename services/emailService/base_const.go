package emailService

const (
	StatusToSend = iota
	StatusSendSuccess
	StatusUsed
	StatusExpired
)
const (
	TypeLogin = iota + 1
	TypeRegister
	TypeForget
	TypeReset
)

var (
	// ScopeTypeList 来源与类型的关联列表
	ScopeTypeList = map[string]int64{
		"login": TypeLogin,
		"register": TypeRegister,
		"forget": TypeForget,
		"reset": TypeReset,
	}
)