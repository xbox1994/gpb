package combiner

import (
	"github.com/xbox1994/wps-gpb/repository/creater"
	"github.com/xbox1994/wps-gpb/repository/loginer"
	"github.com/xbox1994/wps-gpb/repository/model"
)

type RepoCombiner interface {
	CreateAndCombine(repoCreator creater.RepoCreator, repoCreatePreInfo loginer.RepoCreatePreInfo, answers model.Answer)
}
