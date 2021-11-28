package controller

// 上下文错误码
type ContextError struct {
	// 错误码
	Code int64 `json:"error,omitempty"`
	// 错误信息
	Msg string `json:"msg,omitempty"`
	// 返回参数
	Data interface{} `json:"data,omitempty"`
}

// 实现错误展示信息
func (that ContextError) Error() string {
	return that.Msg
}

// 实例化错误信息
func NewError(code int64, msg string, data interface{}) ContextError {
	var err = ContextError{
		Code: 0,
		Msg:  "success",
		Data: nil,
	}
	err.Code = code
	if msg != "" {
		err.Msg = msg
	}
	if data != nil {
		err.Data = data
	}
	return err
}
