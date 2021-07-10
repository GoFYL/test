package main

import (
	"fmt"
	"time"

	calc "github.com/GoFYL/studygo/day05/10calc"
)

var a = 3

const pi = 3.14

func init() {
	fmt.Println("执行")
	fmt.Println(a, pi)
}
func main() {
	ret := calc.Add(5, 3)
	fmt.Println(ret)
	time.Now()
}
