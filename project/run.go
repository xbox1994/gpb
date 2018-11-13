package project

import (
	"github.com/xbox1994/gpb/common/project_type"
	"github.com/xbox1994/gpb/project/beegocli"
	"github.com/xbox1994/gpb/project/gomicrocli"
	"gopkg.in/AlecAivazis/survey.v1"
	"os"
)

func CreateProjects(projectStructure string, projectName string) (err error) {
	switch projectStructure {
	case project_type.OneIndependent:
		err = selectAndGenerateProject(projectName, "", ".")
	case project_type.TwoIndependent:
		err = selectAndGenerateProject(projectName, "-admin", ".")
		err = selectAndGenerateProject(projectName, "-server", ".")
	case project_type.TwoIndependentWithParent:
		os.Mkdir(projectName, os.ModeDir)
		err = selectAndGenerateProject(projectName, "-admin", projectName)
		err = selectAndGenerateProject(projectName, "-server", projectName)
	}
	return err
}

func selectAndGenerateProject(projectName string, projectSuffix string, path string) error {
	projectFramework := ""
	projectStructureArray := []string{project_type.Beego, project_type.GoMicro}
	err := survey.AskOne(
		&survey.Select{
			Message: "Select " + projectSuffix + " project framework:",
			Options: projectStructureArray,
		},
		&projectFramework, nil)
	if err != nil {
		panic(err)
	}
	switch projectFramework {
	case project_type.Beego:
		err = beegocli.CreateProject(projectName+projectSuffix, path)
	case project_type.GoMicro:
		err = gomicrocli.CreateProject(projectName+projectSuffix, path)
	}
	return err
}
