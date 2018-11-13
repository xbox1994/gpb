package loginer

import "wps-gpb/repository/model"

type RepoCreatePreInfo struct {
	Cookie          string
	RepoNamespaceId string
}

type GitWebInterfaceLoginer interface {
	Login(loginInfo model.LoginInfo) RepoCreatePreInfo
}
