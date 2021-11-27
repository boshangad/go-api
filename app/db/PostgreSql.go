package db

import (
	"database/sql"

	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	"github.com/boshangad/v1/ent"
	_ "github.com/jackc/pgx/v4/stdlib"
)

type PostgreSql struct {
	// 数据库类型
	Driver string `json:"driver,omitempty" yaml:"driver"`
	// DSN 连接
	Dsn string `json:"dsn,omitempty" yaml:"dsn"`
	// 服务地址
	Host string `json:"host,omitempty" yaml:"host"`
	// 登录用户名
	Username string `json:"username,omitempty" yaml:"username"`
	// 用户密码
	Password string `json:"password,omitempty" yaml:"password"`
	// 数据库名称
	Dbname string `json:"dbname,omitempty" yaml:"dbname"`
	// 端口
	Port string `json:"port,omitempty" yaml:"port"`
}

func (that PostgreSql) dsnStr() string {
	return "host=" + that.Host + " port=" + that.Port + " user=" + that.Username + " dbname=" + that.Dbname + " password=" + that.Password
}

// 打开连接
func (that *PostgreSql) Open() (client *ent.Client, err error) {
	// postgresql://user:password@127.0.0.1/database
	db, err := sql.Open("pgx", that.dsnStr())
	if err != nil {
		return nil, err
	}
	drv := entsql.OpenDB(dialect.Postgres, db)
	return ent.NewClient(ent.Driver(drv)), nil
}
