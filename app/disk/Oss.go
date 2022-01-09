package disk

// Oss 阿里云对象存储
type Oss struct {
	// 磁盘保存路径
	Path string `mapstructure:"path" json:"path" yaml:"path"`
	// 磁盘访问URL
	Url string `mapstructure:"url" json:"url" yaml:"url"`
	// 以下参数云存储有效
	// 磁盘所属区域
	Region string `mapstructure:"region" json:"region" yaml:"region"`
	// 存储桶名称
	BucketName string `mapstructure:"bucketName" json:"bucketName" yaml:"bucketName"`
	// 访问密钥
	AccessKeyId string `mapstructure:"accessKeyId" json:"accessKeyId" yaml:"accessKeyId"`
	// 访问私钥
	AccessKeySecret string `mapstructure:"accessKeySecret" json:"accessKeySecret" yaml:"accessKeySecret"`
}

func (that Oss) Upload(filename, path string) (err error) {
	return nil
}
