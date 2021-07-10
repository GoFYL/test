package main

import "fmt"

func main() {
	// //1. &:取地址
	// n := 18
	// p := &n
	// fmt.Println(p)
	// fmt.Printf("%T\n", p)
	// //2. *: 根据地址取值
	// m := *p
	// fmt.Println(m)

	var a1 *int //nil
	fmt.Println(a1)
	var a2 = new(int) //分配内存
	fmt.Println(a2)
	fmt.Println(*a2)
	*a2 = 100
	fmt.Println(*a2)
}
