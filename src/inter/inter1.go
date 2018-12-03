package main

import "fmt"

type Person interface {
	Teacher
	Run()
	Eat()
}

type Teacher interface {
	Class()
	Learn()
}

type Personstr struct {
	name string
	age  int
}

func (p Personstr) Run() {
	fmt.Println("name=", p.name, "   age=", p.age, "。。。在走路")
}

func (p Personstr) Eat() {
	fmt.Println("name=", p.name, "   age=", p.age, "。。。在吃东西")
}

func (p Personstr) Class() {
	fmt.Println("name=", p.name, "   age=", p.age, "。。。在上课")
}
func (p Personstr) Learn() {
	fmt.Println("name=", p.name, "   age=", p.age, "。。。在学习")
}
func main() {
	var person Person = Personstr{"张三", 45}
	person.Run()
	person.Eat()
	person.Class()
	person.Learn()
}
