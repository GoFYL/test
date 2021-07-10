package main

import "fmt"

//结构体嵌套

type address struct {
	province string
	city     string
}

type person struct {
	name    string
	age     int
	address //匿名嵌套结构体	address:address
}

type company struct {
	name string
	addr address
}

func main() {
	p1 := person{
		name: "fyl",
		age:  18,
		address: address{
			province: "台州",
			city:     "黄岩",
		},
	}
	fmt.Println(p1)
	fmt.Println(p1.name, p1.address.city)
	fmt.Println(p1.city) //现在自己结构体找这个字段，找不到就去匿名嵌套结构体中查找该字段
	//若匿名嵌套结构体字段冲突，那就写全
}
