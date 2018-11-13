package creater

import (
	"github.com/xbox1994/gpb/repository/loginer"
	"github.com/xbox1994/gpb/repository/model"
)

type RepoCreator interface {
	CreateRemoteRepo(answer model.Answer, repoCreatePreInfo loginer.RepoCreatePreInfo)
}
