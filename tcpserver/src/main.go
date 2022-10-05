package main

import (
	"bufio"
	"fmt"
	"net"
)

func main() {
	listen, err := net.Listen("tcp", "127.0.0.1:3000")
	if err != nil {
		fmt.Println("listen() failed, err:", err)
		return
	}
	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("Accept() failed, err:", err)
			continue
		}

		go process(conn) // 启动一个协程来处理连接
	}
}

func process(conn net.Conn) {
	defer conn.Close() // 关闭连接
	for {
		reader := bufio.NewReader(conn)
		var buf [128]byte
		n, err := reader.Read(buf[:])
		if err != nil {
			fmt.Println("raed failed, err:", err)
		}
		recvStr := string(buf[:n])
		fmt.Println("接受到的数据是:", recvStr)
		conn.Write([]byte(recvStr))
	}
}
