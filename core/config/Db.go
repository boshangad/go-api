package config

// dbConnection 数据库连接器
type dbConnection struct {
	// 连接的服务
	Driver string `json:"driver"`
	// 数据库文件地址
	Database string `json:"database,omitempty"`
	// 连接的用户名
	Username string `json:"username,omitempty"`
	// 密码
	Password string `json:"password,omitempty"`
	// 连接协议
	Protocol string `json:"protocol,omitempty"`
	// 地址
	Host string `json:"host,omitempty"`
	// 端口
	Port int `json:"port,omitempty"`
	// 数据库名称
	Dbname string `json:"dbname,omitempty"`
	// 其它参数
	Extra map[string]string `json:"extra,omitempty"`
}

// Db连接配置
type dbConfig struct {
	Default string `json:"default,omitempty"`
	Connections map[string]dbConnection `json:"connections,omitempty"`
}

// 初始化数据
func (dc *dbConfig) initDefaultData() {
	if dc.Default == "" {
		dc.Default = "db"
	}
}