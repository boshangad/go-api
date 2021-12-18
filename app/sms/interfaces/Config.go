package interfaces

// 配置
type Config interface {
	// 字符串
	GetString(string) string
	// 获取数值
	GetInt64(string) int64
	// 获取数值
	GetUint64(string) uint64
}
