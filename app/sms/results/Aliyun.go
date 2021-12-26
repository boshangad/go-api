package results

type Aliyun struct {
	CodeOrigin string `xml:"Code" json:"Code"`
	Message    string `xml:"Message" json:"Message"`
	BizId      string `xml:"BizId" json:"BizId"`
	RequestId  string `xml:"RequestId" json:"RequestId"`
}

func (that Aliyun) Code() string {
	return that.CodeOrigin
}

func (that Aliyun) Error() string {
	return that.Message
}

func (that Aliyun) IsSuccess() bool {
	return that.CodeOrigin == "OK"
}
