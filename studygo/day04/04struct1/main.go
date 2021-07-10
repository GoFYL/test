package main

import "fmt"

//结构体是值类型
type person struct {
	name, gender string
}

//go语言中国函数参数永远传的是拷贝
func f(x person) {
	x.gender = "女" //修改的是副本的gender
}

func f2(x *person) {
	// (*x).gender = "女" //根据内存地址找到那个原变量，修改的就是原变量
	x.gender = "女" //语法糖，自动根据指针找对应的变量
}

func main() {
	var p person
	p.name = "ss"
	p.gender = "男"
	f(p)
	fmt.Println(p.gender)
	f2(&p)
	fmt.Println(p.gender)

	//2.结构体指针2
	//2.1 key-value初始化
	var p3 = &person{
		name:   "刷",
		gender: "男",
	}
	fmt.Printf("%#v\n", p3)

	//2.2使用值列表的形式初始化，值得顺序要和结构体定义时字段的顺序一致
	var p4 = &person{
		"刷",
		"男",
	}
	fmt.Printf("%#v", p4)
}
