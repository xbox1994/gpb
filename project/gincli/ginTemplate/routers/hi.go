package routers

import "github.com/gin-gonic/gin"

func HiRoutes(route *gin.Engine){
	hi := route.Group("/hi")
	hi.GET("/wps", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "hi wps",
		})
	})
}