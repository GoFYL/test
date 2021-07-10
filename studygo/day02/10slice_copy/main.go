package main

import "fmt"

func main() {
	a1 := []int{1, 2, 3}
	a2 := a1
	a3 := []int{}
	a4 := make([]int, len(a1))
	copy(a4, a1)
	fmt.Println(a1, a2, a3, a4)
	a1[0] = 100
	fmt.Println(a1, a2, a3, a4)
	fmt.Printf("%T,%p", a1, a1)
}
