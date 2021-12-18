package db

import (
	"encoding/json"
	"log"
	"strings"
	"time"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"

	"github.com/boshangad/v1/app/helpers"
	"github.com/boshangad/v1/ent"
	_ "github.com/go-sql-driver/mysql"
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
	MaxIdleConns helpers.Int `json:"maxIdleConns,omitempty" yaml:"maxIdleConns"`
	// 打开到数据库的最大连接数
	MaxOpenConns helpers.Int `json:"maxOpenConns,omitempty" yaml:"maxOpenConns"`
	// 可以重用连接的最长时间
	ConnMaxLifetime helpers.Int64 `json:"connMaxLifetime,omitempty" yaml:"connMaxLifetime"`
	// 连接可能空闲的最长时间
	ConnMaxIdleTime helpers.Int64 `json:"connMaxIdleTime,omitempty" yaml:"connMaxIdleTime"`
}

// 获取Dsn连接字符串
func (that *Mysql) dsnStr() (dsnStr string) {
	// dsn数据源格式 [username[:password]@][protocol[(address)]]/dbname[?param1=value1&...&paramN=valueN]
	if strings.TrimSpace(that.Dsn) == "" {
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
	db.SetMaxIdleConns(int(that.MaxIdleConns))
	db.SetMaxOpenConns(int(that.MaxOpenConns))
	if that.ConnMaxLifetime != 0 {
		db.SetConnMaxLifetime(time.Duration(that.ConnMaxLifetime))
	}
	if that.ConnMaxIdleTime != 0 {
		db.SetConnMaxIdleTime(time.Duration(that.ConnMaxIdleTime))
	}
	client := ent.NewClient(ent.Driver(drv))
	return client, nil
}

func (that *Mysql) Close() error {
	return nil
}

// 实例化Mysql配置项
func NewMysql(data map[string]interface{}) *Mysql {
	dataByte, err := json.Marshal(data)
	if err != nil {
		log.Println("failed to convert mysql configuration information to byte", err)
		return nil
	}
	mysql := Mysql{}
	if err = json.Unmarshal(dataByte, &mysql); err != nil {
		log.Println("mysql configuration information transfer structure failed", err)
		return nil
	}
	mysql.Driver = dialect.MySQL
	return &mysql
}
