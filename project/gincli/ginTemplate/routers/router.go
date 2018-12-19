package routers

import "github.com/gin-gonic/gin"

func Routes(route *gin.Engine){
	HelloRoutes(route)
	HiRoutes(route)
}