package cache

import (
	"github.com/boshangad/v1/app/helpers"
	"github.com/boshangad/v1/app/validators"
)

// 创建访问的Key值
func buildingKey(key string) string {
	if key == "" {
		panic("key cannot be empty")
	}
	err := validators.DefaultValidate.Var(key, "required")
	if err != nil {
		return helpers.Md5(key)
	}
	return key
}
