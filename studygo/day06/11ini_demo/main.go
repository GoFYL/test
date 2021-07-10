package main

import "fmt"

//ini配置文件解析器

//MysqlConfig MySQL配置结构体
type MysqlConfig struct {
	Address  string `ini:"address"`
	Port     int    `ini:"port"`
	Username string `ini:"username"`
	Password string `ini:"password"`
}



func loadIni(v interface{}) {
	//1.读文件得到字节类型数据
	//2.一行一行读数据
	// 2.1 如果是注释就跳过
	// 2.2 如果是[开头的就表示是节(section)
	// 2.3 如果不是[开头的就是=分割的键值对
}

func main() {
	var mc MysqlConfig
	loadIni(&mc)
	fmt.Println(mc.Address, mc.Password, mc.Port, mc.Username)
}
