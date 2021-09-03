package global

import (
	"log"
	"os"
	"path/filepath"
)

// GetPathWithExec 获取项目执行路径
func GetPathWithExec() string {
	if RootPath == "" {
		ex, err := os.Executable()
		if err != nil {
			log.Println("获取执行目录失败", err)
			return ""
		}
		RootPath = filepath.Dir(ex)
	}
	return RootPath
}

// GetPathWithStatic 获取静态资源路径
func GetPathWithStatic() string {
	if StaticPath == "" {
		StaticPath = GetPathWithExec() + "/static"
	}
	return StaticPath
}

// GetPathWithUpload 获取上传文件路径
func GetPathWithUpload() string {
	if UploadPath == "" {
		UploadPath = GetPathWithStatic() + "/upload"
	}
	return UploadPath
}

// GetStaticFileUrl 获取可访问路径
func GetStaticFileUrl(path string) (str string) {
	str = ""
	if path == "" {
		return
	}
	str = "/static/upload/" + path
	return
}