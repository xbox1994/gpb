package combiner

import (
	"grb/repository/creater"
	"grb/repository/loginer"
	"grb/repository/model"
)

type RepoCombiner interface {
	CreateAndCombine(repoCreator creater.RepoCreator, repoCreatePreInfo loginer.RepoCreatePreInfo, answers model.Answer)
}
