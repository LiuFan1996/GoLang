package main

import "os"

func main() {
	copyPath := "d:/soft/保护色.doc"
	locaPath := "c:/保护色.doc"

	// 1、 读取文件
	copy, _ := os.Open(copyPath)
	// 2、写入文件
	loca, _ := os.Create(locaPath)

	// 关闭文件
	defer copy.Close()
	defer loca.Close()

	buf := make([]byte, 1024)
	for {
		n, _ := copy.Read(buf)
		if n == 0 { // 文件读取完毕
			return
		}
		loca.Write(buf[:n])
	}
}
