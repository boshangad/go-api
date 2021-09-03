package config

import "fmt"

type appConfig struct {
	// 监听地址
	Host string `json:"host,omitempty"`
	// 监听的端口号
	Port int64 `json:"port,omitempty"`
	// 静态资源问卷地址
	StaticPath string `json:"static_path,omitempty"`
	// 上传文件路径
	UploadPath string `json:"upload_path,omitempty"`
	// 静态资源路径URL地址
	StaticUrl string `json:"static_url,omitempty"`
	// 可访问接口路径
	APIUrl string `json:"api_url,omitempty"`
}

// 初始化应用配置的默认数据
func (c *appConfig) initDefaultData() *appConfig {
	// 检查文件是否存在
	if c.Host == "" {
		c.Host = "127.0.0.1"
	}
	if c.Port == 0 {
		c.Port = 8080
	}
	if c.StaticPath == "" {
		c.StaticPath = "/static"
	}
	if c.UploadPath == "" {
		c.UploadPath = fmt.Sprintf("/upload")
	}
	if c.StaticUrl == "" {
		c.StaticUrl = fmt.Sprintf("%s%s", c.APIUrl, c.StaticPath)
	}
	return c
}

// GetApiUrl 获取链接地址
func (c appConfig) GetApiUrl() string {
	return fmt.Sprintf("https://%s:%d", c.Host, c.Port)
}

// GetStaticUrl 获取链接地址
func (c appConfig) GetStaticUrl() string {
	return fmt.Sprintf("%s%s", c.GetApiUrl(), c.StaticPath)
}
