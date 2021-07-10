package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// net/http server

//
func f1(w http.ResponseWriter, r *http.Request) {
	str, err := ioutil.ReadFile("./xx.txt")
	if err != nil {
		w.Write([]byte(fmt.Sprintf("%v", err)))
		// w.Write([]byte("暂无内容!"))
		return
	}
	w.Write([]byte(str))
}
func f2(w http.ResponseWriter, r *http.Request) {
	//对于GET请求，参数都放在URL上（query param），请求体中是没有数据的。
	queryparam := r.URL.Query() //自动帮我们识别URL中的 query param
	name := queryparam.Get("name")
	age := queryparam.Get("age")
	fmt.Println(name, age)
	fmt.Println(r.Method)
	fmt.Println(ioutil.ReadAll(r.Body)) //我在服务端打印客户端发来的请求的body
	w.Write([]byte("ok"))
}

func main() {
	http.HandleFunc("/posts/Go/15_sokect/", f1)
	http.HandleFunc("/xxx/", f2)
	http.ListenAndServe("127.0.0.1:9000", nil)
}
