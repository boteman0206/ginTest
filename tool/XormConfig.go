package tool

import (
	"fmt"
	"ginPro1/Model"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

var DbEngin *XormEngin

type XormEngin struct {
	*xorm.Engine
}

//todo init orm engin
func OrmEngin(con *Config) (*XormEngin, error) {

	database := con.DataBase
	conn := database.User + ":" + database.Password +
		"@tcp(" + database.Host + ":" + database.Port + ")/" + database.DbName

	fmt.Println("conn : ", conn)
	engine, e := xorm.NewEngine(database.Driver, conn)

	if e != nil {
		return nil, e
	}

	// todo 调试是否打印sql语句
	engine.ShowSQL(true)

	// todo 初始化语句
	sync2 := engine.Sync2(new(Model.Code))
	if sync2 != nil {
		return nil, sync2
	}
	orm := new(XormEngin)
	orm.Engine = engine

	//全局的Db
	DbEngin = orm

	return orm, nil
}
