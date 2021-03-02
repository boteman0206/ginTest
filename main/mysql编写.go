package main

import (
	"database/sql"
	"fmt"
	"ginPro1/user"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func main() {

	conStr := "root:1234@tcp(127.0.0.1:3306)/login"
	db, err := sql.Open("mysql", conStr)

	fmt.Println("db: ", db)
	if err != nil {
		log.Fatal(err.Error())
		return
	}

	// 创建数据库
	result, _ := db.Query("select id, name, password from user")

	var per user.SqlUser

	// todo 遍历读取result结果集合
	for result.Next() {
		// 注意结构体大写
		scan := result.Scan(&per.Id, &per.Name, &per.Password)
		if scan != nil {
			fmt.Println("打印错误。。。。")
		}
		fmt.Println("per:  ", per)
	}

}
