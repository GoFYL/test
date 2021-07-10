package main

import "fmt"

//二进制实际用途

const (
	eat   int = 4
	sleep int = 2
	da    int = 1
)

//111
//左边1表示吃饭  100
//中间1表示睡觉  010
//右边1表示打豆  001

func f(arg int) {
	fmt.Printf("%b\n", arg)
}

func main() {
	f(eat | da) //101
}
