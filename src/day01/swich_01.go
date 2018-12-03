package main

import (
	 "math/rand"
	"fmt"
)

func main() {

	//swich 的一种使用方式
	switch score:=rand.Intn(100); {
	case score >= 90:
		fmt.Println("优秀")
	case score>=80:
		fmt.Println("良好",score)
	case score>=60:
		fmt.Println("良好")
	default:
		fmt.Println("不及格")
	}
}
