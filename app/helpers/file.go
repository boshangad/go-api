package helpers

import (
	"os"
)

// IsExists 文件或路径是否存在
func IsExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

// IsDir 是否是目录路径
func IsDir(path string) bool {
	s, err := os.Stat(path)
	if err != nil {
		return false
	}
	return s.IsDir()
}

// IsFile 是否是一个文件
func IsFile(file string) bool {
	s, err := os.Stat(file)
	if err != nil {
		return false
	}
	return !s.IsDir()
}
