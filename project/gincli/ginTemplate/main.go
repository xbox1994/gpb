package main

import (
	"ginTemplate/config"
	"ginTemplate/dao"
	"ginTemplate/routers"
	"github.com/gin-gonic/gin"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

func main() {
	pflag.Parse()
	////init config about viper
	//if err := config.Init(*cfg); err != nil {
	//	panic(err)
	//}
	config.Init()
	dao.Demo()
	gin.SetMode(viper.GetString("GIN_MODE"))
	r := gin.Default()
	routers.Routes(r)
	r.Run() // listen and serve on 0.0.0.0:8080
}
