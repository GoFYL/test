package main

import (
	"fmt"
	"strings"
)

func main() {
	s1 := []string{"1", "2", "3"}
	// s1[3] = "4" //错误的写法，会导致编译错误：索引越界
	// fmt.Println(s1)

	s1 = append(s1, "1")

	fmt.Println(s1)

	a1 := "asd   sd ds d  s"
	a2 := strings.Split(a1, "  ")
	fmt.Println(a2)
}
