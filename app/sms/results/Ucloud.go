package results

import "strconv"

// ucloud网关返还数据结构
type Ucloud struct {
	// 错误码
	RetCode int
	// 操作指令名称
	Action string
	// 返回错误消息，当 RetCode 非 0 时提供详细的描述信息
	Message string
	// 本次提交发送的短信的唯一ID，可根据该值查询本次发送的短信列表
	SessionNo string
	// 本次提交的自定义业务标识ID，仅当发送时传入有效的UserId，才返回该字段。
	UserId string
}

// 错误码
func (that Ucloud) Code() string {
	return strconv.Itoa(that.RetCode)
}

// 错误信息
func (that Ucloud) Error() string {
	return that.Message
}

func (that Ucloud) IsSuccess() bool {
	return that.RetCode == 0
}
