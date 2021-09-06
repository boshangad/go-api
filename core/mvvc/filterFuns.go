package mvvc

import (
	"strings"
)

var filterFuncs = make(map[string]func(string) string)

func init()  {
	filterFuncs["trim"] = strings.TrimSpace
}

// Filter 过滤参数
func Filter(data interface{}, funcs []string) interface{} {
	if str, ok := data.(string); ok {
		for _, f := range funcs {
			if fun, ok := filterFuncs[f]; ok {
				str = fun(str)
			}
		}
		return str
	} else if strArr, ok := data.([]string); ok {
		for k, str := range strArr {
			strArr[k] = Filter(str, funcs).(string)
		}
		return strArr
	}
	return data
}