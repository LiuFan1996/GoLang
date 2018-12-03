package main

import (
	"fmt"
)

func main() {
	a := []int{1, 2, 3, 4, 5}
	var b = a[0:3]
	var c = [...]int{3, 6, 9, 2, 6, 4}
	d := c[0:2]
	sliceInfo(b)

	fmt.Printf("sum of b is %d\n", sum(b))
	fmt.Printf("sum of d is %d\n", sum(d))

}
func sum(a []int) int {
	s := 0
	for i := 0; i < len(a); i++ {
		s += a[i]
	}
	return s

}

func sliceInfo(x []int) {
	fmt.Printf("len is %d ,cap is %d,  slice is %v\n", len(x), cap(x), x)

}
