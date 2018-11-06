package loginer

import "grb/repository/model"

type RepoCreatePreInfo struct {
	Cookie          string
	RepoNamespaceId string
}

type GitWebInterfaceLoginer interface {
	Login(loginInfo model.LoginInfo) RepoCreatePreInfo
}
