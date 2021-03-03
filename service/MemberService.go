package service

import (
	"fmt"
	"ginPro1/Model"
	"ginPro1/dao"
	"ginPro1/tool"
	"math/rand"
	"time"
)

type MemberService struct {
}

func (ms *MemberService) SendCode(phone string) bool {

	if len(phone) == 0 {
		return false
	}

	//产生验证码
	sprintf := fmt.Sprintf("%04v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(1000))
	fmt.Println("sprintF: ", sprintf)

	//保存到数据库
	code := Model.Code{Name: "jack", Code: sprintf, CreateTime: time.Now().Unix()}

	//todo 获取数据库连接
	codeDao := dao.CodeDao{tool.DbEngin}
	insertCode := codeDao.InsertCode(code)

	if insertCode > 0 {
		fmt.Println("插入数据成功。。。。")
	} else {
		fmt.Println("插入数据失败、、、")
	}

	return true
}
