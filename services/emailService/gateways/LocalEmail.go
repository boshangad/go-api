package gateways

import (
	"crypto/tls"
	"fmt"
	"net/smtp"
	"strings"

	"github.com/jordan-wright/email"
	"github.com/mitchellh/mapstructure"
)

// 本地邮件网关
type LocalGateway struct {
	// 发信地址
	FromAddress string `mapstructure:"from-address" json:"from_address" yaml:"from-address"`
	// 发信人昵称
	FromAlias string `mapstructure:"from-alias" json:"from_alias" yaml:"from-alias"`
	// 服务器地址
	Host string `mapstructure:"host" json:"host" yaml:"host"`
	// 端口
	Port int `mapstructure:"port" json:"port" yaml:"port"`
	// 是否SSL
	IsSSL bool `mapstructure:"is-ssl" json:"is_ssl" yaml:"is-ssl"`
	// 用户名
	Username string `mapstructure:"username" json:"username" yaml:"username"`
	// 密码
	Password string `mapstructure:"password" json:"password" yaml:"password"`
}

// SMTP 邮件服务
func (that LocalGateway) sendSmtp(emails []string, subject, content string, cc, bcc []string) (err error) {
	var (
		hostAddr    = ""
		smtpAuth    smtp.Auth
		emailClient *email.Email
	)
	smtpAuth = smtp.PlainAuth("", that.Username, that.Password, that.Host)
	emailClient = email.NewEmail()
	if that.FromAlias != "" {
		emailClient.From = fmt.Sprintf("%s <%s>", that.FromAlias, that.FromAddress)
	} else {
		emailClient.From = that.FromAddress
	}
	emailClient.To = emails
	emailClient.Subject = subject
	emailClient.HTML = []byte(content)
	if cc != nil {
		emailClient.Cc = cc
	}
	if bcc != nil {
		emailClient.Bcc = bcc
	}
	if that.Port > 0 {
		hostAddr = fmt.Sprintf("%s:%d", that.Host, that.Port)
	} else {
		hostAddr = fmt.Sprintf("%s:%d", that.Host, 587)
	}
	if that.IsSSL {
		err = emailClient.SendWithTLS(hostAddr, smtpAuth, &tls.Config{ServerName: that.Host})
	} else {
		err = emailClient.Send(hostAddr, smtpAuth)
	}
	return
}

// 单邮件服务
func (that LocalGateway) Send(emailAddress, subject, content string, cc, bcc []string) (err error) {
	err = that.sendSmtp([]string{emailAddress}, subject, content, cc, bcc)
	return
}

// 多邮件服务
func (that LocalGateway) MultiSend(emails []string, subject, content string, cc, bcc []string) (isSuccess bool, errors map[string]error) {
	err := that.sendSmtp(emails, subject, content, cc, bcc)
	if err != nil {
		errors = make(map[string]error)
		for _, e := range emails {
			errors[e] = err
		}
		return
	}
	return true, nil
}

// 实例化本地邮件服务
func NewLocalGateway(gatewayConfig map[string]interface{}) (gateway *LocalGateway) {
	err := mapstructure.Decode(gatewayConfig, gateway)
	if err == nil {
		panic("Email gateway configuration is abnormal")
	}
	gateway.Host = strings.TrimSpace(gateway.Host)
	if gateway.Host == "" {
		gateway.Host = "localhost"
	}
	gateway.FromAddress = strings.TrimSpace(gateway.FromAddress)
	if gateway.FromAddress == "" {
		panic("Email sender cannot be empty")
	}
	gateway.Username = strings.TrimSpace(gateway.Username)
	gateway.Password = strings.TrimSpace(gateway.Password)
	
	return
}
