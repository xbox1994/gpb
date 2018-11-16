package beegocli

import (
	"fmt"
	"github.com/xbox1994/gpb/common/util"
	"github.com/xbox1994/gpb/project/beegocli/template"
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
	} else {
		panic(err)
	}
	return
}
