package utils

import (
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

// Exists 判断所给路径文件/文件夹是否存在
func Exists(path string) bool {
	// os.Stat获取文件信息
	_, err := os.Stat(path)
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}

// IsDir 判断所给路径是否为文件夹
func IsDir(path string) bool {
	s, err := os.Stat(path)
	if err != nil {
		return false
	}
	return s.IsDir()
}

// IsFile 判断所给路径是否为文件
func IsFile(path string) bool {
	if !Exists(path) {
		return false
	}
	return !IsDir(path)
}

func DownLoadFile(url, filename string) bool {
	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		log.Println(err)
		return false
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Println(err)
		}
	}(resp.Body)
	if !IsDir(filepath.Dir(filename)) {
		err = os.MkdirAll(filepath.Dir(filename), 0664)
		if err != nil {
			log.Println("创建目录路径失败", err)
			return false
		}
	}
	// 创建一个文件用于保存
	out, err := os.Create(filename)
	if err != nil {
		log.Println("创建文件失败", err)
		return false
	}
	defer func(out *os.File) {
		err := out.Close()
		if err != nil {
			log.Println("关闭文件输出流失败", err)
		}
	}(out)
	// 然后将响应流和文件流对接起来
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}