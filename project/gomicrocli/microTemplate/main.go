package main

import (
	"common/log"
	"context"
	"encoding/json"
	"fmt"
	"github.com/micro/go-api/proto"
	"github.com/micro/go-micro"
	"microTemplate/proto"
	"microTemplate/request"
	"microTemplate/service"
	"os"
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
	service := micro.NewService(
		micro.Name("go.micro.api.hello"),
	)
	service.Init()
	hello.RegisterHelloHandler(service.Server(), &Say{})

	if err := service.Run(); err != nil {
		log.Error(nil, err)
		os.Exit(1)
	}
}
