package creator

import "grb/model"

type RepoCreator interface {
	Login(loginInfo model.LoginInfo)
	CreateRepo(answer model.Answer)
}
