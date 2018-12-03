package main

import (
	"fmt"
	"net"
)

type User struct {
	C       chan string // 用于接收信息的管道
	name    string      // 用户名
	address string      // 用户地址
}

// 服务器接收用户信息的管道(包含用户标识)
var msg = make(chan [2]string)

// 当前的在线用户
var online map[string]User = make(map[string]User)

func chat(conn net.Conn) {
	buf := make([]byte, 1024)

	defer conn.Close()

	// 把当前连接(用户)保存到online中，以便服务器群发信息
	address := conn.RemoteAddr().String()
	user := User{make(chan string), address, address}
	online[address] = user

	// 开启新任务：检测用户管道是否有数据，一旦有数据，则写到客户端
	go func() {
		for {
			info := <-user.C
			conn.Write([]byte(info))
		}
	}()

	// 持续读取
	for {
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("读取异常：", err)
			return
		}

		info := address + "：" + string(buf[:n-2])
		msg <- [2]string{address, info} // 将信息存储到管道中
	}
}

func main() {
	server, err := net.Listen("tcp", "127.0.0.1:8080")
	if err != nil {
		fmt.Println("服务器异常：", err)
		return
	}

	// 延迟关闭服务器
	defer server.Close()

	// 开启任务：从服务器信息管道中获取数据，然后将数据写入到每个用户的管道中
	go func() {
		for {
			// 管道中除了发送的信息以外，还包括发送人的唯一标示
			msgWithAddr := <-msg
			// 遍历在线用户，将信息塞入管道
			for k, u := range online {
				// 仅发送给非发送者的用户
				if k != msgWithAddr[0] {
					u.C <- msgWithAddr[1]
				}
			}
		}
	}()

	// 处理多个连接
	for {
		// 等待用户连接服务器：如果用户没有连接，会阻塞
		conn, err := server.Accept()
		fmt.Println(conn.RemoteAddr().String() + "连接至服务器！")
		if err != nil {
			fmt.Println("等待异常：", err)
			return
		}
		go chat(conn) // 01-开启协程，处理当前用户的连接
	}

}
