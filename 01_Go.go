package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var i ,sum int
	for i=0;i<=100;i++ {
	sum+=i;
	}
	fmt.Println(sum)
 	str :="abcde"
	for i, s:=range str  {
		fmt.Printf("i=%d,s=%c\n",i,s)
	}

	for i, _=range str{
		fmt.Println("i=",i)
	}

	for _ ,d:=range str{
		if(d=='c'){
			break
		}
		fmt.Printf("c=%c\n",d)
	}

	for _ ,d:=range str{
		if(d=='c'){
			continue
		}
		fmt.Printf("c=%c\n",d)
	}
	for i=5;i<10;i++{
		if(i==7){
			fmt.Printf("i=%d\n",i)
			goto Span
		}

	}
	Span:
		fmt.Println("goto跳转")

	if c:=100;c>10{
		fmt.Println(c)
	}
	//用时间戳作为随机数的种子
	rand.Seed(time.Now().UnixNano())
	k:=rand.Intn(100)
	fmt.Println(k)
}
