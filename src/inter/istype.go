package main

import "fmt"

func main() {
	//次数引入空接口类型，可以接受任何类型的数据，相当于JAVA中的Object类
	var in interface{} = 1

	if i, ok := in.(int); ok {
		fmt.Printf("int 类型的in=%d\n", i)
	}

	switch ok := in.(type) {
	case int:
		fmt.Printf("int 类型的in=%d", ok)
	case string:
		fmt.Printf("int 类型的in=%s", ok)
	default:
		fmt.Printf("int 类型的in=%v", ok)

	}
}
