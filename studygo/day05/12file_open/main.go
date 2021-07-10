package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

//打开文件
func readFromFile1() {
	fileObj, err := os.Open("./main.go")
	if err != nil {
		fmt.Printf("open file failed, err:%v\n", err)
		return
	}
	//记得关闭文件
	defer fileObj.Close()
	//读文件
	tmp := make([]byte, 128) //指定读的长度
	for {
		n, err := fileObj.Read(tmp[:])
		if err != nil {
			fmt.Printf("read failed,err:%v\n", err)
		}
		fmt.Printf("读了个%d个字节\n", n)
		fmt.Println(string(tmp[:]))
		if n < 128 {
			return
		}
	}
}

//利用bufio这个包读文件
func readFromFile2() {
	fileObj, err := os.Open("./main.go")
	if err != nil {
		fmt.Println("open file failed,err:", err)
		return
	}
	//记得关闭文件
	defer fileObj.Close()
	//创建一个用来从文件中读内容的对象
	reader := bufio.NewReader(fileObj)
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			return
		}
		if err != nil {
			fmt.Println("read line failed,err:", err)
			return
		}
		fmt.Print(line)
	}
}

func readFromFileByIoutil() {
	ret, err := ioutil.ReadFile("./main.go")
	if err != nil {
		fmt.Println("read file failed,err", err)
	}
	fmt.Println(string(ret))
}

func main() {
	// readFromFile1()
	// readFromFile2()
	readFromFileByIoutil()
}
