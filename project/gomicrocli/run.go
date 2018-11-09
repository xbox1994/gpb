package main

import (
	"grb/common/util"
	"path/filepath"
)

func CreateProject(projectName string) (err error) {
	util.CopyDir("project/gomicrocli/microTemplate", projectName)
	err = filepath.Walk("ll", util.Replace("microTemplate", projectName))
	if err != nil {
		panic(err)
	}
	return err
}
