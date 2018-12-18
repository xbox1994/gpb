package router

import (
	"ginTemplate/service"
	"github.com/gin-gonic/gin"
)

func NewUserResource(e *gin.Engine) {
	u := service.UserResource{}

	// Setup Routes
	e.GET("/users", u.GetAllUsers)
	e.GET("/users/:id", u.GetUserByID)

	e.POST("/users", u.GetUserByIDPost)
}
