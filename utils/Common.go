package utils

import (
	"runtime"
	"strings"
)

// 获取正在运行的函数名
func GetFunctionName(isFullName bool) string {
	pc, _, _, _ := runtime.Caller(1)
	name := runtime.FuncForPC(pc).Name()
	if (isFullName) {
		return name
	}
	return SubStrWithLast(name, ".", 0, true)
}

// 截取从开头到最后一次发现的字符串
func SubStrWithBegin(s string, substr string, offset int, desc bool) string {
	var index int
	if desc {
		index = strings.LastIndex(s, substr)
	} else {
		index = strings.Index(s, substr)
	}
	if index == -1 {
		return s
	}
	return s[0: index]
}

// 截取从开头到最后一次发现的字符串
func SubStrWithLast(s string, substr string, offset int, desc bool) string {
	var index int
	if desc {
		index = strings.LastIndex(s, substr)
	} else {
		index = strings.Index(s, substr)
	}
	if index == -1 {
		return s
	}
	return s[index + 1: strings.Count(s, "") - 1]
}


func Substr(s string, pos, length int) string {
	runes := []rune(s)
	l := pos + length
	if l > len(runes) {
		l = len(runes)
	}
	return string(runes[pos:l])
}

func GetParentDirectory(dirctory string) string {
	return Substr(dirctory, 0, strings.LastIndex(dirctory, "/"))
}