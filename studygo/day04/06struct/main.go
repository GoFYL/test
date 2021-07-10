package main

import "fmt"

//结构体占用一块连续的内存空间

type x struct {
	a int8
	b int8
	c int8
}

func main() {
	m := x{
		a: int8(10),
		b: int8(20),
		c: int8(30),
	}
	fmt.Printf("%p %p %p\n", &m.a, &m.b, &m.c)
}
