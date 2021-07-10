package main

import "fmt"

//引入接口的实例

type speaker interface {
	speak() //只要实现了speak方法的变量都是speaker类型
}
type cat struct {
}

type dog struct {
}

func (c cat) speak() {
	fmt.Println("喵喵喵")
}

func (d dog) speak() {
	fmt.Println("汪汪汪")
}

func da(x speaker) {
	//接受一个参数，传进来什么我就打什么

	x.speak() //挨打了就要叫
}
func main() {
	var c1 cat
	var d1 dog
	da(c1)
	da(d1)

	var ss speaker //定义一个接口类型：speaker的变量：ss
	ss = c1
	da(ss)
	fmt.Println()
}
