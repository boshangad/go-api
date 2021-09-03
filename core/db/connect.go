package db

import (
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
	"github.com/tend/wechatServer/core/config"
	"github.com/tend/wechatServer/ent"
	"github.com/tend/wechatServer/utils"
	"log"
	"strings"
)

// 打开连接器客户端
func (dbContainer) openClient(driverName string, dataSourceName string) (*ent.Client, error) {
	drv, err := sql.Open(driverName, dataSourceName)
	if err != nil {
		return nil, err
	}
	// Get the underlying sql.DB object of the driver.
	db := drv.DB()
	err = db.Ping()
	if err != nil {
		_ = drv.Close()
		return nil, err
	}
	//c.Drv = drv
	//db.SetMaxIdleConns(10)
	//db.SetMaxOpenConns(100)
	//db.SetConnMaxLifetime(time.Hour)
	return ent.NewClient(ent.Driver(drv)), nil
}

// 连接到数据库并返回连接的客户端
func (that *dbContainer) connectDb() *ent.Client {
	key := that.key
	connection, ok := config.Get().Db.Connections[key]
	if !ok {
		//Int64String int64 json:",string" // “Int64String”:“0”
		log.Panicln(fmt.Sprintf("数据库[%s]配置未被找到", key))
	}
	driver := strings.ToLower(connection.Driver)
	var client *ent.Client
	var err error
	switch {
	case driver == dialect.MySQL || driver == "mariadb":
		// [username[:password]@][protocol[(address)]]/dbname[?param1=value1&...&paramN=valueN]
		dataSourceStr := ""
		if connection.Username != "" {
			dataSourceStr = connection.Username
			if connection.Password != "" {
				dataSourceStr = fmt.Sprintf("%s:%s", dataSourceStr, connection.Password)
			}
			dataSourceStr = fmt.Sprintf("%s@", dataSourceStr)
		}
		if connection.Host != "" {
			if connection.Protocol == "" {
				connection.Protocol = "tcp"
			}
			fullHost := connection.Host
			if connection.Port > 0 {
				fullHost = fmt.Sprintf("%s:%d", connection.Host, connection.Port)
			}
			dataSourceStr = fmt.Sprintf("%s%s(%s)", dataSourceStr, connection.Protocol, fullHost)
		}
		dataSourceStr = fmt.Sprintf("%s/%s", dataSourceStr, connection.Dbname)
		// 附加额外的参数
		if connection.Extra != nil && len(connection.Extra) > 0 {
			dataSourceStr += "?" + utils.MapToQueryString(connection.Extra)
		}
		client, err = that.openClient(dialect.MySQL, dataSourceStr)
	case driver == "postgresql" || driver == dialect.Postgres:
		dataSourceStr := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable",
			connection.Host, connection.Port, connection.Username, connection.Password, connection.Dbname)
		client, err = that.openClient(dialect.Postgres, dataSourceStr)
	case driver == "sqlite" || driver == dialect.SQLite:
		client, err = that.openClient(dialect.SQLite, connection.Database)
	default:
		log.Panicln("未定义的连接方式")
	}
	if err != nil {
		log.Panicln(fmt.Errorf("connection DB(%s:%s) error：%s", connection.Driver, connection.Host, err))
	}
	that.isConnect = true
	that.Client = client
	return client
}

// 关闭数据库连接
func (that dbContainer) close() error {
	if that.isConnect {
		err := that.Client.Close()
		if err != nil {
			return err
		}
	}
	return nil
}
// 重启数据库连接
func (that *dbContainer) reload() error {
	err := that.close()
	if err != nil {
		return err
	}
	client := that.connectDb()
	that.Client = client
	return nil
}