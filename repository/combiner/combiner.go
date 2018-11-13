package combiner

import (
	"wps-gpb/repository/creater"
	"wps-gpb/repository/loginer"
	"wps-gpb/repository/model"
)

type RepoCombiner interface {
	CreateAndCombine(repoCreator creater.RepoCreator, repoCreatePreInfo loginer.RepoCreatePreInfo, answers model.Answer)
}
