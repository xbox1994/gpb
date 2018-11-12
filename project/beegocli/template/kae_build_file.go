package template

import (
	"errors"
	"fmt"
	"os"
)

// 珠海vender lib目录下bee工具自动生成的kae_build.sh删除掉
var kaeBuildFile = &Template{
	Mode:     Del,
	FilePath: `kae_build.sh`,
	StdOut:   delKaeBuildFile,
}

func init() {
	AvailableTemplates = append(AvailableTemplates, kaeBuildFile)
}

func delKaeBuildFile(template *Template, args ...string) (err error) {
	if len(args) < 2 {
		err = errors.New(`params error`)
	} else {
		projectPath := args[0]
		projectName := args[1]
		absPath := fmt.Sprintf("%s/%s/%s", projectPath, projectName, template.FilePath)
		if file, err1 := os.Stat(absPath); err1 != nil {
			if !os.IsNotExist(err1) {
				err = err1
			}
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
