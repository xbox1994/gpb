package service

import (
	"ginTemplate/dao"
	"ginTemplate/model"
	"github.com/gin-gonic/gin"
	"strconv"
)

type UserResource struct {
}

func (r *UserResource) GetAllUsers(c *gin.Context) {
	users, _ := dao.NewHelloDAO().GetInfo()
	c.JSON(200, users)
}

func (r *UserResource) GetUserByID(c *gin.Context) {
	id := c.Params.ByName("id")
	i, _ := strconv.Atoi(id)
	c.JSON(200, &model.UserInfo{
		Id:   int64(i),
		Name: "unknown",
	})
}

func (r *UserResource) GetUserByIDPost(c *gin.Context) {
	var user model.UserInfo
	c.BindJSON(&user)
	c.JSON(200, user)
}
