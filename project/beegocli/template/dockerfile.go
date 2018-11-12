package template

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

var dockerfile = &Template{
	Mode:     Create,
	FilePath: `Dockerfile`,
	Content: `FROM alpine:latest
RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*

WORKDIR /app
ADD {{.appName}} .
ADD conf conf
RUN chmod +x {{.appName}}

ENTRYPOINT [ "./{{.appName}}" ]`,
	StdOut: createDockerfile,
}

func init() {
	AvailableTemplates = append(AvailableTemplates, dockerfile)
}

func createDockerfile(template *Template, args ...string) (err error) {
	if len(args) < 2 {
		err = errors.New(`params error`)
	} else {
		projectPath := args[0]
		projectName := args[1]
		absPath := fmt.Sprintf("%s/%s/", projectPath, projectName)
		content := strings.Replace(template.Content, `{{.appName}}`, projectName, -1)
		if file, err1 := os.Create(absPath + template.FilePath); err1 == nil {
			_, err = file.Write([]byte(content))
			file.Close()
		} else {
			err = err1
			fmt.Println(err)
		}
	}
	return
}
