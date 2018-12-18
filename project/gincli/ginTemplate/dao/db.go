package dao

import (
	"common/log"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

const (
	DatabaseName = "ginTemplate"
)

func init() {
	user := beego.AppConfig.String("db_user")
	pwd := beego.AppConfig.String("db_pwd")
	host := beego.AppConfig.String("db_host")
	idleConns, _ := beego.AppConfig.Int("db_idle_conns")
	maxOpenConns, _ := beego.AppConfig.Int("db_max_open_conns")
	log.Info(nil, user, host, idleConns, maxOpenConns)

	url := fmt.Sprintf("%s:%s@%s/%s?charset=utf8", user, pwd, host, DatabaseName)
	orm.RegisterDataBase("default", "mysql", url, idleConns, maxOpenConns)
	orm.RegisterDataBase(DatabaseName, "mysql", url, idleConns, maxOpenConns)
}
