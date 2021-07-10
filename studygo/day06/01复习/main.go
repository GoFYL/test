package main

import (
	"fmt"
)

//类型断言

func main() {
	var a interface{}
	a = 100
	//如何判断a的类型
	//类型断言
	//x.(T)

	if _, ok := a.(int8); !ok {
		fmt.Println("猜错了")
	}
	switch a.(type) {
	case int8:
		fmt.Println("int8")
	case int64:
		fmt.Println("int64")
	case int:
		fmt.Println("int")
	default:
		fmt.Println()
	}
}
