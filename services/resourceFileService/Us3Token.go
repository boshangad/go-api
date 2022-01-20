package resourceFileService

import (
	"fmt"
	"net/http"
	"net/url"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/boshangad/v1/app/disk"
	"github.com/boshangad/v1/global"
	"github.com/google/uuid"
)

type Us3Token struct {
	// 磁盘名称
	DiskName string `json:"diskName,omitempty" form:"diskName"`
	// 文件名称
	Filename string `json:"filename,omitempty" form:"filename"`
	// 请求方法
	Method string `json:"method,omitempty" form:"method"`
	// 请求类型
	ContentType string `json:"contentType,omitempty" form:"contentType"`
	// 文件的Md5值
	ContentMd5 string `json:"contentMd5,omitempty" form:"contentMd5"`
}

type Us3TokenResult struct {
	Method    string `json:"method,omitempty"`
	Key       string `json:"key,omitempty"`
	Time      int64  `json:"time,omitempty"`
	UploadUrl string `json:"uploadUrl,omitempty"`
	Signature string `json:"signature,omitempty"`
}

func (that Us3Token) Build() (result *Us3TokenResult, err error) {
	if that.DiskName == "" {
		that.DiskName = global.Config.Disk.Default
		if global.Config.Disk.Default == "" {
			that.DiskName = "ucloud"
		}
	}
	us3 := disk.NewUs3Disk(global.Config.Disk.Disks[that.DiskName])
	if us3 == nil {
		return nil, fmt.Errorf("us3 type disk configuration error")
	}
	ext := filepath.Ext(that.Filename)
	key := GetBasePath(that.ContentType, that.Filename) + time.Now().Format("20060102") + "/" + strings.ReplaceAll(uuid.New().String(), "-", "") + ext
	u := url.Values{}
	nowUnixMicro := time.Now().UnixMicro()
	req, err := http.NewRequest(strings.ToUpper(that.Method), us3.GetRequestUrl(key, u), nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-MD5", that.ContentMd5)
	req.Header.Add("Content-Type", that.ContentType)
	req.Header.Add("Date", strconv.FormatInt(nowUnixMicro, 10))
	result = &Us3TokenResult{
		Method: req.Method,
		Key:    key,
		Time:   nowUnixMicro,
		UploadUrl: strings.Join([]string{
			req.URL.Scheme,
			"://",
			req.URL.Host,
			req.URL.RequestURI(),
		}, ""),
		Signature: us3.AuthToken(req),
	}
	return result, nil
}
