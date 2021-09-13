package gateways

import "strings"

type Gateway interface {
	Send(data Data, config TemplateConfig) (result string, err error)
}

type TemplateConfig struct {
	Scopes []string `json:"scopes,omitempty"`
	TemplateId string `json:"template_id,omitempty"`
	TemplateText string `json:"template_text,omitempty"`
}

type Config struct {
	Gateway string `json:"gateway,omitempty"`
	SignName string `json:"sign_name,omitempty"`
	RegionId string `json:"region_id,omitempty"`
	AccessKey string `json:"access_key,omitempty"`
	AccessSecret string `json:"access_secret,omitempty"`
	Templates []*TemplateConfig `json:"templates,omitempty"`
	TempData map[string]*TemplateConfig `json:"-"`
}

type Data struct {
	Scope string `json:"scope,omitempty"`
	DialCode string `json:"dial_code,omitempty"`
	Mobile string `json:"mobile,omitempty"`
	Content map[string]interface{} `json:"content,omitempty"`
}

// Init 初始化配置
func (that *Config) Init()  {
	if that.Templates == nil {
		that.Templates = []*TemplateConfig{}
	}
	that.TempData = make(map[string]*TemplateConfig)
	for _, template := range that.Templates {
		if template.Scopes != nil {
			for _, scope := range template.Scopes {
				that.TempData[scope] = template
			}
		}
	}
}

func NewGateway(config Config) Gateway {
	gateway := strings.ToLower(config.Gateway)
	switch gateway {
	case "aliyun":
		return NewAliyunGateway(config)
	}
	return nil
}