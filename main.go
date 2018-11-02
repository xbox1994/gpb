package main

import (
	"errors"
	"fmt"
	"gopkg.in/AlecAivazis/survey.v1"
	"grb/model"
	"grb/repo_combiner"
	"grb/repo_creator"
	"os"
)

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
			Default: "GitLab 6.3.0 LDAP",
		},
	},
	{
		Name:     "repoName",
		Prompt:   &survey.Input{Message: "Main repository name:"},
		Validate: survey.Required,
	},
	{
		Name:     "repoNamespace",
		Prompt:   &survey.Input{Message: "Repository namespace:"},
		Validate: survey.Required,
	},
	{
		Name:     "includeCommon",
		Prompt:   &survey.Confirm{Message: "Create common repo?"},
		Validate: survey.Required,
	},
	{
		Name:     "username",
		Prompt:   &survey.Input{Message: "Git login username"},
		Validate: survey.Required,
	},
	{
		Name:     "password",
		Prompt:   &survey.Password{Message: "Git login password"},
		Validate: survey.Required,
	},
}

func main() {
	answers := model.Answer{}

	//err := survey.Ask(qs, &answers)
	//if err != nil {
	//	fmt.Println(err.Error())
	//	return
	//}

	// 测试数据
	answers = model.Answer{
		GitHostAddress:   "http://wpsgit.kingsoft.net/",
		GitServerVersion: "GitLab 6.3.0 LDAP",
		RepoName:         "grbtest",
		RepoNamespace:    "galaxy",
		Username:         "wangtianyi1",
		Password:         "",
		IncludeCommon:    false,
	}

	// 选择creator
	var repoCreator repo_creator.RepoCreator
	if "GitLab 6.3.0 LDAP" == answers.GitServerVersion {
		repoCreator = &repo_creator.Gitlab630Ldap{}
		repoCreator.Login(model.LoginInfo{
			GitHostAddress: answers.GitHostAddress,
			Username:       answers.Username,
			Password:       answers.Password,
		})
	} else {
		fmt.Println(errors.New(answers.GitServerVersion + " no implement yet"))
		os.Exit(1)
	}

	// 在远端与本地创建并合并子项目到父项目
	repo_combiner.RepoCombiner{
		RepoCreator: repoCreator,
	}.CreateAndCombineRepo(answers)

}
