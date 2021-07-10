package main

import "fmt"

//结构体

type person struct {
	name   string
	age    int
	gender string
	hobby  []string
}

func main() {
	//声明一个person类型的变量p
	var p person
	p.name = "fyl"
	p.age = 9000
	p.gender = "男"
	p.hobby = []string{"篮球", "足球"}
	fmt.Println(p)
	//访问变量p的字段
	fmt.Printf("%T\n", p)
	fmt.Println(p.name)

	var s struct {
		x int
	}
	s.x = 1
	fmt.Println(s)
}
