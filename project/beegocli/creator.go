package beegocli

import (
	"fmt"
	"grb/common/project_type"
	"grb/project/beegocli/template"
	"grb/util"
	"os"
	"os/exec"
)

func CheckBeeExisted() bool {
	cmd := exec.Command("bee", "version")
	if err := cmd.Run(); err != nil {
		fmt.Println(err)
		return false
	}
	return true
}

func InstallBee() (err error) {
	cmd := exec.Command("go", "get", "github.com/beego/bee")
	err = cmd.Run()
	return
}

func CreateProject(projectName string, path string) (err error) {
	defer func() {
		if err != nil {
			fmt.Println(err)
		}
	}()
	if beeExisted := CheckBeeExisted(); beeExisted == false {
		if err = InstallBee(); err != nil {
			return
		}
	}
	if err = util.Run(exec.Command("bee", "api", projectName), path); err == nil {
		for _, template := range template.AvailableTemplates {
			if err = template.StdOut(template, path, projectName); err != nil {
				break
			}
		}
	}
	return
}

func CreateProjects(projectStructure string, projectName string) (err error) {
	switch projectStructure {
	case project_type.OneIndependent:
		err = CreateProject(projectName, "")
	case project_type.TwoIndependent:
		err = CreateProject(projectName+"-admin", "")
		err = CreateProject(projectName+"-server", "")
	case project_type.TwoIndependentWithParent:
		os.Mkdir(projectName, os.ModeDir)
		err = CreateProject(projectName+"-admin", projectName)
		err = CreateProject(projectName+"-server", projectName)
	}
	return err
}
