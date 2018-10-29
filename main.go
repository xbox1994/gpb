package main

import (
	"fmt"
	"gopkg.in/AlecAivazis/survey.v1"
)

// the questions to ask
var qs = []*survey.Question{
	{
		Name:     "gitHostAddress",
		Prompt:   &survey.Input{Message: "Git Host address:"},
		Validate: survey.Required,
	},
	{
		Name: "gitServerVersion",
		Prompt: &survey.Select{
			Message: "Git server version:",
			Options: []string{"GitLab 6.3.0 LDAP"},
			Default: "GitLab 6.3.0",
		},
	},
	{
		Name:     "repoName",
		Prompt:   &survey.Input{Message: "Main repository name:"},
		Validate: survey.Required,
	},
	{
		Name:     "username",
		Prompt:   &survey.Input{Message: "Git login username"},
		Validate: survey.Required,
	},
	{
		Name:     "password",
		Prompt:   &survey.Input{Message: "Git login password"},
		Validate: survey.Required,
	},
}

func main() {
	answers := struct {
		GitHostAddress   string
		GitServerVersion string
		RepoName         string
		Username         string
		Password         string
	}{}

	err := survey.Ask(qs, &answers)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(answers.GitHostAddress, answers.GitServerVersion, answers.RepoName, answers.Username, answers.Password)
}
