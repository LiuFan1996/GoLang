package tea

import "fmt"

func main() {
	ch := make(chan string) // 默认能存能取
	fmt.Println(ch)
	ch_read := make(<-chan string)
	//ch_read <- "xxx" //Invalid operation: ch_read <- "xxx" (send to receive-only type <-chan string)
	//<-ch_read
	fmt.Println(ch_read)
	ch_write := make(chan<- string)
	ch_write <- "hello"
	//<-ch_write
	fmt.Println(ch_write)
}
