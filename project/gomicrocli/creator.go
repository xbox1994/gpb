package gomicrocli

import (
	"fmt"
	"grb/common/util"
	"path/filepath"
)

func CreateProject(projectName string, path string) (err error) {
	projectPath := projectName
	if path != "" {
		projectPath = path + "/" + projectName
	}

	err = util.CopyDir("project/gomicrocli/microTemplate", projectPath)
	if err != nil {
		panic(err)
	}
	err = filepath.Walk(projectPath, util.Replace("microTemplate", projectName))
	if err != nil {
		panic(err)
	}
	fmt.Printf("Create project %s success", projectName)
	return err
}
