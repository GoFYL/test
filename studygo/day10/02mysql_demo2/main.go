package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql" //init()
)

//go连接MySQl示例

var db *sql.DB //是一个连接池对象

func initDB() (err error) {
	//连数据库
	dsn := "root:123456789@tcp(127.0.0.1:3306)/sql_test"
	//连接数据库
	db, err = sql.Open("mysql", dsn) //不会校验用户名和密码是否正确
	if err != nil {                  //dsn格式不正确的时候会报错
		fmt.Printf("dsn:%s invalid,err:%v \n", dsn, err)
		return
	}
	err = db.Ping() //尝试连接数据库
	if err != nil {
		fmt.Printf("open%s failed,err:%v \n", dsn, err)
		return
	}

	db.SetMaxOpenConns(5)    //设置数据库连接池的最大连接数
	db.SetConnMaxIdleTime(5) //设置数据库连接池的最大空闲连接数
	return
}

type user struct {
	id   int
	name string
	age  int
}

//查询
func query(id int) {
	// db.QueryRow()
	var u1 user
	//1.写查询单条记录的sql语音
	sqlStr := `select id,name,age from user where id=?;`
	//2.执行并拿出结果
	//从连接池里拿一个连接出来去数据库查询单条记录
	//必须对rowObj对象调用Scan方法,因为该方法会释放数据库连接
	db.QueryRow(sqlStr, id).Scan(&u1.id, &u1.name, &u1.age)
	fmt.Println(u1)
}

func insert() {

}

func main() {
	err := initDB()
	if err != nil {
		fmt.Printf("initDB failed,err:%v \n", err)
	}
	fmt.Println("连接数据库成功")
	query(1)
}
