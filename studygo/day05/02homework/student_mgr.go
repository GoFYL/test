package main

import "fmt"

/*
学生管理系统
有一个物件:
	1.它保存了一些数据  -->结构体字段
	2.它有四个功能      -->结构体方法
*/

type student struct {
	id   int
	name string
}

func newStudent(id int, name string) student {
	return student{
		id:   id,
		name: name,
	}
}

//造一个学生的管理者
type stuManeger struct {
	allStudent map[int]student
}

//查看学生

func (s stuManeger) showStudent() {
	for _, stu := range s.allStudent {
		fmt.Printf("学号:%d,姓名:%s\n", stu.id, stu.name)
	}
}

//增加学生
func (s stuManeger) addStudent() {
	var (
		id   int
		name string
	)
	for {
		fmt.Printf("输入学生id:")
		fmt.Scanln(&id)
		fmt.Printf("输入学生姓名:")
		fmt.Scanln(&name)
		newStudent(id, name)
		if _, ok := s.allStudent[id]; ok {
			fmt.Printf("id已存在请检查\n")
		} else {
			s.allStudent[id] = student{
				id,
				name,
			}
			fmt.Println("添加成功！")
			break
		}
	}
}

//修改学生
func (s stuManeger) editStudent() {
	var (
		id   int
		name string
	)
	for {
		fmt.Printf("输入学生id:")
		fmt.Scanln(&id)
		fmt.Printf("输入学生姓名:")
		fmt.Scanln(&name)
		newStudent(id, name)
		if _, ok := s.allStudent[id]; !ok {
			fmt.Printf("id不存在请检查\n")
		} else {
			s.allStudent[id] = student{
				id,
				name,
			}
			fmt.Println("修改成功！")
			break
		}
	}
}

//删除学生
func (s stuManeger) deleteStudent() {
	var (
		id int
	)
	for {
		fmt.Printf("输入学生id:")
		fmt.Scanln(&id)
		if _, ok := s.allStudent[id]; !ok {
			fmt.Printf("id不存在请检查\n")
		} else {
			delete(s.allStudent, id)
			fmt.Println("删除成功！")
			break
		}
	}
}
