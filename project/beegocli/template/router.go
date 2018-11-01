package template

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

var router = &Template{
	Mode:     Modify,
	FilePath: `/routers/router.go`,
	Content: `// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"{{.packagePath}}/controllers"

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
`,
	StdOut: writeRouter,
}

func init() {
	AvailableTemplates = append(AvailableTemplates, router)
}

func writeRouter(template *Template, args ...string) (err error) {
	if len(args) < 2 {
		err = errors.New(`params error`)
	} else {
		fmt.Println("writeRouter")
		projectPath := args[0]
		projectName := args[1]
		absPath := fmt.Sprintf("%s/%s", projectPath, projectName)
		packagePath := getPackagePath(absPath)
		content := strings.Replace(template.Content, `{{.packagePath}}`, packagePath, -1)
		if file, err1 := os.Create(absPath + template.FilePath); err1 == nil {
			_, err = file.Write([]byte(content))
			fmt.Println("writeRouter over")
			file.Close()
		} else {
			err = err1
			fmt.Println(err)
		}
	}
	return
}
