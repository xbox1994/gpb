package creator

import "grb/repository/model"

type RepoCreator interface {
	Login(loginInfo model.LoginInfo)
	CreateRepo(answer model.Answer)
}
