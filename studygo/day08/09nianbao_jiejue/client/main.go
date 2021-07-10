package main

import (
	"fmt"
	"net"

	"github.com/GoFYL/studygo/day08/09nianbao_jiejue/protocol"
)

// socket_stick/client/main.go

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:30000")
	if err != nil {
		fmt.Println("dial failed, err", err)
		return
	}
	defer conn.Close()
	for i := 0; i < 20; i++ {
		msg := `Hello, Hello. How are you?`
		//调用协议编码数据
		b, _ := protocol.Encode(msg)
		conn.Write([]byte(b))
		// time.Sleep(time.Second)
	}
}
