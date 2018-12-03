package main

func main() {
	
}
//无参无返回值
func test(){

}

//不定参数无返回值
func test1(n int, arrs...int){

}
//有参无返回值
func test2(n,s int){

}

//无参有返回值
func test3()int{
return 0
}

//无参有不定返回值
//func test4() (arrs...int){
//	return 1,2,3
//}

//无参多值返回
func test5()(n,k int){
	return 1,2
}
//有参有返回值
func test6(n int)(k int){
	return n
}
//多参多返回值
func test7(n,k,j int)(q,w,e int){
	return n,k,j
}