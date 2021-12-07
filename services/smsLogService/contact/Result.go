package contact

type Result struct {
	// 访问状态
	Status int
	// 访问正文
	Body string
	// 成功
	Success bool
	// 错误码
	Code int64
	// 错误字符串
	Err string
	// 错误文案
	Message string
}

func (that Result) Error() string {
	return ""
}
