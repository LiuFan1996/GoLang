package main

import "fmt"

func main() {
	var data interface{} = 1
	if d, ok := data.(string); ok {
		fmt.Printf("int 类型  d = %v", d)
	}

	switch value := data.(type) {
	case int:
		fmt.Printf("int  data = %v", value)
	case string:
		fmt.Printf("string  data = %v", value)
	default:
		fmt.Printf("other  data = %v", value)
	}

}
