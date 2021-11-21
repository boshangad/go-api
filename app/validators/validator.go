package validators

import v10 "github.com/go-playground/validator/v10"

var (
	DefaultValidate *v10.Validate = v10.New()
)

// 是否是一个URL地址
func IsUrl(url string) bool {
	return DefaultValidate.Var(url, "url") == nil
}

// 是否是一个URI地址
func IsUri(uri string) bool {
	return DefaultValidate.Var(uri, "uri") == nil
}
