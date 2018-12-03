package main

import "fmt"

type Pserson struct {
	name string
	age   int
	sex   string
}
func main() {
	//方法是特殊的函数，定义在某一特定的类型上，通过类型的实例来进行调用，这个实例被叫接收者(receiver)。
	//函数将变量作为参数：Function1(recv)
	//方法在变量上被调用：recv.Method1()
	//接收者必须有一个显式的名字，这个名字必须在方法中被使用。
	//receiver_type 叫做 （接收者）基本类型，这个类型必须在和方法同样的包中被声明。

	//Go语言不允许为简单的内置类型添加方法，所以下面定义的方法是非法的。
	//func (a int) save(){
	//
	//}

	//但是为int类型起一个别名类型就可以
	//type myint int
	//func (a myint) save()(n int){
	//	return n
	//}
	var p Pserson
	p.set("张三",90,"男")
	fmt.Println(p)

}
type myint int
func (a myint) save()(n int){
	return n
}

func (p *Pserson) set(name string,age int,sex string){
	p.name=name
	p.age=age
	p.sex=sex
}