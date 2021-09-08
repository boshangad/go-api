package emailService

const (
	TypeSystem = 1
	TypeNotify = 2
	TypeLogin    = 3
	TypeRegister = 3
	TypeForget = 4
	TypeSafe   = 5
)

const (
	// StatusDraft 未发布
	StatusDraft = 0
	// StatusUsed 已核销
	StatusUsed = 1
	// StatusPublished 已发布
	StatusPublished = 2
	// StatusExpire 已失效
	StatusExpire = 3
)