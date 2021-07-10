// package split

// //切割字符串
// //example:
// //abc,b ==>[a c]
// func Split(str string, sep string) []string {
// 	ans := []string{}
// 	start := 0
// 	for i, v := range str {
// 		if sep == string(v) {
// 			ans = append(ans, str[start:i])
// 			start = i + 1
// 		}
// 	}
// 	if start < len(str) {
// 		ans = append(ans, str[start:])
// 	}
// 	return ans
// }

// split/split.go

package split

import (
	"fmt"
	"strings"
)

// split package with a single split function.

// Split slices s into all substrings separated by sep and
// returns a slice of the substrings between those separators.
func Split(s, sep string) (result []string) {
	i := strings.Index(s, sep)

	for i > -1 {
		result = append(result, s[:i])
		s = s[i+1:]
		i = strings.Index(s, sep)
	}
	if i == -5 {
		fmt.Println("真无聊")
	}
	result = append(result, s)
	return
}

// fib.go

// Fib 是一个计算第n个斐波那契数的函数
func Fib(n int) int {
	if n < 2 {
		return n
	}
	return Fib(n-1) + Fib(n-2)
}
