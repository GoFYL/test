package main

import (
	"fmt"
	"sync"
)

//channel

var b chan int //需要指定通道中元素的类型
var wg sync.WaitGroup

func noBufChannel() {
	fmt.Println(b)     //nil
	b = make(chan int) //不带缓冲区通道初始化  类似接力棒，得有下一棒接收才能发送
	wg.Add(1)
	go func() {
		defer wg.Done()
		x := <-b
		fmt.Println("后台goroutine从通道b中取到了，", x)
	}()
	b <- 10
	fmt.Println("10发送到通道b中")
	wg.Wait()
}

func bufChannel() {
	fmt.Println(b)         //nil
	b = make(chan int, 10) //带缓冲区的通道初始化

	b <- 10
	fmt.Println("10发送到通道b中")
	b <- 20 //hang住了
	fmt.Println("20发送到通道b中")
	x := <-b
	fmt.Println("从通道b中取出", x)
	close(b)
}

func main() {
	// bufChannel()
	noBufChannel()
}
