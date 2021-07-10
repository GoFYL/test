package main

import "fmt"

//匿名字段
//字段比较少也比较简单的场景
//不常用

type person struct {
	string
	int
}

func main() {
	p1 := person{
		"dd",
		188,
	}
	//匿名结构体
	var a = struct {
		x int
		y int
	}{10, 20}
	fmt.Println(a)

	fmt.Println(p1)
	fmt.Println(p1.string)
}
