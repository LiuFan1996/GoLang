package main

import "fmt"

func main() {
	ch := make(chan int, 10)
	go func() {
		for i := 0; i < 20; i++ {
			ch <- i
			fmt.Println("存放i=", i)
		}
	}()
	for i := 0; i < 20; i++ {
		fmt.Println("获取i=", <-ch)
	}
	////只写
	//var c chan<- int
	//c <- 15
	////只读
	//var t <-chan int
	//f := <-t
	//fmt.Println(f)

	//close 与range 注意close 的使用，引文当使用range遍历管道时
	// 如果不关闭管道，很容易形成死循环，会阻塞在遍历函数range一行
	c1 := make(chan int)
	go func() {
		for i := 0; i <= 3; i++ {
			c1 <- i
		}
		close(c1)
	}()
	for val := range c1 {
		fmt.Println(val)
	}

	//使用select关键字实现斐波那契数列 1 1 2 3 5 8 13 21 34 55 89
	x, y := 1, 1
	for i := 0; i < 6; i++ {
		fmt.Println(x)
		x, y = y, x+y
	}

	var f1 = make(chan int) //管道一控制次数
	var f2 = make(chan int) //管道二控制退出
	go func() {
		x, y = 1, 1
		for {
			select {
			case f1 <- x:
				x, y = y, x+y
			case <-f2:
				fmt.Print("退出 ")
				return
			}
		}
	}()

	for i := 0; i < 6; i++ {
		fmt.Print(<-f1, " ")
	}
	f2 <- 0

	//通过管道判断那个协程速度更快
	go func() {
		for i := 0; i < 100; i++ {

		}
	}()
}
