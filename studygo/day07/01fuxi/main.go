package main

import "encoding/json"

//反射

type student struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	str := `{"name":"zhoulin","age":990}`
	var stu1 student
	json.Unmarshal([]byte(str), &stu1)
}
