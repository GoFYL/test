package main

import "fmt"

//存放元素的容器
//必须指定存放的元素的类型和长度

func main() {
	// a1 = []bool{true, false}
	// fmt.Printf("%v", a1)

	// a3 := [...]int{1, 2, 3, 4}
	// for i := 0; i < len(a3); i++ {
	// 	fmt.Println(a3[i])
	// }

	// for _, v := range a3 {
	// 	fmt.Println(v)
	// }
	// a11 := [3][2]int{
	// 	{2, 3},
	// 	{3, 4},
	// 	{5, 6},
	// }
	// fmt.Println(a11)

	//数组是值类型
	b1 := [...]int{1, 3, 5, 7, 8}
	b2 := b1
	b2[0] = 100
	fmt.Println(b1, b2)
	b1[1] = 100
	fmt.Println(b1, b2)
	b3 := []int{1, 3, 5, 7, 8}
	b4 := b3
	b4[0] = 100
	fmt.Println(b3, b4)
	b3[1] = 100
	fmt.Println(b3, b4)
	// sum := 0
	// for _, v := range b1 {
	// 	sum += v
	// }
	// fmt.Println(sum)

	// var ret [][2]int = make([][2]int, len(b1)*len(b1))
	// index := 0
	// for i := 0; i < len(b1); i++ {
	// 	for j := i + 1; j < len(b1); j++ {
	// 		if b1[i]+b1[j] == 8 {
	// 			ret[index][0] = i
	// 			ret[index][1] = j
	// 			index++
	// 		}
	// 	}
	// }
	// fmt.Println(ret)
}
