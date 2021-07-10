package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql" //init()
)

//go连接MySQl示例
func query() {

}

func insert() {

}

func main() {
	//连数据库
	dsn := "root:123456789@tcp(127.0.0.1:3306)/goday10"
	//连接数据库
	db, err := sql.Open("mysql", dsn) //不会校验用户名和密码是否正确
	if err != nil {                   //dsn格式不正确的时候会报错
		fmt.Printf("dsn:%s invalid,err:%v \n", dsn, err)
		return
	}
	err = db.Ping() //尝试连接数据库
	if err != nil {
		fmt.Printf("open%s failed,err:%v \n", dsn, err)
		return
	}
	fmt.Println("成功")
}
