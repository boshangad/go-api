package disk

type AliyunDisk struct {
	// 磁盘类型
	Type string `mapstructure:"type" json:"type" yaml:"type"`
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

func (that AliyunDisk) Upload(filename, path string) (err error) {
	return nil
}
