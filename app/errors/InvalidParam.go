package errors

import (
	"net/http"
)

// 参数错误信息
type InvalidParamData struct {
	// 具体错误信息
	Errors []InvalidError `json:"errors,omitempty"`
}

// InvalidParam represents an exception caused by invalid parameters passed to a method.
type InvalidParam struct {
	// 状态码
	status int `json:"-"`
	// 错误码
	Code int64 `json:"error,omitempty"`
	// 错误消息
	Msg string `json:"msg,omitempty"`
	// 数据
	Data *InvalidParamData `json:"data,omitempty"`
}

// 获取状态值
func (that InvalidParam) GetStatus() int {
	if that.status == 0 {
		return ErrInvalidParam
	}
	return that.status
}

// 设置状态码
func (that *InvalidParam) SetStatus(status int) {
	that.status = status
}

func (that InvalidParam) Error() string {
	return that.Msg
}

// 实例化新的参数错误实例
func NewInvalidParam() *InvalidParam {
	return &InvalidParam{
		status: http.StatusBadRequest,
		Code:   ErrInvalidParam,
		Msg:    "invalid parameter",
	}
}
