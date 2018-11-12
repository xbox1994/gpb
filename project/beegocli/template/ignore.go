package template

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

var ignoreFile = &Template{
	Mode:     Create,
	FilePath: `.gitignore`,
	Content: `.idea
log
{{.appName}}
`,
	StdOut: createIgnorefile,
}

func init() {
	AvailableTemplates = append(AvailableTemplates, ignoreFile)
}

func createIgnorefile(template *Template, args ...string) (err error) {
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
