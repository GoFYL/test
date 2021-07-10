package main

import "fmt"

func main() {
	//切片
	a1 := [...]int{1, 2, 3, 4, 5, 6, 7, 8} //定义数组
	a2 := []int{1, 2, 3, 4, 5, 6, 7, 8}    //定义切片
	// var s []int //定义一个存放int类型元素的切片
	s1 := a1[1:2]
	s2 := a2[1:2]
	//切片是引用类型，都指向了底层的一个数组
	fmt.Println(s1)
	fmt.Println(s2)
	// s2 := a1[1:5]
	// s3 := s2[2:]
	// fmt.Println(s2, cap(s2), s3, cap(s3))
}
