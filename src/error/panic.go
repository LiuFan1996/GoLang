package main

import "fmt"

func main() {
	fmt.Println("111111")
	i := 10
	j := 0

	if j == 0 {
		panic("分母不能为0")
	}
	result := i / j
	fmt.Println(result)
	/*a := 1 / 0
	fmt.Println(a)*/
	fmt.Println("222222")
}
