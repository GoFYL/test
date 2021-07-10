package main

import (
	"fmt"
	"strconv"
)

//strconv

func main() {
	//把字符串中解析对应的数据
	str := "100000"
	ret1, _ := strconv.ParseInt(str, 10, 64)
	fmt.Println(ret1)

	//
	retInt, _ := strconv.Atoi(str)
	fmt.Println(retInt)

	//把数字转换成字符串类型
	i := 97
	// ret2 := string(i) //"a"
	ret2 := fmt.Sprintf("%d", i) //"97"
	fmt.Printf("%#v", ret2)

	ret3 := strconv.Itoa(i)
	fmt.Printf("%#v", ret3)

	//从字符串中解析出布尔
	boolstr := "true"
	ret4, _ := strconv.ParseBool(boolstr)
	fmt.Printf("%#v", ret4)
}
