package disk

import (
	"log"

	"github.com/mitchellh/mapstructure"
)

// 磁盘
type Disk map[string]interface{}

// 区域域名
type RegionDomain struct {
	// 外网
	Extranet string
	// 内网
	Intranet string
}

type UploadInterface interface {
	Upload(filename string, path string) (err error)
}

type DiskInterface interface{}

func NewOssDisk(c Disk) *Oss {
	var (
		o   = Oss{}
		err = mapstructure.Decode(c, &o)
	)
	if err != nil {
		log.Panicln("oss disk initialization failed", err)
	}
	return &o
}

func NewUs3Disk(c Disk) *Us3 {
	var (
		o   = Us3{}
		err = mapstructure.Decode(c, &o)
	)
	if err != nil {
		log.Panicln("us3 disk initialization failed", err)
	}
	if o.AccessKeyId == "" || o.AccessKeySecret == "" {
		log.Panicln("us3 disk initialization failed", "not find access parameter")
	}
	if o.BucketName == "" {
		log.Panicln("us3 disk initialization failed", "bucket name required")
	}
	return &o
}

func NeKodoDisk(c Disk) *Kodo {
	var (
		o   = Kodo{}
		err = mapstructure.Decode(c, &o)
	)
	if err != nil {
		log.Panicln("kodo disk initialization failed", err)
	}
	return &o
}

func NewLocalDisk(c Disk) *Kodo {
	var (
		o   = Kodo{}
		err = mapstructure.Decode(c, &o)
	)
	if err != nil {
		log.Panicln("local disk initialization failed", err)
	}
	return &o
}

// 实例化磁盘
func NewDisk(c Disk) DiskInterface {
	var diskType = "local"
	if t, ok := c["type"]; ok {
		diskType = t.(string)
	}
	switch diskType {
	case "oss":
		return NewOssDisk(c)
	case "us3":
		return NewUs3Disk(c)
	case "kodo":
		return NeKodoDisk(c)
	case "local":
		return NewLocalDisk(c)
	default:
		return NewLocalDisk(c)
	}
}
