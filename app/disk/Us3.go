package disk

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"net/http"
	"strconv"
	"time"
)

var us3RegionList = map[string]RegionDomain{
	// 北京
	"bj": {
		Extranet: "cn-bj.ufileos.com",
		Intranet: "ufile.cn-north-02.ucloud.cn",
	},
	// 上海二
	"sh2": {
		Extranet: "cn-sh2.ufileos.com",
		Intranet: "internal-cn-sh2-01.ufileos.com",
	},
	// 广州
	"gd": {
		Extranet: "cn-gd.ufileos.com",
		Intranet: "internal-cn-gd-02.ufileos.com",
	},
	// 香港
	"hk": {
		Extranet: "hk.ufileos.com",
		Intranet: "internal-hk-01.ufileos.com",
	},
	// 洛杉矶
	"us-ca": {
		Extranet: "us-ca.ufileos.com",
		Intranet: "internal-us-ca-01.ufileos.com",
	},
	// 新加坡
	"sg": {
		Extranet: "sg.ufileos.com",
		Intranet: "internal-sg-01.ufileos.com",
	},
	// 雅加达
	"idn-jakarta": {
		Extranet: "idn-jakarta.ufileos.com",
		Intranet: "	internal-idn-jakarta-01.ufileos.com",
	},
	// 台北
	"tw-tp": {
		Extranet: "tw-tp.ufileos.com",
		Intranet: "internal-tw-tp.ufileos.com",
	},
	// 拉各斯
	"afr-nigeria": {
		Extranet: "afr-nigeria.ufileos.com",
		Intranet: "internal-afr-nigeria.ufileos.com",
	},
	// 圣保罗
	"bra-saopaulo": {
		Extranet: "bra-saopaulo.ufileos.com",
		Intranet: "internal-bra-saopaulo.ufileos.com",
	},
	// 迪拜
	"uae-dubai": {
		Extranet: "uae-dubai.ufileos.com",
		Intranet: "internal-uae-dubai.ufileos.com",
	},
	// 法兰克福
	"ge-fra": {
		Extranet: "ge-fra.ufileos.com",
		Intranet: "internal-ge-fra.ufileos.com",
	},
	// 胡志明市
	"vn-sng": {
		Extranet: "vn-sng.ufileos.com",
		Intranet: "internal-vn-sng.ufileos.com",
	},
	// 华盛顿
	"us-ws": {
		Extranet: "us-ws.ufileos.com",
		Intranet: "internal-us-ws.ufileos.com",
	},
	// 孟买
	"ind-mumbai": {
		Extranet: "ind-mumbai.ufileos.com",
		Intranet: "internal-ind-mumbai.ufileos.com",
	},
	// 首尔
	"kr-seoul": {
		Extranet: "kr-seoul.ufileos.com",
		Intranet: "internal-kr-seoul.ufileos.com",
	},
	// 东京
	"jpn-tky": {
		Extranet: "jpn-tky.ufileos.com",
		Intranet: "internal-jpn-tky.ufileos.com",
	},
	// 曼谷
	"th-bkk": {
		Extranet: "th-bkk.ufileos.com",
		Intranet: "internal-th-bkk.ufileos.com",
	},
}

// Us3 Ucloud对象存储
type Us3 struct {
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

func (that Us3) Name() string {
	return "us3"
}

// 获取鉴权token
func (that Us3) AuthToken(method, key, contentMd5, contentType string, ucloudHeaders map[string]string) string {
	var (
		dateTime                   = time.Now().UnixMicro()
		canonicalizedUCloudHeaders = ""
		canonicalizedResource      = ""
		stringToSign               = ""
		signature                  = ""
		authorization              = ""
	)
	// 请求方式
	if method == "" {
		method = http.MethodPut
	}
	// / + bucketName + "/" + key
	canonicalizedResource = "/" + that.BucketName + "/" + key
	// HTTP-Verb + "\n" + Content-MD5 + "\n" + Content-Type + "\n" + Date + "\n"
	stringToSign = method + "\n" + contentMd5 + "\n" + contentType + "\n" + strconv.FormatInt(dateTime, 10) + "\n" + canonicalizedUCloudHeaders + canonicalizedResource
	// Base64( HMAC-SHA1( ucloudPrivateKey, UTF-8-Encoding-Of(stringToSign)))
	mac := hmac.New(sha1.New, []byte(that.AccessKeySecret))
	mac.Write([]byte(stringToSign))
	signature = base64.StdEncoding.EncodeToString(mac.Sum(nil))
	// authorization = "UCloud" + " " + ucloudPublicKey + ":" + signature
	authorization = "UCloud" + " " + that.AccessKeyId + ":" + signature
	return authorization
}

func (that Us3) Upload(filename, path string) (err error) {
	return nil
}
