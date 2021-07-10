package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

//打开文件写内容

func writeDemo1() {
	fileObj, err := os.OpenFile("./xx.txt", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		fmt.Println("open file failed,err:", err)
		return
	}
	defer fileObj.Close()

	fmt.Println(fileObj)
	//write
	fileObj.Write([]byte("周林 mengbi le！\n"))
	//writeString
	fileObj.WriteString("周林解释不了")
}

func writeDemo2() {
	fileObj, err := os.OpenFile("./xx.txt", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		fmt.Println("open file failed,err:", err)
		return
	}
	defer fileObj.Close()
	//创建一个写的对象
	wr := bufio.NewWriter(fileObj)
	wr.WriteString("hello,沙河\n") //写到缓存中
	wr.Flush()                   //将缓存中的内容写入文件
}

func writeDemo3() {
	str := "hello,沙河"
	err := ioutil.WriteFile("./xx.txt", []byte(str), 0666)
	if err != nil {
		fmt.Println("write file failed,err:", err)
		return
	}
}
func main() {
	// writeDemo1()
	writeDemo2()
	// writeDemo3()
}
