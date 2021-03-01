package main

import (
	"fmt"
	"ginPro1/user"
	"github.com/gin-gonic/gin"
	"log"
	"reflect"
)

func main() {

	engine := gin.Default()

	//i := gin.New() todo 和default的区别

	engine.GET("/hello", func(context *gin.Context) {

		context.Writer.Write([]byte("hello world!"))

		name := context.DefaultQuery("name", "null")

		name1 := context.Query("name")
		fmt.Println("name1 : ", name1)

		fmt.Println("name : ", name)
	})

	// todo 两种处理方式
	engine.Handle("GET", "/get1", func(context *gin.Context) {

		context.Writer.Write([]byte("get1 return...."))
	})

	engine.POST("/post1", func(context *gin.Context) {
		name := context.PostForm("name")
		age := context.PostForm("age")

		fmt.Println("name : ", name)
		fmt.Println("age : ", age, reflect.TypeOf(age))

		//todo 获取参数方式二
		qname, ok := context.GetPostForm("name")
		qname1, ok1 := context.GetPostForm("name1")
		fmt.Println("qName : ", qname, ok)
		fmt.Println("qName : ", qname1, ok1)

		context.Writer.Write([]byte("post1 run ... "))
	})

	// todo 路由参数的获取
	engine.GET("/user/:id", func(context *gin.Context) {

		id := context.Param("id")
		fmt.Println("id : ", id)
		context.Writer.Write([]byte(id))
	})

	// todo 结构体去接受数据
	engine.GET("/user", func(context *gin.Context) {
		full := context.FullPath()
		fmt.Println("user2 full path : ", full)

		user2 := user.User{}
		err := context.ShouldBindQuery(&user2)
		if err != nil {
			context.Writer.Write([]byte("error to change.."))
		}
		fmt.Println("user2 : ", user2)

	})

	//todo post结构体接受
	engine.POST("user", func(context *gin.Context) {
		fmt.Println("post user run...")

		user1 := user.User{}
		err := context.ShouldBind(&user1)
		if err != nil {
			log.Fatal("is error >>>")
			return
		}
		fmt.Println("post user : ", user1)

	})

	//todo 接受json格式的数据

	engine.POST("/json", func(context *gin.Context) {
		fmt.Println("json is run....")

		var user11 user.User
		err := context.BindJSON(&user11)
		if err != nil {
			log.Fatal("json false...")
		}
		fmt.Println("json user : ", user11, user11.UserName)
	})

	engine.Run()
}
