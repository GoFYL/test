package main

import "fmt"

type animal struct {
	name string
}

//给animal实现一个移动的方法
func (a animal) move() {
	fmt.Printf("%s会动!", a.name)
}

//狗类
type dog struct {
	feet   uint8
	animal //animal拥有的方法,dog此时也有了
}

//给dog实现汪汪汪的发那个发
func (d dog) wang() {
	fmt.Printf("%s:汪汪汪~", d.name)

}
func main() {
	d1 := dog{
		animal: animal{name: "周林"},
		feet:   4,
	}
	fmt.Println(d1)
	d1.wang()
	d1.move()
}
