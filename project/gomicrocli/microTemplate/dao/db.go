package dao

import (
	"common/log"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

const (
	DATABASE      = "microTemplate"
	MAX_IDLE_CONN = 10
	MAX_OPEN_CONN = 100
)

func init() {
	user := beego.AppConfig.String("db_user")
	pwd := beego.AppConfig.String("db_pwd")
	host := beego.AppConfig.String("db_host")
	log.Info(nil, host)

	url := fmt.Sprintf("%s:%s@%s/%s?charset=utf8", user, pwd, host, DATABASE)
	orm.RegisterDataBase("default", "mysql", url, MAX_IDLE_CONN, MAX_OPEN_CONN)
	orm.RegisterDataBase(DATABASE, "mysql", url, MAX_IDLE_CONN, MAX_OPEN_CONN)
}
