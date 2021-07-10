package main

import (
	"encoding/json"
	"fmt"
)

//json
type person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	str := `{"name":"zhoulin","age":9000}`
	var p person
	err := json.Unmarshal([]byte(str), &p)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(p.Name, p.Age)
}
