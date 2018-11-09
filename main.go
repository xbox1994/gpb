package main

import (
	"gopkg.in/AlecAivazis/survey.v1"
	"grb/common/project_type"
	"grb/project"
	"grb/repository"
)

func main() {
	projectStructure := ""
	projectStructureArray := []string{
		project_type.OneIndependent,
		project_type.TwoIndependent,
		project_type.TwoIndependentWithParent}
	err := survey.AskOne(
		&survey.Select{
			Message: "Select project structure you want to create:",
			Options: projectStructureArray,
		},
		&projectStructure, nil)
	if err != nil {
		panic(err)
	}

	isCreateRemoteAndLocalRepository := false
	err = survey.AskOne(
		&survey.Confirm{Message: "Create remote and local git repository?"},
		&isCreateRemoteAndLocalRepository, nil)
	if err != nil {
		panic(err)
	}

	projectName := ""
	if isCreateRemoteAndLocalRepository {
		projectName = repository.Create(projectStructure)
	} else {
		err := survey.AskOne(
			&survey.Input{Message: "Project Name:"},
			&projectName, nil)
		if err != nil {
			panic(err)
		}
	}
	project.CreateProjects(projectStructure, projectName)
}
