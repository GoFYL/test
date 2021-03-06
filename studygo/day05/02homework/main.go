package main

import (
	"fmt"
	"os"
)

var smr stuManeger //声明一个全局的变量 学生管理者smr
//菜单函数
func showMenu() {
	fmt.Println("欢迎光临学生管理系统")
	fmt.Println(`
	1.查看所有学生
	2.添加学生
	3.修改学生
	4.删除学生
	5.退出
	`)
}

func main() {
	smr = stuManeger{
		allStudent: make(map[int]student, 50),
	}
	// stu1 := stuManeger{
	// 	allStudent[0]=-1
	// }
	var (
		choice int
	)
	for {
		showMenu()
		//等待用户输入
		fmt.Print("请选择序号:")
		fmt.Scanln(&choice)
		fmt.Println("你输入的是:", choice)
		switch choice {
		case 1:
			smr.showStudent()
		case 2:
			smr.addStudent()
		case 3:
			smr.editStudent()
		case 4:
			smr.deleteStudent()
		case 5:
			os.Exit(1)
		default:
			fmt.Println("滚")
		}
	}
}
