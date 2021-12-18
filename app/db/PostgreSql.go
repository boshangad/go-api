package db

import (
	"database/sql"
	"encoding/json"
	"log"
	"strconv"

	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	"github.com/boshangad/v1/app/helpers"
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
	Port helpers.Int `json:"port,omitempty" yaml:"port"`
}

func (that PostgreSql) dsnStr() string {
	return "host=" + that.Host + " port=" + strconv.Itoa(int(that.Port)) + " user=" + that.Username + " dbname=" + that.Dbname + " password=" + that.Password
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

func (that *PostgreSql) Close() error {
	return nil
}

// 实例化PostgreSql
func NewPostgreSql(data map[string]interface{}) *PostgreSql {
	dataByte, err := json.Marshal(data)
	if err != nil {
		log.Println("failed to convert postgreSql configuration information to byte", err)
		return nil
	}
	postgreSql := PostgreSql{}
	if err = json.Unmarshal(dataByte, &postgreSql); err != nil {
		log.Println("postgreSql configuration information transfer structure failed", err)
		return nil
	}
	postgreSql.Driver = dialect.Postgres
	return &postgreSql
}
