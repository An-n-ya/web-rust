package main

import (
	"fmt"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:3000")
	if err != nil {
		fmt.Println("err: ", err)
		return
	}
	defer conn.Close() // 关闭连接

	_, err = conn.Write([]byte("Hello"))
	if err != nil {
		return
	}

	buf := [512]byte{}
	// 返回了字节长度
	n, err := conn.Read(buf[:])
	if err != nil {
		fmt.Println("recv failed, err:", err)
		return
	}
	fmt.Println("Response from server", string(buf[:n]))

}
