package creater

import (
	"github.com/xbox1994/wps-gpb/repository/loginer"
	"github.com/xbox1994/wps-gpb/repository/model"
)

type RepoCreator interface {
	CreateRemoteRepo(answer model.Answer, repoCreatePreInfo loginer.RepoCreatePreInfo)
}
