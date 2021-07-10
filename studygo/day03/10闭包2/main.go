package main

import "fmt"

//闭包是什么？
//闭包是一个函数，这个函数包含了他外部作用域的一个变量

//底层原理
//1.函数可以作为返回值
//2.函数内部查找变量的顺序，现在自己内部找，找不到去外层找

func main() {
	ret := bibao(f2, 100, 200)
	f1(ret)
}

func f1(f func()) {
	fmt.Println("this is f1")
	f()
}

func f2(x, y int) {
	fmt.Println("this is f2")
	fmt.Println((x + y))
}

//f2作为参数传入
func bibao(f func(int, int), x, y int) func() {
	tmp := func() {
		f(x, y)
	}
	return tmp
}
