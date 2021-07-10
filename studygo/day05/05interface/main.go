package main

import "fmt"

//接口实现

type animal interface {
	move()
	eat(string)
}

type cat struct {
	name string
	feet int8
}

func (c cat) move() {
	fmt.Println("猫动")
}
func (c cat) eat(food string) {
	fmt.Printf("猫吃%s\n", food)
}

type chicken struct {
	feet int8
}

func (c chicken) move() {
	fmt.Println("鸡动")
}
func (c chicken) eat(food string) {
	fmt.Printf("吃%s\n", food)
}

func main() {
	var a1 animal
	fmt.Printf("%T\n", a1)
	c1 := cat{
		"淘气",
		4,
	}
	a1 = c1
	a1.eat("小黄鱼")
	fmt.Printf("%T\n", a1)

	kfc := chicken{
		2,
	}
	a1 = kfc
	a1.eat("饲料")
	fmt.Printf("%T\n", a1)
}
