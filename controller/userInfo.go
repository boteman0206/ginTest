package controller

import "github.com/gin-gonic/gin"

type UserController struct {
}

// todo userInfo的引擎
func (user UserController) UserRouter(engine *gin.Engine) {

	engine.GET("/userInfo", user.userInfo)

}

// todo userInfo的处理函数
func (user UserController) userInfo(ctx *gin.Context) {

	ctx.JSON(200, gin.H{
		"code": 900,
		"msg":  "hello first gin",
	})
}
