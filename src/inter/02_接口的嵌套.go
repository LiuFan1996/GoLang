package main

import "fmt"

type InterFu interface {
	Eat()
	Run()
}

type InterZi interface {
	InterFu
	Test()
}

type Dog2 struct {
	name string
	age  int
}

func (d Dog2) Eat() {
	fmt.Println("eat eat eat")
}

func (d Dog2) Run() {
	fmt.Println("eat eat eat")
}

func (d Dog2) Test() {
	fmt.Println("eat eat eat")
}

func main() {
	var inter InterFu = Dog2{}
	inter.Run()
	var inter2 InterZi = Dog2{}
	inter2.Test()
	/*Object o1 = 1;
	Object o2 = "xxx";*/

	/*var o1 interface{} = 1
	var o2 interface{} = "xxx"
	var o3 interface{} = Dog2{}*/

}
