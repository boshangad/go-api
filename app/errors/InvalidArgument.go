package errors

import "net/http"

// InvalidArgument represents an exception caused by invalid parameters passed to a method.
type InvalidArgument struct {
	Status int            `json:"status,omitempty"`
	Code   int64          `json:"error,omitempty"`
	Msg    string         `json:"msg,omitempty"`
	Errors []InvalidError `json:"-"`
}

func (that InvalidArgument) Error() string {
	return that.Msg
}

// 添加错误信息
func (that *InvalidArgument) AddErrors(err ...InvalidError) *InvalidArgument {
	that.Errors = append(that.Errors, err...)
	return that
}

// 实例化
func NewInvalidArgument() *InvalidArgument {
	return &InvalidArgument{
		Status: http.StatusBadRequest,
		Code:   ErrInvalidArgument,
		Msg:    "invalid parameter",
		Errors: []InvalidError{},
	}
}
