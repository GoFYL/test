package main

import (
	"encoding/json"
	"fmt"
)

//结构体与json

//1.序列化:  把go语言中的结构体变量-->json格式的字符串
//2.反序列化:json格式的字符串 --> Go语言中能够识别的结构体变量

type person struct {
	Name string `json:"name" db:"name" ini:"name"`
	Age  int    `json:"age"`
}

func main() {
	p1 := person{
		Name: "zhoulin",
		Age:  900,
	}
	//序列化
	b, err := json.Marshal(p1)
	if err != nil {
		fmt.Printf("marshal failed,err: %v\n", err)
	}
	fmt.Printf("%v\n", string(b))

	//反序列化
	str := `{"name":"力量","age":18}`
	var p2 person
	json.Unmarshal([]byte(str), &p2) //传指针是为了能在 json.Unmarshal内部修改p2的值
	fmt.Printf("%v\n", p2)
}
