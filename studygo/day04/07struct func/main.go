package main

import "fmt"

//构造函数

type person struct {
	name string
	age  int
}

//构造函数:约定成俗 用new开头
//返回的是结构体还是结构体指针
//当结构体比较大的时候尽量使用结构体指针，减少程序的内存开销
func newPerson(name1 string, age1 int) *person {
	return &person{
		name: name1,
		age:  age1,
	}
}
func main() {
	p1 := newPerson("元帅", 18)
	fmt.Println(p1)
}
