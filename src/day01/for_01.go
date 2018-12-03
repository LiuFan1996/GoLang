package main

import (
	"fmt"
)

func main() {
	//for循环的三种表现形式及使用
	//第一种（类似与其他语言的while）
	var i=10
	for i>0{
		i--
		fmt.Println(i)
	}

	//第二种写法。类似其他语言的for ，但是没有（）
	 sum:=0
	for j:=0;j<10;j++{
		sum+=j
	}
	fmt.Println(sum)


	//第三种写法类似于其他语言的foreach，在go种使用的range函数
	//第三种写法也可以分为好几种，这里只例举一种
	var a =[]int{1,2,3,4,5}
	for _,s1:=range a{
		fmt.Println(s1)
	}

}
