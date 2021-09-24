package utils

import (
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/mem"
	"runtime"
	"time"
)

const (
	B  = 1
	KB = 1024 * B
	MB = 1024 * KB
	GB = 1024 * MB
)

type Server struct {
	// 系统信息
	Os   Os   `json:"os"`
	// CPU信息
	Cpu  Cpu  `json:"cpu"`
	// RAM信息
	Rrm  Rrm  `json:"ram"`
	// 磁盘信息
	Disk Disk `json:"disk"`
}

//Os 系统信息结果
type Os struct {
	GOOS         string `json:"goos"`
	NumCPU       int    `json:"numCpu"`
	Compiler     string `json:"compiler"`
	GoVersion    string `json:"goVersion"`
	NumGoroutine int    `json:"numGoroutine"`
}

//Cpu CPU单元信息
type Cpu struct {
	Cpus  []float64 `json:"cpus"`
	Cores int       `json:"cores"`
}

//Rrm RAM信息
type Rrm struct {
	UsedMB      int `json:"usedMb"`
	TotalMB     int `json:"totalMb"`
	UsedPercent int `json:"usedPercent"`
}

//Disk 磁盘信息
type Disk struct {
	UsedMB      int `json:"usedMb"`
	UsedGB      int `json:"usedGb"`
	TotalMB     int `json:"totalMb"`
	TotalGB     int `json:"totalGb"`
	UsedPercent int `json:"usedPercent"`
}

func init()  {
	InitOS()
}

//InitOS 初始化OS信息
//@function: InitCPU
//@description: OS信息
//@return: o Os, err error
func InitOS() (o Os) {
	o.GOOS = runtime.GOOS
	o.NumCPU = runtime.NumCPU()
	o.Compiler = runtime.Compiler
	o.GoVersion = runtime.Version()
	o.NumGoroutine = runtime.NumGoroutine()
	return o
}

//InitCPU 初始化CPU信息
//@function: InitCPU
//@description: CPU信息
//@return: c Cpu, err error
func InitCPU() (c Cpu, err error) {
	var (
		cores int
		cpus []float64
	)
	if cores, err = cpu.Counts(false); err != nil {
		return
	}
	c.Cores = cores
	if cpus, err = cpu.Percent(time.Duration(200)*time.Millisecond, true); err != nil {
		return
	}
	c.Cpus = cpus
	return
}

//InitRAM 初始化RAM信息
//@function: InitRAM
//@description: ARM信息
//@return: r Rrm, err error
func InitRAM() (r Rrm, err error) {
	var u *mem.VirtualMemoryStat
	if u, err = mem.VirtualMemory(); err != nil {
		return
	}
	r.UsedMB = int(u.Used) / MB
	r.TotalMB = int(u.Total) / MB
	r.UsedPercent = int(u.UsedPercent)
	return
}

//InitDisk 初始化硬盘信息
//@function: InitDisk
//@description: 硬盘信息
//@return: d Disk, err error
func InitDisk() (d Disk, err error) {
	var u *disk.UsageStat
	if u, err = disk.Usage("/"); err != nil {
		return d, err
	} else {
		d.UsedMB = int(u.Used) / MB
		d.UsedGB = int(u.Used) / GB
		d.TotalMB = int(u.Total) / MB
		d.TotalGB = int(u.Total) / GB
		d.UsedPercent = int(u.UsedPercent)
	}
	return d, nil
}