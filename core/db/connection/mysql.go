package connection

import (
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"fmt"
	"github.com/boshangad/go-api/ent"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"strconv"
	"strings"
)

type mysqlConnection struct {
	ConnectionInterface
	Params
}

func (that *mysqlConnection) Open() *ent.Client {
	// [username[:password]@][protocol[(address)]]/dbname[?param1=value1&...&paramN=valueN]
	var (
		dbDsn string = ""
		dbAccount string = ""
		dbProtocol string = ""
		dbName string = that.Dbname
		dbParamsStr string = ""
	)
	if that.Username != "" {
		password := ""
		if that.Password != "" {
			password = ":" + that.Password
		}
		dbAccount = fmt.Sprintf("%s%s@", that.Username, password)
	}
	if that.Protocol != "" {
		address := ""
		if that.Host != "" {
			address = "(" + that.Host + ")"
		}
		dbProtocol = fmt.Sprintf("%s%s", that.Protocol, address)
	}
	// 附加额外的参数
	if that.Extras != nil && len(that.Extras) > 0 {
		dbParamsStr = "?"
		for key, value := range that.Extras {
			v := ""
			if v1, ok := value.(string); ok {
				v = strings.TrimSpace(v1)
			} else if v1, ok := value.(bool); ok {
				v = strconv.FormatBool(v1)
			} else if v1, ok := value.(int); ok {
				v = strconv.FormatInt(int64(v1), 10)
			} else if v1, ok := value.(int8); ok {
				v = strconv.FormatInt(int64(v1), 10)
			} else if v1, ok := value.(int16); ok {
				v = strconv.FormatInt(int64(v1), 10)
			} else if v1, ok := value.(int32); ok {
				v = strconv.FormatInt(int64(v1), 10)
			} else if v1, ok := value.(int64); ok {
				v = strconv.FormatInt(v1, 10)
			} else if v1, ok := value.(uint); ok {
				v = strconv.FormatUint(uint64(v1), 10)
			} else if v1, ok := value.(uint8); ok {
				v = strconv.FormatUint(uint64(v1), 10)
			} else if v1, ok := value.(uint16); ok {
				v = strconv.FormatUint(uint64(v1), 10)
			} else if v1, ok := value.(uint32); ok {
				v = strconv.FormatUint(uint64(v1), 10)
			} else if v1, ok := value.(uint64); ok {
				v = strconv.FormatUint(v1, 10)
			} else {
				continue
			}
			dbParamsStr += strings.TrimSpace(key) + "=" + v + "&"
		}
		dbParamsStr = strings.TrimRight(dbParamsStr, "&")
	}
	// 格式化为连接
	dbDsn = fmt.Sprintf("%s%s/%s%s", dbAccount, dbProtocol, dbName, dbParamsStr)
	drv, err := sql.Open(dialect.MySQL, dbDsn)
	if err != nil {
		log.Panicln("mysql open fail", err, dbDsn)
		return nil
	}
	// Get the underlying sql.DB object of the driver.
	db := drv.DB()
	err = db.Ping()
	if err != nil {
		_ = drv.Close()
		log.Panicln("mysql connect fail", err)
		return nil
	}
	//c.Drv = drv
	//db.SetMaxIdleConns(10)
	//db.SetMaxOpenConns(100)
	//db.SetConnMaxLifetime(time.Hour)
	return ent.NewClient(ent.Driver(drv))
}

func NewMysqlConnect(data Params) *mysqlConnection {
	connect := mysqlConnection{}
	connect.Driver = data.Driver
	connect.Database = data.Database
	connect.Username = data.Username
	connect.Password = data.Password
	connect.Protocol = data.Protocol
	connect.Host = data.Host
	connect.Port = data.Port
	connect.Dbname = data.Dbname
	connect.Extras = data.Extras
	return &connect
}