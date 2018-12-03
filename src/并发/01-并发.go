package main

import (
	"fmt"
	"runtime"
	"time"
)

func ct() {
	defer fmt.Println("已经结束")
	for i := 0; i < 100; i++ {
		fmt.Println(i)
	}
	runtime.Goexit()
	for {
		fmt.Println("循环")
	}
}
func main() {

	go ct()
	//runtime.Gosched()
	time.Sleep(time.Second)
}
