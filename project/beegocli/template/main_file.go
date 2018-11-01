package template

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

var mainFile = &Template{
	Mode:     Modify,
	FilePath: `/main.go`,
	Content: `package main

import (
	"fmt"
	_ "{{.packagePath}}/routers"

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
	logs.SetLogger(logs.AdapterConsole, ` + "`{\"level\":7}`)" + `
	//logs.SetLogger(logs.AdapterFile, ` + "`{\"filename\":\"test.log\",\"level\":6}`)" + `
	//logs.SetLogger(logs.AdapterMultiFile, ` + "`{\"filename\":\"test.log\", \"maxDays\":15, \"separate\":[\"emergency\", \"alert\", \"critical\", \"error\", \"warning\", \"notice\", \"info\", \"debug\"],}`)" + ` 
	//logs.SetLogger(logs.AdapterMail, ` + "`{\"username\":\"beegotest@163.com\", \"fromAddress\":\"beegotest@163.com\",\"password\":\"xxxxxxxx\",\"host\":\"smtp.163.com:25\",\"sendTos\":[\"user1@qq.com\",\"user2@qq.com\"],\"level\":4}`)" + `
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
`,
	StdOut: writeMainFile,
}

func init() {
	AvailableTemplates = append(AvailableTemplates, mainFile)
}

func writeMainFile(template *Template, args ...string) (err error) {
	if len(args) < 2 {
		err = errors.New(`params error`)
	} else {
		projectPath := args[0]
		projectName := args[1]
		absPath := fmt.Sprintf("%s/%s", projectPath, projectName)
		packagePath := getPackagePath(absPath)
		content := strings.Replace(template.Content, `{{.packagePath}}`, packagePath, -1)
		if file, err1 := os.Create(absPath + template.FilePath); err1 == nil {
			_, err = file.Write([]byte(content))
			file.Close()
		} else {
			err = err1
		}
	}

	return
}
