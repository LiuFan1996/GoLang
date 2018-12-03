package tea

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int, 0) //无缓冲的通道

	go func() {
		for i := 0; i < 3; i++ {
			c <- i
			fmt.Println("存放成功：", i)
		}
	}()

	time.Sleep(3 * time.Second) //延时3s

	for i := 0; i < 3; i++ {
		num := <-c //从c中接收数据，并赋值给num
		fmt.Println("获取成功：", num)
	}
}
