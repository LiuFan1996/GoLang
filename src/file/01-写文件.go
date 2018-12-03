package main

import (
	"fmt"
	"os"
)

func main() {
	f, err := os.Create("xxx.txt")
	if err != nil {
		fmt.Println("err = ", err)
		return
	}

	// 关闭文件
	defer f.Close()

	f.WriteString("hello")
	f.Write([]byte(" world"))
}
