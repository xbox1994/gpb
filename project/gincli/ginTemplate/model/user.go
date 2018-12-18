package model

import (
	"github.com/astaxie/beego/orm"
)

type UserInfo struct {
	Id   int64 `json:"id"`
	Name string `json:"name"`
}

func init() {
	orm.RegisterModel(new(UserInfo))
}
