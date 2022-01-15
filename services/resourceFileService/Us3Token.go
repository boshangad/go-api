package resourceFileService

type Us3Token struct {
	DiskName    string `json:"diskName" form:"diskName"`
	Filename    string `json:"filename" form:"filename"`
	Method      string `json:"method" form:"method"`
	ContentType string `json:"contentType" form:"contentType"`
	ContentMd5  string `json:"contentMd5" form:"contentMd5"`
}

type Us3TokenResult struct {
	Method    string `json:"method"`
	Key       string `json:"key"`
	Token     string `json:"token"`
	Time      int64  `json:"time"`
	UploadUrl string `json:"uploadUrl"`
}

func (that Us3Token) Build() (result Us3TokenResult, err error) {

	return result, nil
}
