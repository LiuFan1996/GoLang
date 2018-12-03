package main

import (
	"fmt"
	"net/http"
	"os"
)

// resp表示响应，用于往客户端会送数据
// req表示请求，用于获取客户端发送过来的数据
func hello(resp http.ResponseWriter, req *http.Request) {
	resp.Write([]byte("hello"))
}

func test(resp http.ResponseWriter, req *http.Request) {
	resp.Write([]byte("test"))
}

// 设置web资源在服务器中位置
var webapps = "./webapps/"

func allPage(resp http.ResponseWriter, req *http.Request) {
	uri := req.RequestURI // 获取访问的路径
	// 根据访问的路径，读取对应的文件
	// 假设路径为：/index.html
	filePath := webapps + uri[1:]
	fmt.Println(filePath)
	file, _ := os.Open(filePath)
	defer file.Close()
	buf := make([]byte, 1024)
	for {
		n, _ := file.Read(buf)
		if n == 0 {
			break
		}
		resp.Write(buf[:n])
	}

}

func main() {
	// 绑定函数，当访问的路径是/index.html的时候，执行hello这个函数
	//http.HandleFunc("/index.html", hello)
	//http.HandleFunc("/test.html", test)
	// 如果访问的路径没有注册
	http.HandleFunc("/", allPage)
	// 监听
	http.ListenAndServe("127.0.0.1:8000", nil)
}
