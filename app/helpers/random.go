package helpers

import (
	"math/rand"
	mrand "math/rand"
	"time"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const (
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

func RandomInt(max int) int {
	return int(RandomInt64(int64(max)))
}

func RandomInt8(max int8) int8 {
	return int8(RandomInt64(int64(max)))
}

func RandomInt64(max int64) int64 {
	mrand.Seed(time.Now().UnixNano())
	return mrand.Int63()
}

func RandomUint64(max uint64) uint64 {
	mrand.Seed(time.Now().UnixNano())
	return mrand.Uint64()
	// ui64, _ := crand.Int(crand.Reader, big.NewInt(int64(max)))
	// return ui64.Uint64()
}

func RandomRangeInt(min, max int) int {
	return int(RandomInt(max-min) + min)
}

func RandomRange64(min, max int64) int64 {
	return RandomInt64(max-min) + min
}

// RandStringBytesMaskImpr 随机字符串
func RandStringBytesMaskImpr(n int) string {
	b := make([]byte, n)
	// A rand.Int63() generates 63 random bits, enough for letterIdxMax letters!
	for i, cache, remain := n-1, rand.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = rand.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}
	return string(b)
}
