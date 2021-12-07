package errors

// 错误信息
type InvalidError struct {
	// 参数名称
	Field string `json:"field,omitempty"`
	// 错误信息
	Msg string `json:"msg,omitempty"`
}
