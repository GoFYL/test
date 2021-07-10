package main

import (
	"fmt"
	"io"
	"os"
)

//文件操作

func f1() {
	var fileObj *os.File
	var err error
	fileObj, err = os.Open("./main.go")
	if err != nil {
		fmt.Println("wrong")
		return
	}
	defer fileObj.Close()
}
func f2() {
	//打开要操作的文件
	fileObj, err := os.OpenFile("./sb.txt", os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		fmt.Printf("failed,err:%v\n", err)
		return
	}
	// defer fileObj.Close()

	//因为没有办法直接在文件中插入内容，所以要借助一个临时文件
	tmpFile, err := os.OpenFile("./sb.tmp", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0666)
	if err != nil {
		fmt.Printf("failed,err:%v\n", err)
		return
	}
	defer tmpFile.Close()
	// 读取源文件写入临时文件
	var ret [1]byte
	n, err := fileObj.Read(ret[:])
	if err != nil {
		fmt.Printf("failed,err:%v\n", err)
		return
	}
	//写入临时文件
	tmpFile.Write(ret[:n])

	//再写入要插入的内容
	var s = []byte{'c'}
	tmpFile.Write(s)
	//紧接着把源文件后续的内容写入临时文件
	var x [1024]byte
	for {
		n, err := fileObj.Read(x[:])
		if err == io.EOF {
			tmpFile.Write(x[:n])
			break
		}
		if err != nil {
			fmt.Println("wrong,err:", err)
			return
		}
		tmpFile.Write(x[:n])
	}
	//源文件后续也写入了临时文件中
	fileObj.Close()
	tmpFile.Close()
	os.Rename("./sb.tmp", "./sb.txt")
}

func main() {
	f2()
}