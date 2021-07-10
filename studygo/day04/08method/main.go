package main

import "fmt"

//方法

//标识符:变量名 函数名 类型名 方法名
//GO语言中如果标识符首字母是大写的，就表示对外可见（暴露的，公有的）

//dog 这是一个狗的结构体
type dog struct {
	name string
	age  int
}

//构造函数
func newDog(name1 string, age1 int) dog {
	return dog{
		name: name1,
		age:  age1,
	}
}

//方法是作用于特定类型的函数
//接受者表示的是调用该方法的具体类型变量，多用类型名首字母小写表示
func (d dog) wang() {
	fmt.Printf("%s:汪汪汪\n", d.name)
}

//使用值接受者：传拷贝进去
func (d dog) guonian() {
	d.age++
}

//指针接受者：传内存地址进去
func (d *dog) zhenguonian() {
	d.age++
}

func main() {
	d1 := newDog("ss", 18)
	d1.wang()
	fmt.Println(d1.age) //18
	d1.guonian()
	fmt.Println(d1.age) //18
	d1.zhenguonian()
	fmt.Println(d1.age) //19
}
