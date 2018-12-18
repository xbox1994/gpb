package dao

import (
	"ginTemplate/model"
)

type HelloDao struct {
}

func NewHelloDAO() *HelloDao {
	return &HelloDao{}
}

func (r *HelloDao) GetInfo() (*model.UserInfo, error) {
	//c := &model.UserInfo{}
	//o := orm.NewOrm()
	//o.Using("default")

	//err := o.QueryTable("").One(c)
	//if err != nil {
	//	if err == orm.ErrNoRows {
	//		return nil, nil
	//	}
	//	return nil, err
	//}
	//return c, nil

	return &model.UserInfo{
		Id:   1,
		Name: "test",
	}, nil
}
