package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func HelloRoutes(route *gin.Engine){
	hello := route.Group("/hello")
	hello.GET("/world", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": viper.GetString("GIN_MODE"),
		})
	})
}