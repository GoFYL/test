package main

import "fmt"

//接口示例2

//接口是一种类型，是一种特殊的类型，它规定了变量有哪些方法
//我不关心一个变量是什么类型，我只关心能调用它的什么方法

//定义一个car接口类型
//不管是什么结构体，只要有run方法都能是car类型
type car interface {
	run()
}
type falali struct {
	brand string
}

func (f falali) run() {
	fmt.Println("速度70")
}

type baoshijie struct {
	brand string
}

func (b baoshijie) run() {
	fmt.Println("速度80")
}

func driver(c car) {
	c.run()
}

func main() {
	f1 := falali{
		"法拉利",
	}
	b1 := baoshijie{
		"保时捷",
	}
	driver(f1)
	driver(b1)
}
