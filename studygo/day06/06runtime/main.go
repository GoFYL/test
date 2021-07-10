package main

import (
	"fmt"
	"path"
	"runtime"
)

//runtime.Caller()

func f() {
	pc, file, line, ok := runtime.Caller(0)
	if !ok {
		fmt.Println("failed")
		return
	}
	funcName := runtime.FuncForPC(pc).Name()
	fmt.Println(funcName)
	fmt.Println(file) //06runtime/main.gp
	fmt.Println(path.Base(file))
	fmt.Println(line) //11
}

func f1() {
	f()
}

func main() {
	f1()
}
