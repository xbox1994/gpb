package beegocli

import (
	"fmt"
	"grb/project/beegocli/template"
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

func CreateProject(projectName string) (err error) {
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
	cmd := exec.Command("bee", "api", projectName)
	if err = cmd.Run(); err == nil {
		if basePath, err := os.Getwd(); err == nil {
			for _, template := range template.AvailableTemplates {
				if err = template.StdOut(template, basePath, projectName); err != nil {
					break
				}
			}
		}
	}
	return
}
