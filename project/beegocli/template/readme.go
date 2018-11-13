package template

import (
	"errors"
	"fmt"
	"os"
)

var readmeConfig = &Template{
	Mode:     Create,
	FilePath: `README.md`,
	Content:  `建议使用dep工具执行"dep init"来添加项目依赖，并将vendor目录提交到git`,
	StdOut:   createReadmeConfig,
}

func init() {
	AvailableTemplates = append(AvailableTemplates, readmeConfig)
}

func createReadmeConfig(template *Template, args ...string) (err error) {
	if len(args) < 2 {
		err = errors.New(`params error`)
	} else {
		projectPath := args[0]
		projectName := args[1]
		absPath := fmt.Sprintf("%s/%s/", projectPath, projectName)
		if file, err1 := os.Create(absPath + template.FilePath); err1 == nil {
			_, err = file.Write([]byte(template.Content))
			file.Close()
		} else {
			err = err1
			fmt.Println(err)
		}
	}
	return
}
