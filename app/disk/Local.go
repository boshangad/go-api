package disk

import (
	"io"
	"os"
)

type Local struct {
	// 磁盘类型
	Type string `mapstructure:"type" json:"type" yaml:"type"`
	// 磁盘保存路径
	Path string `mapstructure:"path" json:"path" yaml:"path"`
	// 磁盘访问URL
	Url string `mapstructure:"url" json:"url" yaml:"url"`
}

// 上传文件
func (that Local) Upload(filename, path string) (err error) {
	src, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer src.Close()

	out, err := os.Create(path)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, src)
	return err
}
