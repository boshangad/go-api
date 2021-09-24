package utils

import (
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

//PathExists 目录路径和文件路径是否存在
//@function: PathExists
//@description: 目录路径和文件路径是否存在
//@param: path string
//@return: bool, error
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

// PathIsDir 路径是否为目录路径
//@function: PathIsDir
//@description: 路径是否为目录路径
//@param: path string
//@return: bool
func PathIsDir(path string) bool {
	s, err := os.Stat(path)
	if err != nil {
		return false
	}
	return s.IsDir()
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
	if !PathIsDir(filepath.Dir(filename)) {
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