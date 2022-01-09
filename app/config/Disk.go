package config

type Disks map[string]interface{}

type Disk struct {
	Default string           `json:"default,omitempty" yaml:"default"`
	Disks   map[string]Disks `json:"disks,omitempty" yaml:"disks"`
}
