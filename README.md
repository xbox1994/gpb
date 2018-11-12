# grb
go repo builder

# 运行方式
1. 依赖：Golang、Git，将当前用户的ssh key保存到git远端服务器
2. 在本项目根目录下执行`go run main.go`
3. 将生成的项目移动到$GOPATH/src
4. 按照项目内的README文件安装依赖，运行应用

# 参考输入
	GitHostAddress:   "http://wpsgit.kingsoft.net/",
	GitServerVersion: "GitLab 6.3.0 LDAP",
	RepoName:         "grbtest",
	RepoNamespace:    "galaxy",
	Username:         "wangtianyi1",
	Password:         "xxx",