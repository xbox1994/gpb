package creater

import (
	"grb/repository/loginer"
	"grb/repository/model"
)

type RepoCreator interface {
	CreateRemoteRepo(answer model.Answer, repoCreatePreInfo loginer.RepoCreatePreInfo)
}
