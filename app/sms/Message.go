package sms

import (
	"net/url"

	"github.com/boshangad/v1/app/sms/interfaces"
)

// 消息实体
type Message struct {
	Content  string     `json:"content"`
	Data     url.Values `json:"data"`
	Template string     `json:"template"`
}

func (that Message) GetType() string {
	return "text"
}

func (that Message) GetContent(g interfaces.Gateway) string {
	return that.Content
}

func (that Message) GetData(g interfaces.Gateway) (values url.Values) {
	return that.Data
}

func (that Message) GetTemplate(g interfaces.Gateway) string {
	return that.Template
}

func (that Message) GetGateways() (d map[string]interfaces.Gateway) {
	return
}

func NewMessage(data map[string]interface{}) *Message {
	return &Message{}
}
