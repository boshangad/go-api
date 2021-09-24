package config

type DB struct {
	Default string
	Connections map[string]DbParams
}

type DbParams struct {
	// 数据库类型
	Driver string
	// 链接地址
	Path string
	// 登录用户名
	Username string
	// 用户密码
	Password string
	// 链接方式
	Protocol string
	// 数据库名称
	Dbname string
	// 高级连接参数
	Config string
	// 空闲中的最大连接数
	MaxIdleConns int
	// 打开到数据库的最大连接数
	MaxOpenConns int
	// 是否通过zap写入日志文件
	LogZap bool
}

func (m *DbParams) Dsn() string {
	return m.Username + ":" + m.Password + "@tcp(" + m.Path + ")/" + m.Dbname + "?" + m.Config
}