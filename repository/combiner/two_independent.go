package combiner

import (
	"github.com/xbox1994/gpb/repository/creater"
	"github.com/xbox1994/gpb/repository/loginer"
	"github.com/xbox1994/gpb/repository/model"
)

type TwoIndependentCombiner struct {
}

func (r TwoIndependentCombiner) CreateAndCombine(
	repoCreator creater.RepoCreator,
	repoCreatePreInfo loginer.RepoCreatePreInfo,
	answers model.Answer) {
	// 在远端与本地同时创建所有Repo
	mainRepoName := answers.RepoName
	answers.RepoName = mainRepoName + "-admin"
	createRepo(repoCreator, repoCreatePreInfo, answers, mainRepoName)
	answers.RepoName = mainRepoName + "-server"
	createRepo(repoCreator, repoCreatePreInfo, answers, mainRepoName)
}
