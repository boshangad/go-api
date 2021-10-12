package utils

import "strconv"

func GetInt(val string) (i int) {
	i, _ = strconv.Atoi(val)
	return
}

func GetInt8(val string) (i int8) {
	i64, _ := strconv.ParseInt(val, 10, 8)
	return int8(i64)
}

func GetInt16(val string) (i int16) {
	i64, _ := strconv.ParseInt(val, 10, 16)
	return int16(i64)
}

func GetInt32(val string) (i int32) {
	i64, _ := strconv.ParseInt(val, 10, 32)
	return int32(i64)
}

func GetInt64(val string) (i int64) {
	i64, _ := strconv.ParseInt(val, 10, 64)
	return i64
}