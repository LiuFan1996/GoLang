package main

import (
	"fmt"
)

func XXX() (err error) {
	defer func() { //在发生异常时，设置恢复
		if x := recover(); x != nil {
			//panic value被附加到错误信息中；
			//并用err变量接收错误信息，返回给调用者。
			err = fmt.Errorf("internal error: %v", x)

		}
	}()
	panic("func TestB(): panic")
	return
}

func main() {
	err := XXX()
	fmt.Println(err)
	fmt.Println("aaaaaaaaaa")
}
