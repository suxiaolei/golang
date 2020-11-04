package controller

import "github.com/gin-gonic/gin"

func Register(ctx *gin.Context) {
	ctx.JSON(utils.NewSucc("Register success !!!", gin.H{
		"msg": "pong",
	})）
}
