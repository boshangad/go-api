package helpers

import (
	"os"
	"path/filepath"
	"reflect"
	"strings"
)

// 判断字符串是否在数组内
// @param string s 要验证的字符串
// @param []string strs 字符串组
func InArray(s string, strs []string) bool {
	for _, str := range strs {
		if s == str {
			return true
		}
	}
	return false
}

// 获取应用执行文件路径
func GetCurrentDirectory() string {
	// 返回绝对路径 filepath.Dir(os.Args[0])去除最后一个元素的路径
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		panic(err)
	}
	// 将\替换成/
	return strings.Replace(dir, "\\", "/", -1)
}

// TrimSpace 移除左右空格,支持结构体
func TrimSpace(target interface{}) {
	t := reflect.TypeOf(target)
	if t.Kind() != reflect.Ptr {
		return
	}
	t = t.Elem()
	v := reflect.ValueOf(target).Elem()
	for i := 0; i < t.NumField(); i++ {
		switch v.Field(i).Kind() {
		case reflect.String:
			v.Field(i).SetString(strings.TrimSpace(v.Field(i).String()))
		}
	}
}
