package repo_combiner

import (
	"fmt"
	"grb/model"
	"grb/repo_creator"
	"io/ioutil"
	"log"
	"net/url"
	"os"
	"os/exec"
	"strings"
)

type RepoCombiner struct {
	RepoCreator repo_creator.RepoCreator
}

func (r RepoCombiner) CreateAndCombineRepo(answers model.Answer) {
	// 在远端与本地同时创建所有Repo
	mainRepoName := answers.RepoName
	os.Mkdir(mainRepoName, os.ModeDir)
	r.RepoCreator.CreateRepo(answers)
	run(exec.Command("git", "init"), mainRepoName)

	answers.RepoName = mainRepoName + "-admin"
	r.createSubRepo(answers, mainRepoName)
	answers.RepoName = mainRepoName + "-server"
	r.createSubRepo(answers, mainRepoName)
	if answers.IncludeCommon {
		answers.RepoName = mainRepoName + "-common"
		r.createSubRepo(answers, mainRepoName)
	}
	answers.RepoName = mainRepoName + "-vendor"
	r.createSubRepo(answers, mainRepoName)

	// 将子Repo加入到父Repo中
	run(exec.Command("git", "add", "."), mainRepoName)
	run(exec.Command("git", "commit", "-m", "\"init\""), mainRepoName)
	parse, _ := url.Parse(answers.GitHostAddress)
	gitRepoPath := "git@" + parse.Host + ":" + answers.RepoNamespace + "/" + mainRepoName + ".git"
	run(exec.Command("git", "remote", "add", "origin", gitRepoPath), mainRepoName)
	run(exec.Command("git", "push", "-u", "origin", "master"), mainRepoName)
}

func (r RepoCombiner) createSubRepo(subRepoAnswers model.Answer, mainRepoName string) {
	var subRepoFolderName string
	if strings.HasSuffix(subRepoAnswers.RepoName, "vendor") {
		subRepoFolderName = "vendor"
	} else {
		subRepoFolderName = subRepoAnswers.RepoName
	}
	subRepoFolderPath := mainRepoName + "/" + subRepoFolderName
	os.Mkdir(subRepoFolderPath, os.ModeDir)
	log.Println("create README file for " + subRepoAnswers.RepoName)
	r.RepoCreator.CreateRepo(subRepoAnswers)
	run(exec.Command("git", "init"), subRepoFolderPath)
	ioutil.WriteFile(subRepoFolderPath+"/README", []byte(""), 0644)
	run(exec.Command("git", "add", "."), subRepoFolderPath)
	run(exec.Command("git", "commit", "-m", "\"init\""), subRepoFolderPath)
	parse, _ := url.Parse(subRepoAnswers.GitHostAddress)
	gitRepoPath := "git@" + parse.Host + ":" + subRepoAnswers.RepoNamespace + "/" + subRepoAnswers.RepoName + ".git"
	run(exec.Command("git", "remote", "add", "origin", gitRepoPath), subRepoFolderPath)
	run(exec.Command("git", "push", "-u", "origin", "master"), subRepoFolderPath)
	run(exec.Command("git", "submodule", "add", gitRepoPath, subRepoFolderName), mainRepoName)
}

func run(cmd *exec.Cmd, dir string) {
	fmt.Println(cmd.Args)
	cmd.Dir = dir
	e := cmd.Run()
	if e != nil {
		fmt.Println(e)
	}
}
