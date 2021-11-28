package helpers

import (
	"golang.org/x/crypto/bcrypt"
)

// 密码加密
func PasswordHash(password string, cost int) (hash string, err error) {
	var bytes []byte
	if cost < bcrypt.MinCost || cost > bcrypt.MaxCost {
		cost = bcrypt.DefaultCost
	}
	bytes, err = bcrypt.GenerateFromPassword([]byte(password), cost)
	hash = string(bytes)
	return
}

// 验证密码
func PasswordVerify(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
