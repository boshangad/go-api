package db

import (
	"github.com/boshangad/v1/ent"
	_ "github.com/lib/pq"
)

type PostgreSql struct {
	client *ent.Client
	// 数据库类型
	Driver string `mapstructure:"driver" json:"driver" yaml:"driver"`
	// 服务地址
	Host string `mapstructure:"host" json:"host" yaml:"host"`
	// 登录用户名
	Username string `mapstructure:"username" json:"username" yaml:"username"`
	// 用户密码
	Password string `mapstructure:"password" json:"password" yaml:"password"`
	// 数据库名称
	Dbname string `mapstructure:"dbname" json:"dbname" yaml:"dbname"`
	// 端口
	Port string `mapstructure:"port" json:"port,omitempty" yaml:"port"`
}

func (that PostgreSql) Dsn() string {
	return "host=" + that.Host + " port=" + that.Port + " user=" + that.Username + " dbname=" + that.Dbname + " password=" + that.Password
}

// 打开连接
func (that *PostgreSql) Open() (*ent.Client, error) {
	if that.client == nil {
		client, err := ent.Open(that.Driver, that.Dsn())
		if err != nil {
			return nil, err
		}
		that.client = client
	}
	return that.client, nil
}

// 关闭数据库连接
func (that *PostgreSql) Close() error {
	if that.client != nil {
		c := that.client
		that.client = nil
		e := c.Close()
		return e
	}
	return nil
}
