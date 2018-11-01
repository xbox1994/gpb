package main

import (
	"fmt"
	_ "grb/testcli/routers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
)

func initMysql() {
	name := beego.AppConfig.String("db_name")
	user := beego.AppConfig.String("db_user")
	pwd := beego.AppConfig.String("db_pwd")
	host := beego.AppConfig.String("db_host")
	idleConns := beego.AppConfig.DefaultInt("db_idle_conns", 1024)
	maxOpenConns := beego.AppConfig.DefaultInt("db_max_open_conns", 1024)

	url := fmt.Sprintf("%s:%s@%s/%s?charset=utf8", user, pwd, host, name)
	//注册数据库
	orm.RegisterDataBase("default", "mysql", url, idleConns, maxOpenConns)
}

func initLogs() {
	logs.SetLogger(logs.AdapterConsole, `{"level":7}`)
	//logs.SetLogger(logs.AdapterFile, `{"filename":"test.log","level":6}`)
	//logs.SetLogger(logs.AdapterMultiFile, `{"filename":"test.log", "maxDays":15, "separate":["emergency", "alert", "critical", "error", "warning", "notice", "info", "debug"],}`) 
	//logs.SetLogger(logs.AdapterMail, `{"username":"beegotest@163.com", "fromAddress":"beegotest@163.com","password":"xxxxxxxx","host":"smtp.163.com:25","sendTos":["user1@qq.com","user2@qq.com"],"level":4}`)
}

func initSwagger() {
	enableDocs, _ := beego.AppConfig.Bool("EnableDocs")
	if enableDocs == true {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
}

func init() {
	initMysql()
	initLogs()
	initSwagger()
}

func main() {
	beego.Run()
}
