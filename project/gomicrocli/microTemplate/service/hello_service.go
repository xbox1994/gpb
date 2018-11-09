package service

import (
	"microTemplate/dao"
	"microTemplate/request"
)

func Do(request *request.HelloRequest) (string, error) {
	dao.NewHelloDAO().GetInfo()
	return request.Name, nil
}
