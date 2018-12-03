package main

import (
	"fmt"
)

func main() {

	//第一初始化 共四种初始化方式 []中可以写长度，也可以不写
	var a []int
	b:=[]int{1,2,3}
	var c=[4]int{1,2,3,4}
	d:=[]int{1,2,3,4,5}
	fmt.Println(a,b,c,d)

	//遍历 ，使用for循环遍历
	for i ,f:=range d  {
		fmt.Println(i,f)
	}
	fmt.Println(typeof(c))
	//数组作为参数传递
	//当数组作为参数传递的时候，要注意目标函数中的数组参数类型有么1有定义长度，
	// 如果没有定义，则，传递参数时，也不能定义长度
	
}
func typeof(v interface{}) string {
	return fmt.Sprintf("%T", v)
}