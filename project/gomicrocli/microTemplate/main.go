package main

import (
	"common/log"
	"context"
	"encoding/json"
	"fmt"
	"github.com/hashicorp/consul/api"
	"github.com/micro/go-api/proto"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry/consul"
	"microTemplate/proto"
	"microTemplate/request"
	"microTemplate/service"
	"os"
	"time"
)

type Say struct {
	Client hello.HelloService
}

func (s *Say) Hello(ctx context.Context, req *go_api.Request, rsp *go_api.Response) error {
	var helloRequest request.HelloRequest
	json.Unmarshal([]byte(req.Body), &helloRequest)

	result, codeE := service.Do(&helloRequest)
	if codeE != nil {
		rsp.StatusCode = 500
		errBody, _ := json.Marshal(request.HelloResponse{
			Code:    0,
			Message: fmt.Sprintf("%v", codeE),
		})
		rsp.Body = string(errBody)
		return nil
	}

	rsp.StatusCode = 200
	errBody, _ := json.Marshal(request.HelloResponse{
		Code:    0,
		Message: result,
	})
	rsp.Body = string(errBody)
	return nil
}

func main() {
	config := api.DefaultConfig()
	config.Address = os.Getenv("MICRO_REGISTRY_ADDRESS_CUSTOM")
	if config.Address == "" {
		config.Address = "127.0.0.1:8500"
	}
	service := micro.NewService(
		micro.Name("go.micro.api.hello"),
		micro.Registry(consul.NewRegistry(consul.Config(config), consul.TCPCheck(time.Second))),
	)
	service.Init()
	hello.RegisterHelloHandler(service.Server(), &Say{})

	if err := service.Run(); err != nil {
		log.Error(nil, err)
		os.Exit(1)
	}
}
