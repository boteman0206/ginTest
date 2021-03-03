package controller

import (
	"fmt"
	"ginPro1/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type MemberController struct {
}

func (mer *MemberController) MemberRouter(cn *gin.Engine) {

	cn.GET("/sendMsg", mer.sendMsg)
}

func (mer *MemberController) sendMsg(cn *gin.Context) {

	query := cn.DefaultQuery("code", "")
	fmt.Println("query: ", query)

	//调用服务层处理数据
	memberService := new(service.MemberService)
	code := memberService.SendCode(query)

	// 处理返回
	if code {
		cn.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg":  "send msg successful",
		})
		return
	}

	cn.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "send msg fail",
	})
}
