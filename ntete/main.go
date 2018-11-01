package main

import (
	"fmt"
	_ "grb/ntete/routers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"moservice/lib/metrics/yunsrv/beego"
)

func init() {
	beego.SetLogFuncCall(true)
	beego.SetLogger("console", fmt.Sprintf("{\"level\":%d}", logs.LevelDebug))
	level, err := beego.AppConfig.Int("log_level")
	if err != nil {
		level = 4
	}
	// warn 4  debug 7
	beego.SetLevel(level)
}

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}

	beego_filter.InsertMetricsFilter()

	beego.Run()
}
