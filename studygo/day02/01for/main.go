package main

import "fmt"

func main() {
	//i=5 跳出for循环
	for i := 0; i < 10; i++ {
		if i == 5 {
			break
		}
		fmt.Println(i)
	}
	fmt.Println("over")
	//i=5 跳过此次for循环，继续下一循环
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if j > 0 {
				goto XX
			}
			fmt.Printf("%d-%d", i, j)
		}
	}
	fmt.Println("over")

XX:
	println("o ver")
}
