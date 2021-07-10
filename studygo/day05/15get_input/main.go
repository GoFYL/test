package main

import (
	"bufio"
	"fmt"
	"os"
)

//获取用户输入时如果有空格\

func useScan() {
	var s string
	fmt.Printf("请输入内容:")
	fmt.Println(&s)
	fmt.Println("你输入的内容是，", s)
}

func useBufio() {
	var s string
	reader := bufio.NewReader(os.Stdin)
	s, _ = reader.ReadString('\n')
	fmt.Println("你输入的内容是，", s)
}
func main() {
	// useScan()
	useBufio()

	// logger.Debug()
	// logger.Trace()
	// logger.Warning()
	// logger.Info()
	// logger.Error("日志内容")

	//写日志
	// fmt.Fprintln(os.Stdout, "这是一条日志记录!")
	// os.OpenFile("./xxx.log")
}
