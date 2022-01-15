package disk

import (
	"bytes"
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha1"
	"encoding/base64"
	"encoding/binary"
	"encoding/hex"
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/boshangad/v1/app/helpers"
	"github.com/google/uuid"
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
	Path string `mapstructure:"path" json:"path,omitempty" yaml:"path"`
	// 磁盘访问URL
	Url string `mapstructure:"url" json:"url,omitempty" yaml:"url"`
	// 以下参数云存储有效
	// 磁盘所属区域
	Region string `mapstructure:"region" json:"region,omitempty" yaml:"region"`
	// 存储桶名称
	BucketName string `mapstructure:"bucketName" json:"bucketName,omitempty" yaml:"bucketName"`
	// 访问密钥
	AccessKeyId string `mapstructure:"accessKeyId" json:"accessKeyId,omitempty" yaml:"accessKeyId"`
	// 访问私钥
	AccessKeySecret string `mapstructure:"accessKeySecret" json:"accessKeySecret,omitempty" yaml:"accessKeySecret"`
	// 是否内网环境
	IsIntranet bool `mapstructure:"isIntranet" json:"isIntranet,omitempty" yaml:"isIntranet"`
}

func (that Us3) Name() string {
	return "us3"
}

// 获取鉴权token
func (that Us3) AuthToken(key string, request *http.Request) string {
	var (
		canonicalizedUCloudHeaders = ""
		canonicalizedResource      = ""
		stringToSign               = ""
		signature                  = ""
		authorization              = ""
	)
	if request.Header != nil {
		var (
			keys   []string
			values = make(map[string][]string)
		)
		for k, _ := range request.Header {
			if len(k) < 10 {
				continue
			}
			var (
				value = request.Header.Get(k)
				k1    = strings.ToLower(k)
				k2    = k1[0:9]
				k3    = k1[9:]
			)
			if k2 == "x-ucloud-" {
				keys = append(keys, k3)
				if v2, ok := values[k3]; ok {
					v2 = append(v2, value)
					values[k3] = v2
				} else {
					values[k3] = []string{value}
				}
			}
		}
		// 排序
		if len(keys) > 0 {
			sort.Strings(keys)
			for _, key := range keys {
				canonicalizedUCloudHeaders += "x-ucloud-" + key + ":" + strings.Join(helpers.UniqueString(values[key]), ",") + "\n"
			}
		}
	}
	// / + bucketName + "/" + key
	canonicalizedResource = "/" + that.BucketName + "/" + strings.TrimLeft(key, "/")
	// HTTP-Verb + "\n" + Content-MD5 + "\n" + Content-Type + "\n" + Date + "\n"
	stringToSign = request.Method + "\n" + request.Header.Get("Content-MD5") + "\n" + request.Header.Get("Content-Type") + "\n" + request.Header.Get("Date") + "\n" + canonicalizedUCloudHeaders + canonicalizedResource
	// Base64( HMAC-SHA1( ucloudPrivateKey, UTF-8-Encoding-Of(stringToSign)))
	mac := hmac.New(sha1.New, []byte(that.AccessKeySecret))
	mac.Write([]byte(stringToSign))
	signature = base64.StdEncoding.EncodeToString(mac.Sum(nil))
	// authorization = "UCloud" + " " + ucloudPublicKey + ":" + signature
	authorization = "UCloud" + " " + that.AccessKeyId + ":" + signature
	return authorization
}

func (that Us3) UploadPolicy() {

}

