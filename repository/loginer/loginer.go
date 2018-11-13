package loginer

import "github.com/xbox1994/gpb/repository/model"

type RepoCreatePreInfo struct {
	Cookie          string
	RepoGroupNameId string
}

type GitWebInterfaceLoginer interface {
	Login(loginInfo model.LoginInfo) RepoCreatePreInfo
}
