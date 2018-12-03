package main

import "fmt"

/*
	键盘输入求和：a+aa+aaa+aaaa
	输入 a，和a出现的的位数
*/
func main() {
	var a,n int
	fmt.Println("请输入a=")
	fmt.Scan(&a)
	fmt.Println("请输入a出现的最大位数")
	fmt.Scan(&n)
    sum:= sumAN(a,n)
    fmt.Println("a的n位之和位",sum)
}
func sumAN(a,n int )(value int){
	var sum int =0
	var fcn int =0
	for i:=0;i<n ;i++  {
		sum+=a*CiFan(10,i)
		fcn+=sum
		fmt.Println(sum,fcn)
	}
return fcn
}
func CiFan(a,n int)(value int){
	var sum int=1
	if n==0{
		return 1
	}else if n==1 {
		return a
	}
	for i:=0;i<n ;i++  {
		sum*=a
	}
	return sum
}