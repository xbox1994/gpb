package combiner

import (
	"fmt"
	"grb/repository/creator"
	"grb/repository/model"
	"grb/util"
	"io/ioutil"
	"net/url"
	"os"
	"os/exec"
)

type RepoCombiner struct {
	RepoCreator creator.RepoCreator
}

func (r RepoCombiner) CreateAndCombineRepo(answers model.Answer) {
	// 在远端与本地同时创建所有Repo
	mainRepoName := answers.RepoName
	os.Mkdir(mainRepoName, os.ModeDir)
	r.RepoCreator.CreateRepo(answers)
	util.Run(exec.Command("git", "init"), mainRepoName)

	answers.RepoName = mainRepoName + "-admin"
	r.createSubRepo(answers, mainRepoName)
	answers.RepoName = mainRepoName + "-server"
	r.createSubRepo(answers, mainRepoName)

	// 将子Repo加入到父Repo中
	util.Run(exec.Command("git", "add", "."), mainRepoName)
	util.Run(exec.Command("git", "commit", "-m", "\"init\""), mainRepoName)
	parse, _ := url.Parse(answers.GitHostAddress)
	gitRepoPath := "git@" + parse.Host + ":" + answers.RepoNamespace + "/" + mainRepoName + ".git"
	util.Run(exec.Command("git", "remote", "add", "origin", gitRepoPath), mainRepoName)
	util.Run(exec.Command("git", "push", "-u", "origin", "master"), mainRepoName)
}

func (r RepoCombiner) createSubRepo(subRepoAnswers model.Answer, mainRepoName string) {
	var subRepoFolderName string
	subRepoFolderName = subRepoAnswers.RepoName
	subRepoFolderPath := mainRepoName + "/" + subRepoFolderName
	os.Mkdir(subRepoFolderPath, os.ModeDir)
	fmt.Println("create README file for " + subRepoAnswers.RepoName)
	r.RepoCreator.CreateRepo(subRepoAnswers)
	util.Run(exec.Command("git", "init"), subRepoFolderPath)
	ioutil.WriteFile(subRepoFolderPath+"/README", []byte(""), 0644)
	util.Run(exec.Command("git", "add", "."), subRepoFolderPath)
	util.Run(exec.Command("git", "commit", "-m", "\"init\""), subRepoFolderPath)
	parse, _ := url.Parse(subRepoAnswers.GitHostAddress)
	gitRepoPath := "git@" + parse.Host + ":" + subRepoAnswers.RepoNamespace + "/" + subRepoAnswers.RepoName + ".git"
	util.Run(exec.Command("git", "remote", "add", "origin", gitRepoPath), subRepoFolderPath)
	util.Run(exec.Command("git", "push", "-u", "origin", "master"), subRepoFolderPath)
	util.Run(exec.Command("git", "submodule", "add", gitRepoPath, subRepoFolderName), mainRepoName)
}
