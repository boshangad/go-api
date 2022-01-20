package resourceFileService

import (
	"path/filepath"
	"strings"
)

func GetBasePath(mimeType, filename string) string {
	var uploadBasic = ""
	if mimeType != "" {
		mimeTypes := strings.Split(mimeType, "/")
		if mimeTypes[0] == "image" {
			uploadBasic = "image"
		} else if mimeTypes[0] == "video" {
			uploadBasic = "video"
		} else if mimeTypes[0] == "audio" {
			uploadBasic = "audio"
		}
	}
	// 基础上传路径不存在
	if uploadBasic == "" {
		ext := strings.TrimLeft(filepath.Ext(filename), ".")
		switch strings.ToLower(ext) {
		case "webm":
			fallthrough
		case "mp4":
			uploadBasic = "video"
		case "wav":
			fallthrough
		case "wma":
			fallthrough
		case "mp3":
			uploadBasic = "audio"
		case "xls":
			fallthrough
		case "xlsm":
			fallthrough
		case "xlsx":
			fallthrough
		case "wps":
			fallthrough
		case "wpt":
			fallthrough
		case "doc":
			fallthrough
		case "docx":
			fallthrough
		case "ppt":
			fallthrough
		case "pptx":
			fallthrough
		case "pdf":
			uploadBasic = "document"
		case "jpg":
			fallthrough
		case "gif":
			fallthrough
		case "png":
			fallthrough
		case "bmp":
			fallthrough
		case "jpeg":
			fallthrough
		case "icon":
			uploadBasic = "image"
		default:
			uploadBasic = "file"
		}
	}
	return "upload/" + uploadBasic + "/"
}
