package main

import (
	"fmt"
	"os"
)

func main() {
	file, _ := os.Open("f:/sdsd.txt")
	create, _ := os.Create("e:/sdsd.txt")

	buffer := make([]byte, 1024)
	fmt.Println("文件加载开始")
	for {
		n, _ := file.Read(buffer)
		if n == 0 {
			return
		}
		fmt.Println(string(buffer[:n]))
		create.Write(buffer[:n])
	}

	defer file.Close()
	defer create.Close()
}
