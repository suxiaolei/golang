package controller

func Register(ctx *gin.Contex) {
	ctx.JSON(utils.NewSucc("Register success !!!", gin.H{
		"msg": "Register success !!!",
	}))
}
