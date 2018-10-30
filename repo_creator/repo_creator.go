package repo_creator

import "grb/model"

type RepoCreator interface {
	CreateRepo(answer model.Answer)
}
