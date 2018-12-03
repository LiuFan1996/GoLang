package tea

import "fmt"

/*func main1() {
	// 斐波那契数列：1 1 2 3 5 8 13...
	x, y := 1, 1
	for i := 0; i < 10; i++ {
		fmt.Println(x)
		x, y = y, x+y
	}
}*/

/*func fibonacci(c, quit chan int) {
	x, y := 1, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func main() {
	c := make(chan int)
	quit := make(chan int)

	go func() {
		for i := 0; i < 6; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()

	fibonacci(c, quit)
}*/

func RunnerA(a chan string) {
	for i := 0; i < 100; i++ {

	}
	a <- "A先执行完"
}

func RunnerB(b chan string) {
	for i := 0; i < 100; i++ {

	}
	b <- "B先执行完"
}

func main() {
	chA := make(chan string)
	chB := make(chan string)
	go RunnerA(chA)
	go RunnerB(chB)
	for {
		select {
		case msg := <-chA:
			fmt.Println(msg)
			return
		case msg := <-chB:
			fmt.Println(msg)
			return
		}
	}
}
