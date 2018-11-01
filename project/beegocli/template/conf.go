package template

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

var appConf = &Template{
	Mode:     Modify,
	FilePath: `conf/app.conf`,
	Content: `appname = {{.Appname}}
httpport = 8080
runmode = "${ENV_MODE||dev}"
copyrequestbody = true

[dev]
# swagger文档
EnableDocs = true

# mysql配置
db_user = dev
db_pwd = 12343
db_host = tcp(127.0.0.1:3306)
db_idle_conns=80
db_max_open_conns=80

# redis配置
redis.host = 127.0.0.1:6379
redis.db = 0

[prod]
# swagger文档
EnableDocs = true

# mysql配置
db_user = prod
db_pwd = 12343
db_host = tcp(127.0.0.1:3306)
db_idle_conns=80
db_max_open_conns=80

# redis配置
redis.host = 127.0.0.1:6379
redis.db = 0
`,
	StdOut: writeAppConf,
}

func init() {
	AvailableTemplates = append(AvailableTemplates, appConf)
}

func writeAppConf(template *Template, args ...string) (err error) {
	if len(args) < 2 {
		err = errors.New(`params error`)
	} else {
		projectPath := args[0]
		projectName := args[1]
		absPath := fmt.Sprintf("%s/%s/%s", projectPath, projectName, template.FilePath)
		content := strings.Replace(template.Content, `{{.Appname}}`, projectName, -1)
		if file, err1 := os.Create(absPath); err1 == nil {
			_, err = file.Write([]byte(content))
			file.Close()
		} else {
			err = err1
		}
	}

	return
}
