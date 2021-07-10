package main

import (
	"fmt"
	"time"
)

//时间

func f1() {
	now := time.Now()
	fmt.Println(now.Year())

	//时间戳
	fmt.Println(now.Unix())     //1970.1.1到现在的秒数
	fmt.Println(now.UnixNano()) //1970.1.1到现在的纳秒数
	//time.Unix()
	ret := time.Unix(1623202094, 0)
	fmt.Println(ret)
	fmt.Println(ret.Year())
	fmt.Println(ret.Day())
	//时间间隔
	fmt.Println(time.Second)

	//now+ 时间间隔
	now.Add(time.Hour)
	//定时器
	// timer := time.Tick(time.Second)
	// for t := range timer {
	// 	fmt.Println(t) //一秒钟执行一次
	// }

	//Sub两个时间相减
	nextYear, err := time.Parse("2006-01-02", "2022-06-09")
	if err != nil {
		fmt.Printf("err:%v\n", err)
		return
	}
	d := now.Sub(nextYear)
	fmt.Println(d)
	fmt.Println("________________________")
	//格式化时间,把语言中时间对象，转换成字符串类型的时间
	//2021-08-03
	fmt.Println(now.Format("2006-01-02"))
	//2021/08/03 11:50:50
	fmt.Println(now.Format("2006/01/02 03:04:05"))
	//2021/08/03 11:50:50 AM
	fmt.Println(now.Format("2006/01/02 15:04:05 PM"))
	//2021/08/03 11:50:50.123\
	fmt.Println(now.Format("2006/01/02 03:04:05.000"))
	//按照对应的格式解析字符串类型的时间
	timeObj, err := time.Parse("2006-01-02", "2021-06-09")
	if err != nil {
		fmt.Printf("err:%v\n", err)
		return
	}
	fmt.Println(timeObj)
	fmt.Println(timeObj.Unix())

	//Sleep
	n := 5 // int
	fmt.Println("开始sleep")
	time.Sleep(time.Duration(n) * time.Second)
	fmt.Println("5秒钟过去了")

}

//时区

func f2() {
	now := time.Now() //本地的时间
	fmt.Println(now)
	///明天的这个时间
	//按照指定格式去解析一个字符串格式的时间
	time.Parse("2006-01-02 15:04:05", "2021-06-10 10:28:00")
	//按照东八区的时区和格式去解析一个字符串格式的时间
	//根据字符串加载时区
	loc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		fmt.Println("load loc failed,err:", err)
		return
	}
	//根据指定时区解析时间
	timeObj, err := time.ParseInLocation("2006-01-02 15:04:05", "2021-06-10 10:28:00", loc)
	if err != nil {
		fmt.Println("parse time failed,err:", err)
		return
	}
	fmt.Println(timeObj)
	//时间对象相减
	td := timeObj.Sub(now)
	fmt.Println(td)
}

func main() {
	//f1()
	f2()
}
