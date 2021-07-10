package main

import "fmt"

func main() {
	var m1 map[string]int
	fmt.Println(m1 == nil)        //还没有初始化(没有在内存中开辟空间)
	m1 = make(map[string]int, 10) //估算好改map容量,避免在程序运行期间再动态扩容
	m1["理想"] = 18
	m1["快乐"] = 19
	fmt.Println(m1)
	delete(m1, "快乐")
	fmt.Println(m1)
}
