package tea

import "fmt"

func main() {
	ch := make(chan bool)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(i)
		}
		ch <- true
	}()
	<-ch
}
