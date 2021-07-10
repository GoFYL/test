package main

import "fmt"

//空接口

//interface:关键字
//interface{}:空接口类型

//空接口作为函数参数
func show(a interface{}) {
	fmt.Printf("type:%T value:%v\n", a, a)
}
func main() {
	var m1 map[string]interface{}
	m1 = make(map[string]interface{}, 10)
	m1["name"] = "f"
	m1["age"] = 90
	m1["merried"] = true
	fmt.Println(m1)

	show(false)
	show(nil)
	show(m1)
}
