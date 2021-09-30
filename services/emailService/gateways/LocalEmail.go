package gateways

import (
	"crypto/tls"
	"fmt"
	"github.com/jordan-wright/email"
	"net/smtp"
	"strings"
)

type LocalEmail struct {
	BaseEmail
	// 端口
	Port     int    `mapstructure:"port" json:"port" yaml:"port"`
	// 服务器地址
	Host     string `mapstructure:"host" json:"host" yaml:"host"`
	// 是否SSL
	IsSSL    bool   `mapstructure:"is-ssl" json:"isSSL" yaml:"is-ssl"`
	// 密钥
	Secret   string `mapstructure:"secret" json:"secret" yaml:"secret"`
}

func (that *LocalEmail) SetPort(port int) *LocalEmail {
	that.Port = port
	return that
}

func (that *LocalEmail) SetHost(host string) *LocalEmail {
	that.Host = host
	return that
}

func (that *LocalEmail) SetIsSSL(isSSL bool) *LocalEmail {
	that.IsSSL = isSSL
	return that
}

func (that *LocalEmail) SetSecret(secret string) *LocalEmail {
	that.Secret = secret
	return that
}

func (that LocalEmail) Send(to, subject, body string) (string, error) {
	auth := smtp.PlainAuth("", to, that.Secret, that.Host)
	e := email.NewEmail()
	if that.FromAlias != "" {
		e.From = fmt.Sprintf("%s <%s>", that.FromAlias, that.FromAddress)
	} else {
		e.From = that.FromAddress
	}
	e.To = strings.Split(to, ",")
	e.Subject = subject
	e.HTML = []byte(body)
	var err error
	hostAddr := fmt.Sprintf("%s:%d", that.Host, that.Port)
	if that.IsSSL {
		err = e.SendWithTLS(hostAddr, auth, &tls.Config{ServerName: that.Host})
	} else {
		err = e.Send(hostAddr, auth)
	}
	return "", err
}