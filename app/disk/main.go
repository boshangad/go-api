package disk

import (
	"log"

	"github.com/boshangad/v1/app/config"
	"github.com/mitchellh/mapstructure"
)

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

func NewOssDisk(c config.Disk) *Oss {
	var (
		o   = Oss{}
		err = mapstructure.Decode(c, &o)
	)
	if err != nil {
		log.Panicln("oss disk initialization failed", err)
	}
	return &o
}

func NewUs3Disk(c config.Disk) *Us3 {
	var (
		o   = Us3{}
		err = mapstructure.Decode(c, &o)
	)
	if err != nil {
		log.Panicln("us3 disk initialization failed", err)
	}
	return &o
}

func NeKodoDisk(c config.Disk) *Kodo {
	var (
		o   = Kodo{}
		err = mapstructure.Decode(c, &o)
	)
	if err != nil {
		log.Panicln("kodo disk initialization failed", err)
	}
	return &o
}

func NewLocalDisk(c config.Disk) *Kodo {
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
func NewDisk(c config.Disk) DiskInterface {
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
