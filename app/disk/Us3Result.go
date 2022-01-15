package disk

import "net/http"

type Us3Result struct {
	raw     *http.Response
	req     *http.Request
	RetCode int
	ErrMsg  string
	// 已上传分片的哈希值
	Etag string
	// 请求失败时返回本次请求的会话Id
	XSessionId string
	// 本次分片上传的上传Id
	UploadId string `json:"UploadId,omitempty"`
	// 分片的块大小
	BlkSize int64 `json:"BlkSize,omitempty"`
	// 已上传文件所属Bucket的名称
	Bucket string `json:"Bucket,omitempty"`
	// 已上传文件在Bucket中的Key名称
	Key string `json:"Key,omitempty"`
	// 已上传文件的大小
	FileSize int64
	// 本次分片上传的分片号码
	PartNumber int
}

func (that Us3Result) Error() string {
	return that.ErrMsg
}
