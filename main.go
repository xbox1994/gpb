package main

import (
	"gopkg.in/AlecAivazis/survey.v1"
	"grb/project/beegocli"
	"grb/repository"
)

func main() {
	isCreateRemoteAndLocalRepository := false
	err := survey.AskOne(
		&survey.Confirm{Message: "Create remote and local repository?"},
		&isCreateRemoteAndLocalRepository, nil)
	if err != nil {
		panic(err)
	}

	projectName := ""
	if isCreateRemoteAndLocalRepository {
		projectName = repository.Create()
		beegocli.CreateProject(projectName+"-server", projectName)
		beegocli.CreateProject(projectName+"-admin", projectName)
	} else {
		err := survey.AskOne(
			&survey.Input{Message: "Project Name:"},
			&projectName, nil)
		if err != nil {
			panic(err)
		}
		beegocli.CreateProject(projectName, "")
	}
}
