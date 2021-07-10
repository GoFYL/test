package main

import "fmt"

//关闭通道

func main() {
	ch1 := make(chan int, 2)
	ch1 <- 10
	ch1 <- 20
	close(ch1)
	// for v := range ch1 {
	// 	fmt.Println(v)
	// }
	<-ch1
	<-ch1
	x, ok := <-ch1
	fmt.Println(x, ok) //对于关闭的通道都是可以取值的
}
