package global

const (
	// ErrSuccess 程序运行正常
	ErrSuccess int64 = 0
	ErrNotice int64 = 10000
	// ErrRequiredLogin 使用本数据需要登录
	ErrRequiredLogin int64 = 40500
	// ErrMissLoginParams 缺失登录参数
	ErrMissLoginParams int64 = 40501
	// ErrExpiredLogin 登录已失效
	ErrExpiredLogin int64 = 40510
)