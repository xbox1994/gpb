// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"grb/testcli/controllers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/plugins/cors"
)

func initCors() {
	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowOrigins:     []string{"http://127.0.0.1:8082"},
		AllowMethods:     []string{"GET", "POST", "PUT", "OPTION"},
		AllowHeaders:     []string{"Access-Control-Allow-Origin", "Origin", "Accept", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Access-Control-Allow-Origin", "Origin"},
		AllowCredentials: true,
	}))
}

func initRoute() {
	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/object",
			beego.NSInclude(
				&controllers.ObjectController{},
			),
		),
		beego.NSNamespace("/user",
			beego.NSInclude(
				&controllers.UserController{},
			),
		),
	)
	beego.AddNamespace(ns)
}

func init() {
	initCors()
	initRoute()
}
