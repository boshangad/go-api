package public

import (
	"fmt"

	"github.com/boshangad/v1/app/controller"
	"github.com/boshangad/v1/app/disk"
	"github.com/boshangad/v1/app/global"
	"github.com/boshangad/v1/services/resourceFileService"
)

type ResourceController struct {
}

// Image 图片预览
func (that ResourceController) Image(c *controller.Context) {

	c.JsonOut(global.ErrSuccess, "success", struct {
		Uuid string `json:"uuid"`
		Path string `json:"path"`
		Url  string `json:"url"`
	}{})
}

// Download 文件下载
func (that ResourceController) Download(c *controller.Context) {
	c.JsonOut(global.ErrSuccess, "success", struct {
		Uuid string `json:"uuid"`
		Path string `json:"path"`
		Url  string `json:"url"`
	}{})
}

// Upload 上传文件
func (that ResourceController) Upload(c *controller.Context) {
	c.FormFile("file")
	c.JsonOut(global.ErrSuccess, "success", struct {
		Uuid string `json:"uuid"`
		Path string `json:"path"`
		Url  string `json:"url"`
	}{})
}

// Us3Token Ucloud的Us3对象存储上传token
// https://docs.ucloud.cn/ufile/api/authorization
func (that ResourceController) Us3Token(c *controller.Context) {
	d := disk.NewUs3Disk(global.Config.Disk.Disks["ucloud"])
	fmt.Println(d.Upload("C:\\Users\\huanghu\\Videos\\Captures\\小狗狗.mp4", "upload/video/20220112/4567893487tydeh.mp4"))
	var (
		us3Token = resourceFileService.Us3Token{}
	)
	if err := c.ShouldBindValue(&us3Token); err != nil {
		c.JsonOutError(err)
		return
	}
	us3TokenResult, err := us3Token.Build()
	if err != nil {
		c.JsonOutError(err)
		return
	}
	c.JsonOut(global.ErrSuccess, "success", us3TokenResult)
}
