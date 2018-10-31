package model

type Answer struct {
	GitHostAddress   string
	GitServerVersion string
	RepoName         string
	RepoNamespace    string
	IncludeCommon    bool
	Username         string
	Password         string
}
