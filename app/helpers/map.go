package helpers

// 字符串去重
func UniqueString(x []string) []string {
	d := make([]string, 0, len(x))
	temp := map[string]struct{}{}
	for _, item := range x {
		if _, ok := temp[item]; !ok {
			temp[item] = struct{}{}
			d = append(d, item)
		}
	}
	return d
}
