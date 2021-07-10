package main

import "fmt"

func main() {
	s1 := make([]int, 5, 10)
	fmt.Printf("s1=%v len(s1)=%d cap(s1)=%d", s1, len(s1), cap(s1))

	s3 := []int{1, 3, 5}
	s4 := s3
	fmt.Println(s3,s4)
	
}
