package main

import (
	"time"

	"github.com/GoFYL/studygo/day06/mylogger"
)

var log mylogger.Logger //声明一个全局的接口变量

//测试我们自己写的日志库
func main() {
	log = mylogger.NewConsoleLogger("Info")                       //终端日志实例
	log = mylogger.NewFileLogger("Info", "./", "my.log", 10*1024) //文件日志实例
	for {
		log.Debug("这是一条Debug日志")
		log.Info("这是一条Info日志")
		log.Warning("这是一条Warning日志")
		id := 10010
		name := "lixiang"
		log.Error("这是一条Error日志,id:%d,name :%s", id, name)

		log.Fatal("这是一条Fatal日志")
		time.Sleep(2 * time.Second)
	}
}
