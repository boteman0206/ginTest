package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {

	engine := gin.Default()

	// 返回字符串和字节数组
	engine.GET("/writeString", func(context *gin.Context) {

		const name = "hello go .."

		context.Writer.WriteString(name)
		//context.Writer.Write([]byte("hello byte"))

	})

	// todo 返回json格式的数据
	engine.GET("/writeJson", func(context *gin.Context) {

		context.JSON(http.StatusOK, map[string]interface{}{
			"name": "jack",
			"age":  90,
			"addr": "beijing",
		})

	})

	// todo 返回html的数据格式 必须先设置html的文件目录
	engine.LoadHTMLGlob("html/*")

	engine.GET("/writeHtml", func(context *gin.Context) {

		context.HTML(http.StatusOK, "index.html", nil)

	})

	// 写入文本到html文件中
	engine.GET("/writeTitle", func(context *gin.Context) {

		context.HTML(http.StatusOK, "title.html", gin.H{
			"title": "this is write title.",
		})

	})

	engine.Run()
}
