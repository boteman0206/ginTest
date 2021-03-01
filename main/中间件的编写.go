package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

// 自定义中间件
func myMiddle1(c *gin.Context) {
	name := c.Query("name")
	pwd := c.Query("pwd")

	fmt.Println("myMiddle1 run ...", name, " pwd ", pwd)
	if name != "jack" || pwd != "123" {
		fmt.Println("用户名或者密码不对！")
		c.Abort()
	}
	before := c.Writer.Status()
	fmt.Println(" next before : ", before)

	//todo next将函数一分为二
	c.Next()

	after := c.Writer.Status()
	fmt.Println("next 之后： ", after)

	//todo 直接获取返回的json的值

	// todo 获取上下文设置的值
	value, exists := c.Get("userId")
	if exists {
		fmt.Println("userId value: ", value)
	}

	funcValue, exists := c.Get("func")
	if exists {
		fmt.Println("userId value: ", funcValue)
	}

}

func myMiddle2(c *gin.Context) {

	// todo 设置上下文的值
	c.Set("userId", "89")
	fmt.Println(" myMiddle2 run ...")

}

func main() {
	engine := gin.Default()

	userGroup := engine.Group("/user", myMiddle2, myMiddle1)

	userGroup.GET("/obj", func(context *gin.Context) {
		fmt.Println("  函数 开始 执行 。。。")

		context.Set("func", "setFunc")
		context.JSON(302, gin.H{
			"code": 100,
			"msg":  "myMiddle",
		})
	})
	engine.Run()

}
