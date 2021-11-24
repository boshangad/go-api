package helpers

import (
	mrand "math/rand"
	"time"
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
