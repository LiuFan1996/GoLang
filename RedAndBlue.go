package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

//Go语言 双色球案例
func main() {
	poll:=[33]string{"01","02","03","04","05","06","07","08","09",
	"10","11","12","13","14","15","16","17","18","19","20","21","22",
	"23","24","25","26","27","28","29","30","31","32","33"}
	fmt.Println(poll)
	var Bpoll =[]string{"0","0","0","0","0","0"}
	var Rpoll =[]string{"0"}
	for i:=0;i<7;i++  {
		rand.Seed(time.Now().UnixNano())
		k:=rand.Intn(33)
		key:= poll[k]
		if i<=5{
		if isNot(Bpoll,key){
			Bpoll[i]=key
		}else{
			i--
			continue
		}
		}
		if i==6{
			rand.Seed(time.Now().UnixNano())
			key=poll[rand.Intn(16)]
				Rpoll[i-6]=key
		}
	}
	sort.Strings(Bpoll)
	fmt.Println("蓝球：",Bpoll)
	fmt.Println("红球: ",Rpoll)
}

func isNot(a []string,key string )( value bool){
	for i:=0; i<len(a);i++{
		if a[i]==key{
		return false}
	}
return true
}