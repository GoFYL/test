package main

import (
	"fmt"
	"time"
)

//goroutine

func hello(i int) {
	fmt.Println("hello,", i)
}

//程序启动后会创建一个主goroutine去执行
func main() {
	for i := 0; i < 100; i++ {
		//go 8hello(i)

		// go func() {
		// 	fmt.Println(i)
		// }()

		go func(i int) {
			fmt.Println(i)
		}(i)
	}
	fmt.Println("main")
	time.Sleep(time.Second)
	//main函数结束了,由main函数启动的goroutine也都结束了
}
