package main

import (
	"common/log"
	"ginTemplate/router"
	"github.com/astaxie/beego"
	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)

	r := gin.New()
	r.Use(gin.Recovery())

	router.NewUserResource(r)

	port := beego.AppConfig.DefaultString("PORT", "8080")
	log.Info(nil, "Service starting on port "+port)

	r.Run(":" + port) // listen and serve
}
