package template

import (
	"errors"
	"fmt"
	"os"
)

// 珠海vender lib目录下bee工具自动生成的docker目录删除掉
var dockerDirConf = &Template{
	Mode:     Del,
	FilePath: `docker/`,
	StdOut:   delDockerDir,
}

func init() {
	AvailableTemplates = append(AvailableTemplates, dockerDirConf)
}

func delDockerDir(template *Template, args ...string) (err error) {
	if len(args) < 2 {
		err = errors.New(`params error`)
	} else {
		projectPath := args[0]
		projectName := args[1]
		absPath := fmt.Sprintf("%s/%s/%s", projectPath, projectName, template.FilePath)

		if file, err1 := os.Stat(absPath); err1 != nil {
			err = err1
		} else {
			if file.IsDir() {
				err = os.RemoveAll(absPath)
			} else {
				err = os.Remove(absPath)
			}
		}
	}
	return
}
