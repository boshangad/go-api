package db

import (
	"time"

	"entgo.io/ent/dialect/sql"

	"entgo.io/ent/dialect"
	"github.com/boshangad/v1/ent"
	_ "github.com/go-sql-driver/mysql"
	"github.com/mitchellh/mapstructure"
)

type Mysql struct {
	Driver string `json:"driver,omitempty" yaml:"driver"`
	// dsn连接字符串
	Dsn string `json:"dsn,omitempty" yaml:"dsn"`
	// 服务地址
	Host string `json:"host,omitempty" yaml:"host"`
	// 登录用户名
	Username string `json:"username,omitempty" yaml:"username"`
	// 用户密码
	Password string `json:"password,omitempty" yaml:"password"`
	// 链接方式
	Protocol string `json:"protocol,omitempty" yaml:"protocol"`
	// 数据库名称
	Dbname string `json:"dbname,omitempty" yaml:"dbname"`
	// 高级连接参数
	Params string `json:"params,omitempty" yaml:"params"`
	// 空闲中的最大连接数
	MaxIdleConns int `json:"maxdleConns,omitempty" yaml:"maxIdleConns"`
	// 打开到数据库的最大连接数
	MaxOpenConns int `json:"maxOpenConns,omitempty" yaml:"maxOpenConns"`
	// 可以重用连接的最长时间
	ConnMaxLifetime int64 `json:"connMaxLifetime,omitempty" yaml:"connMaxLifetime"`
	// 连接可能空闲的最长时间
	ConnMaxIdleTime int64 `json:"connMaxIdleTime,omitempty" yaml:"connMaxIdleTime"`
}

// 获取Dsn连接字符串
func (that *Mysql) dsnStr() (dsnStr string) {
	// dsn数据源格式 [username[:password]@][protocol[(address)]]/dbname[?param1=value1&...&paramN=valueN]
	if that.Dsn == "" {
		// 自己组装Dsn连接字符串
		if that.Username != "" && that.Password != "" {
			dsnStr += that.Username + ":" + that.Password + "@"
		} else if that.Username != "" {
			dsnStr += that.Username
		}
		if that.Host != "" {
			if that.Protocol == "" {
				that.Protocol = "tcp"
			}
			dsnStr += that.Protocol + "(" + that.Host + ")"
		} else if that.Protocol != "" {
			dsnStr += that.Protocol
		}
		dsnStr += "/"
		if that.Dbname != "" {
			dsnStr += that.Dbname
		}
		if that.Params != "" {
			dsnStr += that.Params
		}
		return dsnStr
	}
	return that.Dsn
}

// 打开连接
func (that *Mysql) Open() (*ent.Client, error) {
	drv, err := sql.Open(that.Driver, that.dsnStr())
	if err != nil {
		return nil, err
	}
	// 获取数据库驱动中的sql.DB对象。
	db := drv.DB()
	db.SetMaxIdleConns(that.MaxIdleConns)
	db.SetMaxOpenConns(that.MaxOpenConns)
	if that.ConnMaxLifetime != 0 {
		db.SetConnMaxLifetime(time.Duration(that.ConnMaxLifetime))
	}
	if that.ConnMaxIdleTime != 0 {
		db.SetConnMaxIdleTime(time.Duration(that.ConnMaxIdleTime))
	}
	client := ent.NewClient(ent.Driver(drv))
	return client, nil
}

// 实例化Mysql配置项
func NewMysql(data map[string]interface{}) *Mysql {
	mysql := Mysql{}
	err := mapstructure.Decode(data, &mysql)
	mysql.Driver = dialect.MySQL
	if err != nil {
		return nil
	}
	return &mysql
}
