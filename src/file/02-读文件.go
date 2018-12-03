package main

import (
	"fmt"
	"os"
)

func main() {
	f, e := os.Open("./xxx.txt")
	if e != nil {
		fmt.Println("err = ", e)
		return
	}
	defer f.Close()

	byte := make([]byte, 1024)
	n, _ := f.Read(byte)
	fmt.Println(string(byte[:n]))

}
