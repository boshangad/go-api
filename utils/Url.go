package utils

import "strings"

// MapToQueryString 将字符串MAp数据转换为Query String 格式
func MapToQueryString(data map[string]string) string {
	str := ""
	for key,value := range data {
		str += key + "=" + value + "&"
	}
	str = strings.TrimRight(str, "&")
	return str
}
