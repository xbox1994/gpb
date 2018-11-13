package creater

import (
	"wps-gpb/repository/loginer"
	"wps-gpb/repository/model"
)

type RepoCreator interface {
	CreateRemoteRepo(answer model.Answer, repoCreatePreInfo loginer.RepoCreatePreInfo)
}
