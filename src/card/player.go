package card

import "fmt"

type PlayerCardStr struct {
	Pname         string
	Pcard     [17]string
}

func (pc *PlayerCardStr) ToString(){
	fmt.Println("Pname=",pc.Pname)
	fmt.Println("Pcard=",pc.Pcard)
}