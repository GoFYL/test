package main

import (
	"flag"
	"fmt"
)

//flag  获取命令行参数
func main() {
	//创建一个标志位参数
	// name := flag.String("name", "王野", "请输入名字")
	// age := flag.Int("age", 9000, "请输入真实年龄")
	// married := flag.Bool("married", false, "结婚了吗")
	// cTime := flag.Duration("ct", time.Second, "结婚多久了")

	//使用flag
	// flag.Parse() //先解析再使用

	// fmt.Println(*name)
	// fmt.Println(*age)
	// fmt.Println(*married)
	// fmt.Println(*cTime)

	var name string
	flag.StringVar(&name, "name", "王野", "请输入名字")

	//使用flag
	//05flag_demo.exe -name fyl -age=18
	flag.Parse() //先解析再使用
	fmt.Println(name)

	fmt.Println(flag.Args())  //返回命令行后的其他参数，以[]string类型
	fmt.Println(flag.NArg())  //返回命令行参数后的其他参数的个数
	fmt.Println(flag.NFlag()) //返回使用的命令行参数的个数
}
