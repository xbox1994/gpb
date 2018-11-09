# microTemplate

#### gRPC文件编译
* protoc(https://github.com/protocolbuffers/protobuf/releases)
* protoc-gen-go(`go get -u github.com/golang/protobuf/protoc-gen-go`)
* protoc-gen-micro(`go get -u github.com/micro/protoc-gen-micro`)
* go-api(`go get -u github.com/micro/go-api`)

编译proto文件夹中的hello.proto文件：

`protoc --proto_path=proto --proto_path=%GOPATH%\src --micro_out=proto --go_out=proto hello.proto`

#### 服务注册中心
* consul(https://learn.hashicorp.com/consul/getting-started/install.html)

使用方式： `consul agent -dev`

#### api网关
* micro(`go get -u github.com/micro/micro`)

使用方式： `micro api`

#### 本应用
安装依赖：`dep ensure`

启动方式：`go run main.go`

测试方式：
```bash
curl -X POST \
  'http://localhost:8080/hello/hello' \
  -H 'content-type: application/json' \
  -d '{
	"name": "xxx"
}'
```
