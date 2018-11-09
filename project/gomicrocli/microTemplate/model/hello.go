package model

import (
	"github.com/astaxie/beego/orm"
)

type Hello struct {
	Id   int64
	Name string
}

func init() {
	orm.RegisterModel(new(Hello))
}
