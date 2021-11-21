package db

import (
	"time"

	basicSql "database/sql"

	"entgo.io/ent/dialect/sql"

	"github.com/boshangad/v1/ent"
	_ "github.com/go-sql-driver/mysql"
)

type Mysql struct {
	// ent模型客户端
	client *ent.Client
	// 打开的数据库连接器
	db *basicSql.DB
	// 数据库类型
	Driver string `mapstructure:"driver" json:"driver" yaml:"driver"`
	// dsn连接字符串
	Dsn string `mapstructure:"dsn" json:"dsn" yaml:"dsn"`
	// 服务地址
	Host string `mapstructure:"host" json:"host" yaml:"host"`
	// 登录用户名
	Username string `mapstructure:"username" json:"username" yaml:"username"`
	// 用户密码
	Password string `mapstructure:"password" json:"password" yaml:"password"`
	// 链接方式
	Protocol string `mapstructure:"protocol" json:"protocol" yaml:"protocol"`
	// 数据库名称
	Dbname string `mapstructure:"dbname" json:"dbname" yaml:"dbname"`
	// 高级连接参数
	Params string `mapstructure:"params" json:"params" yaml:"params"`
	// 空闲中的最大连接数
	MaxIdleConns int `mapstructure:"maxIdleConns" json:"maxdleConns" yaml:"maxIdleConns"`
	// 打开到数据库的最大连接数
	MaxOpenConns int `mapstructure:"maxOpenConns" json:"maxOpenConns" yaml:"maxOpenConns"`
	// 可以重用连接的最长时间
	ConnMaxLifetime int64 `mapstructure:"connMaxLifetime" json:"connMaxLifetime" yaml:"connMaxLifetime"`
	// 连接可能空闲的最长时间
	ConnMaxIdleTime int64 `mapstructure:"connMaxIdleTime" json:"connMaxIdleTime" yaml:"connMaxIdleTime"`
}

// 获取Dsn连接字符串
func (that *Mysql) dsnStr() string {
	// dsn数据源格式 [username[:password]@][protocol[(address)]]/dbname[?param1=value1&...&paramN=valueN]
	var dsnStr string = that.Dsn
	if dsnStr == "" {
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
		that.Dsn = dsnStr
	}
	return dsnStr
}

// 打开连接
func (that *Mysql) Open() (*ent.Client, error) {
	if that.client == nil {
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
		that.client = client
		that.db = db
		return client, nil
	}
	return that.client, nil
}

// 关闭数据库连接
func (that *Mysql) Close() error {
	if that.client != nil {
		c := that.client
		that.client = nil
		e := c.Close()
		return e
	}
	return nil
}
