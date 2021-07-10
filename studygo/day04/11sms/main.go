package main

import (
	"fmt"
	"os"
)

/*
	函数版学生管理系统
	写一个系统能够查看\新增学生\删除学生
*/

var (
	allStudent map[int]*student //变量声明

)

type student struct {
	id   int
	name string
}

//newStudent是student的构造函数
func newStudent(id int, name string) *student {
	return &student{
		id:   id,
		name: name,
	}
}
func showAllStudent() {
	//把所有的学生都打印出来
	for k, v := range allStudent {
		fmt.Printf("学号:%d 姓名：%s\n", k, v.name)
	}
}
func addStudent() {
	//向allStudent中添加新的学生
	//1.创建一个新学生
	//1.1获取用户输入
	var (
		id   int
		name string
	)
	fmt.Printf("请输入学生学号:")
	fmt.Scanln(&id)
	fmt.Printf("请输入学生姓名:")
	fmt.Scanln(&name)
	//1.2造学生(调用student的构造函数)
	newStu := newStudent(id, name)
	//2.追加到allstudent这个map中
	allStudent[id] = newStu
}
func deleteStudent() {
	//1.请用户输入要删除的学生的学号
	var (
		deleteId int
	)
	fmt.Print("请输入学生学号:")
	fmt.Scanln(&deleteId)
	//2.去allstudent这个map中根据学号删除对应的键值对
	delete(allStudent, deleteId)
}

func main() {
	allStudent = make(map[int]*student, 30) //初始化（开辟内存空间）
	//1.打印菜单
	for true {
		fmt.Println("欢迎光临学生管理系统！")
		fmt.Println(`
		1.查看所有学生
		2.新增学生
		3.删除学生
		4.退出系统
		`)
		fmt.Print("请输入你要干啥：")
		//2.等待用户选择要做什么
		var choice int
		fmt.Scanln(&choice)
		fmt.Printf("你选择了%d这个选项！\n", choice)
		//3.执行对应的函数
		switch choice {
		case 1:
			showAllStudent()
		case 2:
			addStudent()
		case 3:
			deleteStudent()
		case 4:
			os.Exit(1) //退出
		default:
			fmt.Println("滚")
		}
		// Q:
		// 	break
	}
}
