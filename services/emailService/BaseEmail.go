package emailService

type EmailSend interface {
	Send(to, subject, body string) (string, error)
}

type BaseEmail struct {
	// 发信地址
	FromAddress string
	// 发信人昵称
	FromAlias string
}

func (that *BaseEmail) SetFromAlias(alias string) *BaseEmail {
	that.FromAlias = alias
	return that
}

func (that *BaseEmail) SetFromAddress(addr string) *BaseEmail {
	that.FromAddress = addr
	return that
}
