package main

import (
	"fmt"
	"strings"
)

func main() {
	//1.判断字符串中的汉字数量
	//难点是判断一个字符是汉字
	// count := 0
	s1 := "hello do  you do"
	s3 := strings.Split(s1, " ")
	map1 := make(map[string]int, 10)
	for _, v := range s3 {
		map1[v]++
	}
	fmt.Println(map1)
}
