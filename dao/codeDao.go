package dao

import (
	"ginPro1/Model"
	"ginPro1/tool"
	"log"
)

type CodeDao struct {
	*tool.XormEngin
}

func (md CodeDao) InsertCode(code Model.Code) int64 {

	i, e := md.InsertOne(code)
	if e != nil {
		log.Fatal(e.Error())
	}
	return i
}
