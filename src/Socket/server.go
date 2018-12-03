package main

import (
	"fmt"
	"net"
	"regexp"
	"time"
)

type User struct {
	name    string
	address string
	C       chan string
}

var usermap map[string]*User = make(map[string]*User)
var msg = make(chan [2]string)

func char(conn net.Conn) {
	defer conn.Close()
	address := conn.RemoteAddr().String()
	user := &User{address: address, name: address, C: make(chan string)}
	usermap[address] = user

	go func() {
		for {
			info := <-user.C
			conn.Write([]byte(info))
		}
	}()
	msg <- [2]string{address, user.name + ":" + "已经上线"}
	var timeout time.Duration = 10 * time.Second
	timer := time.NewTimer(timeout)
	buferr := make([]byte, 1024)
	go func() {
		for {
			n, _ := conn.Read(buferr)
			if n == 0 {
				msg <- [2]string{address, user.name + ":" + "已经下线"}
				delete(usermap, address)
				fmt.Println(address, ":客户端已退出")
				return
			}
			info := string(buferr[:n-2])

			//重命名 格式rename|xxxx
			if regexp.MustCompile(`^rename\|.+`).MatchString(info) {
				user.name = info[7:]
				conn.Write([]byte("重命名成功"))
				continue
			}
			timer.Reset(timeout)
			//显示当前所哟用户 userlist
			if regexp.MustCompile(`^userlist$`).MatchString(info) {
				result := "当前在线用户列表：\n--------------------------------------\n"
				for _, u := range usermap {
					result += u.name + "\n"
				}
				result += "--------------------------------------"
				user.C <- result
				continue
			}
			//@功能，将信息单独发送到某个用户 格式：@zhangsan|dfhsjdsjdjs
			reg := regexp.MustCompile(`^@(.+?)\|(.+)`)
			if reg.MatchString(info) {
				result := reg.FindStringSubmatch(info)
				username := result[1]
				userinfo := result[2]
				for _, u := range usermap {
					if u.name == username {
						u.C <- userinfo
					}
				}
				continue
			}
			msg <- [2]string{address, user.name + "：" + info}
		}

	}()
	<-timer.C
	delete(usermap, address)
}
func main() {
	server, err := net.Listen("tcp", "127.0.0.1:8080")
	if err != nil {
		fmt.Println("客户端错误")
		return
	}
	defer server.Close()
	go func() {

		for {
			showinfo := <-msg
			for k, user := range usermap {
				if k != showinfo[0] {

					user.C <- showinfo[1]
				}
			}
		}
	}()
	for {
		conn, _ := server.Accept()
		fmt.Println(conn.RemoteAddr().String(), "已经连接服务器")
		go char(conn)
	}

}
