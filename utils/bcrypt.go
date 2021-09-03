package utils

import (
	"golang.org/x/crypto/bcrypt"
)

func PasswordHash(passwd string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(passwd), 10)
	return string(bytes), err
}

func PasswordHashWithCost(passwd string, cost int) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(passwd), cost)
	return string(bytes), err
}

func PasswordVerify(passwd, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(passwd))
	return err == nil
}

