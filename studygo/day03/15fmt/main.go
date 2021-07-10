package main

import (
	"fmt"
	"strings"
)

func main() {
	// var f float64 = 4
	f := 4.001
	if lenk(f) > 2 {
		fmt.Println("错误")
	}
}

func lenk(f float64) int {
	a := strings.Split(fmt.Sprintf("%v", f), ".")
	if len(a) == 1 {
		return 0
	} else {
		return len(a[1])
	}
}
