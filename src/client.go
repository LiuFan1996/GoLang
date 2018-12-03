package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8080")
	if err != nil {
		fmt.Println("连接服务器异常：", err)
		return
	}

	defer conn.Close()

	// 读取服务器信息
	go func() {
		buf := make([]byte, 1024)
		for {
			n, err := conn.Read(buf)
			if err != nil {
				fmt.Println("读取异常：", err)
				return
			}
			fmt.Println(string(buf[:n]))
		}
	}()

	input := make([]byte, 1024)
	for {
		// 读取键盘输入的信息，如果用户没有输入，则阻塞
		n, _ := os.Stdin.Read(input)
		conn.Write(input[:n])
	}

}
