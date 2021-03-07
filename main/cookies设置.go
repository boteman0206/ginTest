package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AuthMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取客户端cookie并校验
		if cookie, err := c.Cookie("abc"); err == nil {
			if cookie == "123" {
				c.Next()
				return
			}
		}
		// 返回错误
		c.JSON(http.StatusUnauthorized, gin.H{"error": "err"})
		// 若验证不通过，不再调用后续的函数处理
		c.Abort()
		return
	}
}

func main() {
	engine := gin.Default()

	engine.GET("/cookie", func(context *gin.Context) {

		val, err := context.Cookie("key_name")
		if err != nil {
			val = "notSet"
			context.SetCookie("key_cookie", "value_cookie", 60, "/",
				"localhost", false, true)
		}

		fmt.Println("val : ", val)
	})

	engine.GET("/login", func(c *gin.Context) {
		// 设置cookie
		c.SetCookie("abc", "123", 60, "/",
			"localhost", false, true)
		// 返回信息
		c.String(200, "Login success!")
	})

	engine.GET("/home", AuthMiddleWare(), func(c *gin.Context) {

		c.JSON(200, gin.H{"data": "home"})
	})

	engine.Run()
}
