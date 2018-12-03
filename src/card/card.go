package card

import (
	"fmt"
	"math/rand"
	"time"
)

type CcardStr struct {
	Ccard     [54]string
}


func (c *CcardStr) ShengCheng() {
	cardNum :=[13]string{"A","02","03","04","05","06","07","08","09","10","J","Q","K"}
	cardType:= [6]string{"红桃","黑桃","方块","草花","小王","大王"}
	for i:=0;i<=12;i++  {
		c.Ccard[i]=cardType[0]+cardNum[i]
		c.Ccard[i+13]=cardType[1]+cardNum[i]
		c.Ccard[i+26]=cardType[2]+cardNum[i]
		c.Ccard[i+39]=cardType[3]+cardNum[i]
	}
	c.Ccard[52]=cardType[4]
	c.Ccard[53]=cardType[5]
	fmt.Println(c.Ccard)
}

func (c CcardStr) XiPai(){
	rand.Seed(time.Now().UnixNano())
	for i:=0;i<1000 ;i++  {
		r1 :=rand.Intn(54)
		r2 :=rand.Intn(54)
		c.Ccard[r1],c.Ccard[r2]=c.Ccard[r2],c.Ccard[r1]
	}
	fmt.Println(c.Ccard)
}
