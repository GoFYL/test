package main

import "fmt"

//如何判断一个链表有没有闭环

//x一次走1步，y一次走2步
//让他们去走，如果某一时刻，他们在用一个节点相遇说明这个链表有闭环
type a struct {
	val  int
	next *a
}

//走台阶 递归很慢重复计算
func f(n int) int {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	return f(n-1) + f(n-2)
}

func f1(n int) int {
	if n == 1 || n == 0 {
		return 1
	}
	if n == 2 {
		return 2
	}
	x := 1
	y := 2
	for i := 3; i <= n; i++ {
		x, y = y, x+y
	}
	return y
}

func main() {
	fmt.Println(f1(3))
}
