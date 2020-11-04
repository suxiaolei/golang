package main

import (
	"github.com/gin-gonic/gin"
	"github.com/suxiaolei/golang/gin-web/controller"
)

func main() {
	r := gin.Default()
	r.GET("/ping", controller.Ping)
	r.GET("/register", controller.Register)
	r.Run()
}
