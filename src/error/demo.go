package main

import (
	"errors"
	"fmt"
)

func main() {
	var err1 error = errors.New("a normal err1")
	fmt.Println(err1) //a normal err1

	var err2 error = fmt.Errorf("%s", "a normal err2")
	fmt.Println(err2) //a normal err2

}
