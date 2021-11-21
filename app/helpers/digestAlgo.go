package helpers

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/hex"
)

// Md5字符串加密
func Md5(v string) string {
	var (
		d []byte = []byte(v)
		m        = md5.New()
	)
	m.Write(d)
	return hex.EncodeToString(m.Sum(nil))
}

// Sha1摘要算法
func Sha1(v string) string {
	var (
		d []byte = []byte(v)
		m        = sha1.New()
	)
	m.Write(d)
	return hex.EncodeToString(m.Sum(nil))
}

// Sha256加密
func Sha256(v string) string {
	var (
		d []byte = []byte(v)
		m        = sha256.New()
	)
	m.Write(d)
	return hex.EncodeToString(m.Sum(nil))
}