// 上传文件
func (that Us3) Upload(file, key string) (result *Us3Result, err error) {
	fileHandle, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer fileHandle.Close()
	f, _ := fileHandle.Stat()
	if f.Size() > 20*1024*1024 {
		_ = fileHandle.Close()
		return that.MultipartUpload(file, key)
	}
	var (
		contentType       string
		b, _              = ioutil.ReadAll(fileHandle)
		b1                = md5.Sum(b)
		b2                = b1[:]
		contentTypeBuffer = b[:512]
	)
	contentType = http.DetectContentType(contentTypeBuffer)
	req, err := http.NewRequest(http.MethodPut, "https://"+that.getRegionBasicUrl()+"/"+strings.TrimLeft(key, "/\\"), bytes.NewBuffer(b))
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", contentType)
	req.Header.Add("Content-MD5", hex.EncodeToString(b2))
	req.Header.Add("Date", strconv.FormatInt(time.Now().UnixMicro(), 10))
	result, err = that.request(req)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// 切片上传
func (that Us3) MultipartUpload(file, key string) (result *Us3Result, err error) {
	fileHandle, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer fileHandle.Close()
	initResult, err := that.InitiateMultipartUpload(file, key)
	if err != nil {
		return nil, err
	}
	var (
		contentType       string
		contentTypeBuffer          = make([]byte, 512)
		partNumber        int64    = 0
		etags             []string = []string{}
		buffer                     = make([]byte, initResult.BlkSize)
		isSuccessFul               = true
	)
	defer func() {
		if !isSuccessFul {
			_, _ = that.AbortMultipartUpload(key, initResult.UploadId)
		}
	}()
	fileHandle.Read(contentTypeBuffer)
	contentType = http.DetectContentType(contentTypeBuffer)
	fileHandle.Seek(0, io.SeekStart)
	var (
		ch = make(chan struct{}, 3)
		wg sync.WaitGroup
		// ctx = context.Background()
	)
	for {
		n, err := fileHandle.Read(buffer)
		if err != nil {
			if err == io.EOF {
				break
			}
			isSuccessFul = false
			return nil, err
		}
		wg.Add(1)
		go func(byteLength int, eqPart int64) {
			defer wg.Done()
			ch <- struct{}{}
			partResult, err := that.UploadPart(bytes.NewBuffer(buffer[:byteLength]), initResult.UploadId, initResult.Key, eqPart, contentType)
			if err != nil {
				isSuccessFul = false
				return
			}
			<-ch
			etags = append(etags, partResult.Etag)
		}(n, partNumber)
		partNumber++
	}
	wg.Wait()
	result, err = that.FinishMultipartUpload(bytes.NewReader([]byte(strings.Join(etags, ","))), initResult.UploadId, initResult.Key, "", contentType)
	if err != nil {
		isSuccessFul = false
		return nil, err
	}
	return result, nil
}

// 秒传
func (that Us3) HitUpload(file, key string) (result *Us3Result, err error) {
	fileHandle, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer fileHandle.Close()
	var (
		hash        string = ""
		fileinfo, _        = fileHandle.Stat()
	)
	hash, err = that.evalFileHash(fileHandle)
	if err != nil {
		return nil, err
	}
	urlQuery := url.Values{
		"Hash":     {hash},
		"FileName": {key},
		"FileSize": {strconv.FormatInt(fileinfo.Size(), 10)},
	}
	req, err := http.NewRequest(
		http.MethodPost,
		"https://"+that.getRegionBasicUrl()+"/uploadhit?"+urlQuery.Encode(),
		nil,
	)
	if err != nil {
		return nil, err
	}
	result, err = that.request(req)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// 删除文件
func (that Us3) Delete(key string) (err error) {
	var dateTime = strconv.FormatInt(time.Now().UnixMicro(), 10)
	req, err := http.NewRequest(http.MethodDelete, "https://"+that.getRegionBasicUrl()+"/"+strings.TrimLeft(key, "/\\"), nil)
	if err != nil {
		return err
	}
	req.Header.Add("content-Md5", uuid.New().String())
	req.Header.Add("Date", dateTime)
	_, err = that.request(req)
	if err != nil {
		return err
	}
	return nil
}

// 初始化切片计划
func (that Us3) InitiateMultipartUpload(file, key string) (result *Us3Result, err error) {
	fileinfo, err := os.Stat(file)
	if err != nil {
		return nil, err
	}
	urlQuery := url.Values{
		"uploads": {""},
	}
	req, err := http.NewRequest(
		http.MethodPost,
		"https://"+that.getRegionBasicUrl()+"/"+strings.TrimLeft(key, "/\\")+"?"+urlQuery.Encode(),
		nil,
	)
	if err != nil {
		return nil, err
	}
	req.Header.Add("content-length", strconv.FormatInt(fileinfo.Size(), 10))
	result, err = that.request(req)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// 上传切片文件
func (that Us3) UploadPart(reader io.Reader, uploadId, key string, partNumber int64, contrntType string) (result *Us3Result, err error) {
	urlQuery := url.Values{
		"uploadId":   {uploadId},
		"partNumber": {strconv.FormatInt(partNumber, 10)},
	}
	req, err := http.NewRequest(
		http.MethodPut,
		"https://"+that.getRegionBasicUrl()+"/"+strings.TrimLeft(key, "/\\")+"?"+urlQuery.Encode(),
		reader,
	)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", contrntType)
	result, err = that.request(req)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// 完成分片上传
func (that Us3) FinishMultipartUpload(reader io.Reader, uploadId, key, newKey, contrntType string) (result *Us3Result, err error) {
	urlQuery := url.Values{
		"uploadId": {uploadId},
		"newKey":   {newKey},
	}
	req, err := http.NewRequest(
		http.MethodPost,
		"https://"+that.getRegionBasicUrl()+"/"+strings.TrimLeft(key, "/\\")+"?"+urlQuery.Encode(),
		reader,
	)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", contrntType)
	result, err = that.request(req)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// 放弃分片上传
func (that Us3) AbortMultipartUpload(key, uploadId string) (result *Us3Result, err error) {
	urlQuery := url.Values{
		"uploadId": {uploadId},
	}
	req, err := http.NewRequest(
		http.MethodDelete,
		"https://"+that.getRegionBasicUrl()+"/"+strings.TrimLeft(key, "/\\")+"?"+urlQuery.Encode(),
		nil,
	)
	if err != nil {
		return nil, err
	}
	result, err = that.request(req)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// 获取请求的路径
func (that Us3) getRegionBasicUrl() string {
	if that.Region == "" {
		that.Region = "gd"
	}
	s, ok := us3RegionList[that.Region]
	if !ok {
		// 区域地址不存在
		return ""
	}
	if that.IsIntranet {
		return that.BucketName + "." + s.Intranet
	}
	return that.BucketName + "." + s.Extranet
}

// 发起请求
func (that Us3) request(req *http.Request) (result *Us3Result, err error) {
	req.Header.Add("Authorization", that.AuthToken(req.URL.Path, req))
	r := Us3Result{
		req: req,
	}
	response, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	body, _ := ioutil.ReadAll(response.Body)
	if err := json.Unmarshal(body, &r); err != nil {
		return nil, err
	}
	log.Println(req.Header.Get("Authorization"), req.URL.RawQuery, string(body))
	r.req = req
	r.raw = response
	r.XSessionId = response.Header.Get("X-SessionId")
	r.Etag = response.Header.Get("Etag")
	if r.RetCode != 0 {
		return nil, r
	}
	return &r, nil
}

// 计算文件Hash
func (that Us3) evalFileHash(file *os.File) (hash string, err error) {
	var (
		bufferSize  int64 = 4 * 1024 * 1024
		fileInfo, _       = file.Stat()
		fileSize          = fileInfo.Size()
		blkcnt            = uint32(fileSize / bufferSize)
	)
	if fileSize%bufferSize != 0 {
		blkcnt++
	}
	bs := make([]byte, 4)
	binary.LittleEndian.PutUint32(bs, blkcnt)
	h := sha1.New()
	buf := make([]byte, 0, 24)
	buf = append(buf, bs...)
	if fileSize <= bufferSize {
		_, err = io.Copy(h, file)
		if err != nil {
			return "", err
		}
	} else {
		var i uint32
		for i = 0; i < blkcnt; i++ {
			shaBlk := sha1.New()
			io.Copy(shaBlk, io.LimitReader(file, bufferSize))
			io.Copy(h, bytes.NewReader(shaBlk.Sum(nil)))
		}
	}
	buf = h.Sum(buf)
	return base64.URLEncoding.EncodeToString(buf), nil
}
