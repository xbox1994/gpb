package combiner

import (
	"fmt"
	"grb/common/util"
	"grb/repository/creater"
	"grb/repository/loginer"
	"grb/repository/model"
	"io/ioutil"
	"net/url"
	"os"
	"os/exec"
)

type SingleCombiner struct {
}

func (r SingleCombiner) CreateAndCombine(
	repoCreator creater.RepoCreator,
	repoCreatePreInfo loginer.RepoCreatePreInfo,
	answers model.Answer) {
	mainRepoName := answers.RepoName
	createRepo(repoCreator, repoCreatePreInfo, answers, mainRepoName)
}

func createRepo(
	repoCreator creater.RepoCreator,
	repoCreatePreInfo loginer.RepoCreatePreInfo,
	answers model.Answer,
	mainRepoName string) {
	var repoFolderName string
	repoFolderName = answers.RepoName
	os.Mkdir(repoFolderName, os.ModeDir)
	fmt.Println("create README file for " + answers.RepoName)
	repoCreator.CreateRemoteRepo(answers, repoCreatePreInfo)
	util.Run(exec.Command("git", "init"), repoFolderName)
	ioutil.WriteFile(repoFolderName+"/README", []byte(""), 0644)
	util.Run(exec.Command("git", "add", "."), repoFolderName)
	util.Run(exec.Command("git", "commit", "-m", "\"init\""), repoFolderName)
	parse, _ := url.Parse(answers.GitHostAddress)
	gitRepoPath := "git@" + parse.Host + ":" + answers.RepoNamespace + "/" + answers.RepoName + ".git"
	util.Run(exec.Command("git", "remote", "add", "origin", gitRepoPath), repoFolderName)
	util.Run(exec.Command("git", "push", "-u", "origin", "master"), repoFolderName)
	util.Run(exec.Command("git", "submodule", "add", gitRepoPath, repoFolderName), mainRepoName)
}
