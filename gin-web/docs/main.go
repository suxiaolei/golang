package main

import (
	"github.com/suxiaolei/golang/gin-web/controller"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/ping", controller.Ping)

	r.Run()
}
