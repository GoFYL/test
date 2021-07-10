package calc

import "fmt"

//包中的标识符（变量名\函数名\结构体\接口等）如果首字母是小写的，表示私有（只能在当前包中使用）
//首字母大写的标识符可以被外部包调用

func init() {
	fmt.Println("调用我执行")
}
func Add(x, y int) int {
	return x + y
}
