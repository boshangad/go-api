package utils

import "regexp"

// ValidateEmail 判断是否为电子邮箱
func ValidateEmail(email string) bool {
	//pattern := `\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*` //匹配电子邮箱
	pattern := `^[0-9a-z][_.0-9a-z-]{0,31}@([0-9a-z][0-9a-z-]{0,30}[0-9a-z]\.){1,4}[a-z]{2,4}$`
	reg := regexp.MustCompile(pattern)
	return reg.MatchString(email)
}

// ValidateMobile 验证是否为有效手机号码
func ValidateMobile(mobile string) bool {
	regular := "^((13[0-9])|(14[5,7])|(15[0-3,5-9])|(17[0,3,5-8])|(18[0-9])|166|198|199|(147))\\d{8}$"
	reg := regexp.MustCompile(regular)
	return reg.MatchString(mobile)
}

// ValidatePhone 验证是否为有效电话号码
func ValidatePhone(mobile string) bool {
	regular := "^(\\(\\d{3,4}-)|\\d{3.4}-)?\\d{7,8}$"
	reg := regexp.MustCompile(regular)
	return reg.MatchString(mobile)
}
