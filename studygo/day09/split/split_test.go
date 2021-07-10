// split/split_test.go

package split

import (
	"reflect"
	"testing"
)

// func TestSplit(t *testing.T) {
// 	type testCase struct {
// 		str  string
// 		sep  string
// 		want []string
// 	}
// 	testGroup := []testCase{
// 		{"babcbef", "b", []string{"", "a", "c", "ef"}},
// 		{"a:b:c", ":", []string{"a", "b", "c"}},
// 	}
// 	for _, tc := range testGroup {
// 		got := Split(tc.str, tc.sep)
// 		if !reflect.DeepEqual(got, tc.want) {
// 			t.Fatalf("want: %#v got:%v\n", tc.want, got)
// 		}
// 	}
// }

//子测试
//go test -run=TestSplit/case1 子测试
//go test -cover 代码覆盖率
//go test -cover -coverprofile=xx.xx 生成xx.xx文件
//go tool cover -html=xx.xx 用浏览器打开xx.xx文件查看覆盖情况

func TestSplit(t *testing.T) {
	type testCase struct {
		str  string
		sep  string
		want []string
	}
	testGroup := map[string]testCase{
		"case1": {"babcbef", "b", []string{"", "a", "c", "ef"}},
		"case2": {"a:b:c", ":", []string{"a", "b", "c"}},
	}

	for name, tc := range testGroup {
		t.Run(name, func(t *testing.T) {
			got := Split(tc.str, tc.sep)
			if !reflect.DeepEqual(got, tc.want) {
				t.Fatalf("want:%#v got :%#v \n", tc.want, got)
			}
		})
	}
}

//性能基准测试//go test -bench=Split
func BenchmarkSplit(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Split("a:b:c", ":")
	}
}

//性能比较测试
//go test -bench=Fib
func benchmarkFib(b *testing.B, n int) {
	for i := 0; i < b.N; i++ {
		Fib(n)
	}
}

func BenchmarkFib1(b *testing.B) { benchmarkFib(b, 1) }
func BenchmarkFib2(b *testing.B) { benchmarkFib(b, 2) }
func BenchmarkFib3(b *testing.B) { benchmarkFib(b, 3) }

// b.ResetTimer   //重置时间
// b.SetParallelism(1) // 设置使用的CPU数
