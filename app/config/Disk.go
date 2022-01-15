package config

import (
	appDisk "github.com/boshangad/v1/app/disk"
)

// 目录
type Disk struct {
	Default string                  `json:"default,omitempty" yaml:"default"`
	Disks   map[string]appDisk.Disk `json:"disks,omitempty" yaml:"disks"`
}
