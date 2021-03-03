package service

import (
	"fmt"
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
	return true
}
