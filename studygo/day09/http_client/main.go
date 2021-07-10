package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

//net/http Client

// //公用一个client 适用于请求比较频繁
// var (
// 	client = http.Client{
// 		Transport: &http.Transport{
// 			DisableKeepAlives: false,
// 		},
// 	}
// )

func main() {
	// resp, err := http.Get("http://127.0.0.1:9000/xxx/?name=fyl&age=1")
	// if err != nil {
	// 	fmt.Println("get url failed,err :", err)
	// 	return
	// }
	urlObj, _ := url.Parse("http://127.0.0.1:9000/xxx/")
	data := url.Values{} //URL values
	data.Set("name", "周林")
	data.Set("age", "9000")
	queryStr := data.Encode() //URL encode 之后的地址
	fmt.Println(queryStr)
	k := urlObj.String() + "?" + queryStr
	req, err := http.NewRequest("GET", k, nil)
	// resp, err := http.DefaultClient.Do(req)
	// if err != nil {
	// 	fmt.Println("get url failed,err :", err)
	// 	return
	// }

	//适用于请求不是特别频繁，用完就关闭该连接
	//禁用KeepAlives的client
	tr := &http.Transport{
		DisableKeepAlives: true,
	}
	client := http.Client{
		Transport: tr,
	}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("get url failed,err :", err)
		return
	}
	defer resp.Body.Close() // 一定要记得关闭resp.Body
	//从resp中把服务端返回的数据读出来
	b, err := ioutil.ReadAll(resp.Body) //我在客户端读出服务端返回的响应的body
	if err != nil {
		fmt.Println("ReadAll failed,err :", err)
		return
	}
	fmt.Println(string(b))
}
