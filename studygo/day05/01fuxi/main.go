package main

import (
	"encoding/json"
	"fmt"
)

type MyInt int    //自定义类型
type newInt = int //类型别名
//类型别名只在代码编写过程中有效，编译完之后就不存在，内置的byte和rune都属于类型别名
type person struct {
	name string
	age  int
}

//构造(结构体变量的)函数
func newPerson(name string, age int) person {
	return person{
		name: name,
		age:  age,
	}
}

//方法
func (p person) dream(str string) {
	fmt.Printf("%s有梦想,梦想是%s\n", p.name, str)
}

//结构体是值类型
func (p person) guonian() {
	p.age++ //此处的p是p4的副本,改的是副本
}

//指针接受者
//1.需要修改结构体变量的值要用指针接受
//2.结构体本身比较大，拷贝的内存比较大时也要用指针接受者
//3.保持一致性，如果有一个方法用了指针接受者，那其他方法保持统一
func (p *person) zhenguonian() {
	p.age++
}
func main() {
	// p4 := newPerson("zhou", 18)
	// p4.dream("当个咸鱼")
	// fmt.Println(p4.age)
	// p4.guonian()
	// fmt.Println(p4.age)
	// p4.zhenguonian()
	// fmt.Println(p4.age)

	//结构体嵌套
	type addr struct {
		province string
		city     string
	}
	type student struct {
		name    string
		address addr //嵌套别的结构体
		addr         //匿名嵌套结构体，使用类型名做名称
	}

	//json序列化
	type point struct {
		// x int `json:"gaiming"`  //json访问不到
		X int `json:"gaiming"`
		Y int `json:"gaiming2"` //json可以访问
	}

	p1 := point{100, 200}
	b, err := json.Marshal(p1) //在json里序列化，point结构体中的变量名首字母要大写,json才能访问到
	if err != nil {
		fmt.Printf("出错了，err:%v", err)
	}
	fmt.Println(string(b))

	//json反序列化:由字符串 -->GO中的结构体变量
	str1 := `{"gaiming":10,"gaiming2":800}`
	var po2 point                            //造一个结构爱变量，准备接受反序列化的值
	err = json.Unmarshal([]byte(str1), &po2) //一定要传指针才能修改结构体变量的值
	if err != nil {
		fmt.Printf("unmarshal failed,err:%v\n", err)
	}
	fmt.Println(po2)
}
