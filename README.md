# grb
go repo builder

# 运行方式
1. 安装工具：Golang、Git、Dep
2. 安装项目依赖：`dep ensure`
3. 当前用户的ssh key已经保存到git远端服务器
4. 在本项目根目录下执行`go run main.go`
5. 将生成的项目移动到$GOPATH/src
6. 按照项目内的README文件安装依赖，运行应用

# 参考输入
	GitHostAddress:   "http://wpsgit.kingsoft.net/",
	GitServerVersion: "GitLab 6.3.0 LDAP",
	RepoName:         "grbtest",
	RepoNamespace:    "galaxy",
	Username:         "wangtianyi1",
	Password:         "xxx",