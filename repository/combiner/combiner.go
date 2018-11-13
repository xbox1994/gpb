package combiner

import (
	"github.com/xbox1994/gpb/repository/creater"
	"github.com/xbox1994/gpb/repository/loginer"
	"github.com/xbox1994/gpb/repository/model"
)

type RepoCombiner interface {
	CreateAndCombine(repoCreator creater.RepoCreator, repoCreatePreInfo loginer.RepoCreatePreInfo, answers model.Answer)
}
