package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()

	// todo 分组
	userGroup := engine.Group("/user")

	userGroup.GET("/info", func(context *gin.Context) {

		fmt.Println("info run ....")
		context.JSON(200, gin.H{
			"code":    1,
			"message": "successful",
		})
	})

	// get
	userGroup.GET("/info/:id", func(context *gin.Context) {

		fmt.Println("info run ....")
		context.JSON(200, gin.H{
			"code":    1,
			"message": "info",
			"id":      context.Param("id"),
		})
	})

	// delete请求
	userGroup.DELETE("/:id", func(context *gin.Context) {

		fmt.Println("id run : ", context.Param("id"))
		context.JSON(200, gin.H{
			"code": 100,
			"id":   context.Param("id"),
		})
	})

	engine.Run()

}
