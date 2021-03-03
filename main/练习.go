package main

import (
	"fmt"
	"ginPro1/controller"
	"ginPro1/tool"
	"github.com/gin-gonic/gin"
)

// 项目入口

func main() {

	// todo 读取配置文件
	config, e := tool.ParseConfig("./config/app.json")
	if e != nil {
		fmt.Println("error...")
		return
	}
	fmt.Println("config: ", config)

	engine := gin.Default()

	// todo userInfo定义路由
	initRouter(engine)

	engine.Run(config.AppHost + ":" + config.AppPort)
}

//todo 路由的初始化设置
func initRouter(router *gin.Engine) {

	new(controller.UserController).UserRouter(router)

	new(controller.MemberController).MemberRouter(router)
}
