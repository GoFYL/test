package main

import (
	"fmt"
	"sync"
)

//channel练习

//1.启动一个goroutine,生成100个数发送到ch1
//2.启动一个goroutine,从ch1中取值，计算其平方放到ch2中
//3.在main中，从ch2取值打印出来

var wg sync.WaitGroup
var once sync.Once

func f1(ch1 chan<- int) { //ch1单项通道 只能往ch1中存值
	defer wg.Done()
	for i := 0; i < 10; i++ {
		ch1 <- i
	}
	close(ch1)
}
func f2(ch1 <-chan int, ch2 chan<- int) {
	//ch1只能取值，ch2只能存值
	defer wg.Done()
	// for {
	// 	x, ok := <-ch1
	// 	if !ok {
	// 		break
	// 	}
	// 	ch2 <- x * x
	// }
	for i := 0; i < 8; i++ {
		x := <-ch1
		ch2 <- x * x
	}
	once.Do(func() { close((ch2)) })
}

func main() {
	a := make(chan int, 10)
	b := make(chan int, 10)
	wg.Add(2)
	go f1(a)
	go f2(a, b)
	for ret := range b {
		fmt.Println(ret)
	}
	for ret := range a {
		fmt.Println(ret)
	}
	wg.Wait()

	// time.Tick()
}
