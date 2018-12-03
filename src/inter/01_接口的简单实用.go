package main

import "fmt"

type Inter interface {
	Eat()
	Run()
}

type Dog struct {
	name string
	age  int
}

func (d Dog) Eat() {
	fmt.Println("吃吃吃")
}

func (d Dog) Run() {
	fmt.Println("跑跑跑")
}

type Cat struct {
	name string
	age  int
}

func (d Cat) Eat() {
	fmt.Println("eat eat eat")
}

func (d Cat) Run() {
	fmt.Println("eat eat eat")
}

func main() {
	var inter Inter = Dog{"阿才", 2}
	inter.Eat()
	var inter2 Inter = Cat{"小花", 5}
	inter2.Eat()
}
