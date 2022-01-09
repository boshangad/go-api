package public

type RouterApi struct {
	Mp       MpRouter
	Account  AccountRouter
	Captcha  CaptchaRouter
	Sms      SmsRouter
	Email    EmailRouter
	Resource ResourceRouter
}
