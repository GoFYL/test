package main

import (
	"fmt"
	"math/rand"
	"time"
)

// channel

//往通道中发送值
func sendNum(ch chan<- int) {
	for {
		num := rand.Intn(10)
		ch <- num
		time.Sleep(2 * time.Second)
	}

}

func main() {
	ch := make(chan int, 1)
	// ch <- 100     //把一个值发送到通道中
	// <-ch          //把通道中100取出来
	// x, ok := <-ch //再取阻塞
	// fmt.Println(x, ok)
	go sendNum(ch)
	for {
		x, ok := <-ch //再取阻塞
		fmt.Println(x, ok)
		time.Sleep(time.Second)
	}
}
