package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup
var notifyCh = make(chan struct{}, 0)

func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Printf("worker:%d start job:%d\n", id, j)
		time.Sleep(time.Second)
		fmt.Printf("worker:%d end job:%d\n", id, j)
		results <- j * 2
		notifyCh <- struct{}{}
	}
}

func main() {
	jobs := make(chan int, 100)
	results := make(chan int, 100)

	// 开启5个任务
	go func() {
		for j := 1; j <= 5; j++ {
			jobs <- j
		}
		close(jobs)
	}()

	// 开启3个goroutine ,,workPool
	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	go func() {
		for i := 0; i < 5; i++ {
			<-notifyCh
		}
		close(results)
	}()

	// for {
	// 	x, ok := <-results
	// 	if !ok { //什么时候ok=false?results这个通道被关闭的时候
	// 		break
	// 	}
	// 	fmt.Println(x)
	// }

	// 输出结果
	for v := range results {
		fmt.Println(v)
	}
}