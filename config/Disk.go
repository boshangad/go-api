package config

type DiskInterface interface {
	Url(path string) string
	ImgThumbUrl(path string, w, h int64, t string) string
}

type Disk struct {
	// 默认磁盘网关
	Default string                   `mapstructure:"default" json:"default,omitempty" yaml:"default"`
	// 磁盘列表
	Disks   map[string]DiskInterface `mapstructure:"disks" json:"disks,omitempty" yaml:"disks"`
}

// DiskLocal 本地磁盘类型
type DiskLocal struct {
	// 存储路径
	Path string `mapstructure:"path" json:"path" yaml:"path"`
}

func (DiskLocal) Url(path string) string {
	return ""
}

func (DiskLocal) ImgThumbUrl(path string, w, h int64, t string) string {
	return ""
}

// DiskAliyun 阿里云对象存储
type DiskAliyun struct {
	Endpoint        string `mapstructure:"endpoint" json:"endpoint" yaml:"endpoint"`
	AccessKeyId     string `mapstructure:"access-key-id" json:"accessKeyId" yaml:"access-key-id"`
	AccessKeySecret string `mapstructure:"access-key-secret" json:"accessKeySecret" yaml:"access-key-secret"`
	BucketName      string `mapstructure:"bucket-name" json:"bucketName" yaml:"bucket-name"`
	BucketUrl       string `mapstructure:"bucket-url" json:"bucketUrl" yaml:"bucket-url"`
	BasePath        string `mapstructure:"base-path" json:"basePath" yaml:"base-path"`
}

func (DiskAliyun) Url(path string) string {
	return ""
}

func (DiskAliyun) ImgThumbUrl(path string, w, h int64, t string) string {
	return ""
}

// DiskQiniu 七牛云对象存储
type DiskQiniu struct {
	Zone          string `mapstructure:"zone" json:"zone" yaml:"zone"`                                // 存储区域
	Bucket        string `mapstructure:"bucket" json:"bucket" yaml:"bucket"`                          // 空间名称
	ImgPath       string `mapstructure:"img-path" json:"imgPath" yaml:"img-path"`                     // CDN加速域名
	UseHTTPS      bool   `mapstructure:"use-https" json:"useHttps" yaml:"use-https"`                  // 是否使用https
	AccessKey     string `mapstructure:"access-key" json:"accessKey" yaml:"access-key"`               // 秘钥AK
	SecretKey     string `mapstructure:"secret-key" json:"secretKey" yaml:"secret-key"`               // 秘钥SK
	UseCdnDomains bool   `mapstructure:"use-cdn-domains" json:"useCdnDomains" yaml:"use-cdn-domains"` // 上传是否使用CDN上传加速
}

func (DiskQiniu) Url(path string) string {
	return ""
}

func (DiskQiniu) ImgThumbUrl(path string, w, h int64, t string) string {
	return ""
}

// DiskTencentCOS 腾讯云对象存储
type DiskTencentCOS struct {
	Bucket     string `mapstructure:"bucket" json:"bucket" yaml:"bucket"`
	Region     string `mapstructure:"region" json:"region" yaml:"region"`
	SecretID   string `mapstructure:"secret-id" json:"secretID" yaml:"secret-id"`
	SecretKey  string `mapstructure:"secret-key" json:"secretKey" yaml:"secret-key"`
	BaseURL    string `mapstructure:"base-url" json:"baseURL" yaml:"base-url"`
	PathPrefix string `mapstructure:"path-prefix" json:"pathPrefix" yaml:"path-prefix"`
}

func (DiskTencentCOS) Url(path string) string {
	return ""
}

func (DiskTencentCOS) ImgThumbUrl(path string, w, h int64, t string) string {
	return ""
}