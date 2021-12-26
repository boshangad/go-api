package core

import (
	"net/url"

	"github.com/boshangad/v1/app/sms/interfaces"
)

// 消息实体
type Message struct {
	// 短信类型
	t string `json:"type,omitempty"`
	// 短信模板
	templates map[string]string `json:"templates,omitempty"`
	// 短信正文
	content string `json:"content,omitempty"`
	// 参数数据
	data url.Values `json:"data,omitempty"`
	// 参数顺序键
	dataKeys []string `json:"dataKeys,omitempty"`
}

func (that Message) GetType() string {
	if that.t == "" {
		return "text"
	}
	return that.t
}

func (that *Message) SetContent(content string) *Message {
	that.content = content
	return that
}

func (that Message) GetContent(g interfaces.Gateway) string {
	return that.content
}

func (that *Message) SetData(data url.Values) *Message {
	that.data = data
	return that
}

func (that Message) GetData(g interfaces.Gateway) (values url.Values) {
	return that.data
}

func (that *Message) SetTemplate(g interfaces.Gateway, templateId string) *Message {
	that.templates[g.Name()] = templateId
	return that
}

func (that Message) GetTemplate(g interfaces.Gateway) string {
	return that.templates[g.Name()]
}

func (that Message) GetGateways() (d map[string]interfaces.Gateway) {
	return
}

func NewMessage() *Message {
	return &Message{}
}
