package emailService

type AliyunEmail struct {
	BaseEmail

}

func (that AliyunEmail) Send(to, subject, body string) (string, error) {

}
